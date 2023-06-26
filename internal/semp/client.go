// terraform-provider-solacebroker
//
// Copyright 2023 Solace Corporation. All rights reserved.
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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Client struct {
	*http.Client
	url                string
	username           string
	password           string
	bearerToken        string
	retries            uint
	retryMinInterval   time.Duration
	retryMaxInterval   time.Duration
	requestMinInterval time.Duration
	rateLimiter        <-chan time.Time
}

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
		client.Client.Timeout = requestTimeoutDuration
		client.requestMinInterval = requestMinInterval
	}
}

func NewClient(url string, options ...Option) *Client {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	url += "SEMP/v2/config"
	client := &Client{
		Client:           http.DefaultClient,
		url:              url,
		retries:          3,
		retryMinInterval: time.Second,
		retryMaxInterval: time.Second * 10,
	}
	for _, o := range options {
		o(client)
	}
	if client.requestMinInterval > 0 {
		client.rateLimiter = time.Tick(client.requestMinInterval)
	} else {
		ch := make(chan time.Time)
		// closing the channel will make receiving from the channel non-blocking (the value received will be the
		//  zero value)
		close(ch)
		client.rateLimiter = ch
	}
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
	dumpData(fmt.Sprintf("%v to %v", request.Method, request.URL), data)
	return c.doRequest(request)
}

func (c *Client) doRequest(request *http.Request) (map[string]any, error) {
	// the value doesn't matter, it is waiting for the value that matters
	<-c.rateLimiter
	if request.Method != http.MethodGet {
		request.Header.Set("Content-Type", "application/json")
	}
	request.SetBasicAuth(c.username, c.password)
	attemptsRemaining := c.retries + 1
	retryWait := c.retryMinInterval
	var response *http.Response
	var err error
loop:
	for attemptsRemaining != 0 {
		response, err = c.Do(request)
		if err != nil {
			response = nil // make sure response is nil
		} else {
			switch response.StatusCode {
			case http.StatusOK:
				break loop
			case http.StatusTooManyRequests:
				// ignore the too many requests body and any errors that happen while reading it
				_, _ = io.ReadAll(response.Body)
				// just continue
			default:
				// ignore errors while reading the error response body
				body, _ := io.ReadAll(response.Body)
				return nil, fmt.Errorf("unexpected status %v (%v) during %v to %v, body:\n%s", response.StatusCode, response.Status, request.Method, request.URL, body)
			}
		}
		time.Sleep(retryWait)
		retryWait *= 2
		if retryWait > c.retryMaxInterval {
			retryWait = c.retryMaxInterval
		}
		attemptsRemaining--
	}
	if response == nil {
		return nil, err
	}
	rawBody, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not perform request: status %v (%v) during %v to %v, response body:\n%s", response.StatusCode, response.Status, request.Method, request.URL, rawBody)
	}
	data := map[string]any{}
	err = json.Unmarshal(rawBody, &data)
	if err != nil {
		return nil, fmt.Errorf("could not parse response body from %v to %v, response body was:\n%s", request.Method, request.URL, rawBody)
	}
	dumpData("response", rawBody)
	rawData, ok := data["data"]
	if ok {
		data, ok = rawData.(map[string]any)
	}
	if !ok {
		return nil, nil
	}
	return data, nil
}

func (c *Client) RequestWithoutBody(ctx context.Context, method, url string) (map[string]interface{}, error) {
	request, err := http.NewRequestWithContext(ctx, method, c.url+url, nil)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(os.Stderr, "===== %v to %v =====\n", request.Method, request.URL)
	return c.doRequest(request)
}

func dumpData(tag string, data []byte) {
	var in any
	_ = json.Unmarshal(data, &in)
	out, _ := json.MarshalIndent(in, "", "\t")
	fmt.Fprintf(os.Stderr, "===== %v =====\n%s\n", tag, out)
}
