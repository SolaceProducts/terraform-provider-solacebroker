// terraform-provider-solacebroker
//
// Copyright 2024 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package semp

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrBadRequest       = errors.New("bad request")
	ErrInvalidPath      = errors.New("invalid path")
)

var firstRequest = true

type Client struct {
	*retryablehttp.Client
	url                string
	username           string
	password           string
	bearerToken        string
	retries            uint
	retryMinInterval   time.Duration
	retryMaxInterval   time.Duration
	requestMinInterval time.Duration
	requestTimeout     time.Duration
	rateLimiter        <-chan time.Time
}

const (
	DefaultRetryMinInterval = 3 * time.Second
	DefaultRetryMaxInterval = 30 * time.Second
	DefaultRequestTimeout   = time.Minute
	DefaultRequestInterval  = 100 * time.Millisecond
	DefaultRetries          = 10
)

var Cookies = map[string]*http.Cookie{}

type Option func(*Client)

func BasicAuth(username, password string) Option {
	return func(client *Client) {
		client.username = username
		client.password = password
	}
}

func BearerToken(bearerToken string) Option {
	return func(client *Client) {
		client.bearerToken = bearerToken
	}
}

func Retries(numRetries uint, retryMinInterval, retryMaxInterval time.Duration) Option {
	return func(client *Client) {
		client.retries = numRetries
		client.retryMinInterval = retryMinInterval
		client.retryMaxInterval = retryMaxInterval
	}
}

func RequestLimits(requestTimeoutDuration, requestMinInterval time.Duration) Option {
	return func(client *Client) {
		client.requestTimeout = requestTimeoutDuration
		client.requestMinInterval = requestMinInterval
	}
}

func NewClient(url string, insecure_skip_verify bool, providerClient bool, options ...Option) *Client {
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure_skip_verify},
		MaxIdleConnsPerHost: 10,
	}
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Transport = tr
	if !providerClient {
		retryClient.Logger = nil
	}
	client := &Client{
		Client:           retryClient,
		url:              url,
		retries:          3,
		retryMinInterval: time.Second,
		retryMaxInterval: time.Second * 10,
	}
	for _, o := range options {
		o(client)
	}
	client.Client.RetryMax = int(client.retries)
	client.Client.RetryWaitMin = client.retryMinInterval
	client.Client.RetryWaitMax = client.retryMaxInterval
	client.HTTPClient.Timeout = client.requestTimeout
	client.HTTPClient.Jar, _ = cookiejar.New(nil)
	if client.requestMinInterval > 0 {
		client.rateLimiter = time.NewTicker(client.requestMinInterval).C
	} else {
		ch := make(chan time.Time)
		// closing the channel will make receiving from the channel non-blocking (the value received will be the
		//  zero value)
		close(ch)
		client.rateLimiter = ch
	}
	firstRequest = true
	return client
}

func (c *Client) RequestWithBody(ctx context.Context, method, url string, body any) (map[string]any, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx, method, c.url+url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("===== %v to %v =====", request.Method, request.URL))
	rawBody, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return parseResponseAsObject(ctx, request, rawBody)
}

func (c *Client) doRequest(request *http.Request) ([]byte, error) {
	if !firstRequest {
		// the value doesn't matter, it is waiting for the value that matters
		<-c.rateLimiter
	} else {
		// only skip rate limiter for the first request
		firstRequest = false
	}
	if request.Method != http.MethodGet {
		request.Header.Set("Content-Type", "application/json")
	}
	// Prefer OAuth even if Basic Auth credentials provided
	if c.bearerToken != "" {
		// TODO: add log
		request.Header.Set("Authorization", "Bearer "+c.bearerToken)
	} else if c.username != "" {
		request.SetBasicAuth(c.username, c.password)
	} else {
		return nil, fmt.Errorf("either username or bearer token must be provided to access the broker")
	}
	var response *http.Response
	var err error
	response, err = c.StandardClient().Do(request)
	if err != nil || response == nil {
		return nil, err
	}
	defer response.Body.Close()
	rawBody, err := io.ReadAll(response.Body)
	if err != nil || (response.StatusCode != http.StatusOK && response.StatusCode != http.StatusBadRequest) {
		return nil, fmt.Errorf("could not perform request: status %v (%v) during %v to %v, response body:\n%s", response.StatusCode, response.Status, request.Method, request.URL, rawBody)
	}
	if _, err := io.Copy(io.Discard, response.Body); err != nil {
		return nil, fmt.Errorf("response processing error: during %v to %v", request.Method, request.URL)
	}
	return rawBody, nil
}

func parseResponseAsObject(ctx context.Context, request *http.Request, dataResponse []byte) (map[string]any, error) {
	data := map[string]any{}
	err := json.Unmarshal(dataResponse, &data)
	if err != nil {
		return nil, fmt.Errorf("could not parse response body from %v to %v, response body was:\n%s", request.Method, request.URL, dataResponse)
	}
	rawData, ok := data["data"]
	if ok {
		// Valid data
		data, _ = rawData.(map[string]any)
		return data, nil
	} else {
		// Analize response metadata details
		rawData, ok = data["meta"]
		if ok {
			data, _ = rawData.(map[string]any)
			if data["responseCode"].(float64) == http.StatusOK {
				// this is valid response for delete
				return nil, nil
			}
			description := data["error"].(map[string]interface{})["description"].(string)
			status := data["error"].(map[string]interface{})["status"].(string)
			if status == "NOT_FOUND" {
				// resource not found is a special type we want to return
				return nil, fmt.Errorf("request failed from %v to %v, %v, %v, %w", request.Method, request.URL, description, status, ErrResourceNotFound)
			}
			tflog.Error(ctx, fmt.Sprintf("SEMP request returned %v, %v", description, status))
			return nil, fmt.Errorf("request failed for %v using %v, %v, %v", request.URL, request.Method, description, status)
		}
	}
	return nil, fmt.Errorf("could not parse response details from %v to %v, response body was:\n%s", request.Method, request.URL, dataResponse)
}

func parseResponseForGenerator(c *Client, ctx context.Context, basePath string, method string, request *http.Request, dataResponse []byte, appendToResult []map[string]any) ([]map[string]any, error) {
	data := map[string]any{}
	err := json.Unmarshal(dataResponse, &data)
	if err != nil {
		return nil, fmt.Errorf("could not parse response body from %v to %v, response body was:\n%s", request.Method, request.URL, dataResponse)
	}
	responseData := []map[string]any{}
	rawData, ok := data["data"]
	if ok {
		switch rawData := rawData.(type) {
		case []interface{}:
			responseDataRaw := rawData
			for _, t := range responseDataRaw {
				responseData = append(responseData, t.(map[string]any))
			}
		case map[string]interface{}:
			responseDataRaw := rawData
			responseData = append(responseData, responseDataRaw)
		}
		metaData, hasMeta := data["meta"]
		appendToResult = append(appendToResult, responseData...)
		if hasMeta {
			pageData, hasPaging := metaData.(map[string]any)["paging"]
			if hasPaging {
				nextPage := fmt.Sprint(pageData.(map[string]any)["nextPageUri"])
				nextPageUrl := strings.Split(nextPage, basePath)
				return c.RequestWithoutBodyForGenerator(ctx, basePath, method, nextPageUrl[1], appendToResult)
			}
		}
		return appendToResult, nil
	} else {
		rawData, ok = data["meta"]
		if ok {
			data, _ = rawData.(map[string]any)
			description := data["error"].(map[string]interface{})["description"].(string)
			status := data["error"].(map[string]interface{})["status"].(string)
			if status == "NOT_FOUND" {
				// resource not found is a special type we want to return
				return nil, fmt.Errorf("%v, %v, %w", description, status, ErrResourceNotFound)
			}
			if status == "INVALID_PATH" {
				// resource not found is a special type we want to return
				return nil, fmt.Errorf("%v, %v, %w", description, status, ErrInvalidPath)
			}
			return nil, fmt.Errorf("%v, %v, %w", description, status, ErrBadRequest)
		}
	}
	return nil, nil
}

func (c *Client) RequestWithoutBody(ctx context.Context, method, url string) (map[string]interface{}, error) {
	request, err := http.NewRequestWithContext(ctx, method, c.url+url, nil)
	if err != nil {
		return nil, err
	}
	tflog.Debug(ctx, fmt.Sprintf("===== %v to %v =====", request.Method, request.URL))
	rawBody, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return parseResponseAsObject(ctx, request, rawBody)
}

func (c *Client) RequestWithoutBodyForGenerator(ctx context.Context, basePath string, method string, url string, appendToResult []map[string]any) ([]map[string]interface{}, error) {
	request, err := http.NewRequestWithContext(ctx, method, c.url+url, nil)
	if err != nil {
		return nil, err
	}
	rawBody, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return parseResponseForGenerator(c, ctx, basePath, method, request, rawBody, appendToResult)
}
