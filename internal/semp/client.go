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
	url          string
	username     string
	password     string
	retries      uint
	retryWait    time.Duration
	retryWaitMax time.Duration
}

type Option func(*Client)

func BasicAuth(username, password string) Option {
	return func(client *Client) {
		client.username = username
		client.password = password
	}
}

func Retries(numRetries uint, retryWait, retryWaitMax time.Duration) Option {
	return func(client *Client) {
		client.retries = numRetries
		client.retryWait = retryWait
		client.retryWaitMax = retryWaitMax
	}
}

func NewClient(url string, options ...Option) *Client {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	url += "SEMP/v2/config"
	client := &Client{
		Client:       http.DefaultClient,
		url:          url,
		retries:      3,
		retryWait:    time.Second,
		retryWaitMax: time.Second * 10,
	}
	for _, o := range options {
		o(client)
	}
	return client
}

func (c *Client) RequestWithBody(method, url string, body any) (map[string]any, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, c.url+url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	dumpData(fmt.Sprintf("%v to %v", request.Method, request.URL), data)
	return c.doRequest(request)
}

func (c *Client) doRequest(request *http.Request) (map[string]any, error) {
	if request.Method != http.MethodGet {
		request.Header.Set("Content-Type", "application/json")
	}
	request.SetBasicAuth(c.username, c.password)
	attemptsRemaining := c.retries + 1
	retryWait := c.retryWait
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
				// just continue
			default:
				body, _ := io.ReadAll(response.Body)
				return nil, fmt.Errorf("unexpected status %v (%v) during %v to %v, body:\n%s", response.StatusCode, response.Status, request.Method, request.URL, body)
			}
		}
		time.Sleep(retryWait)
		retryWait *= 2
		if retryWait > c.retryWaitMax {
			retryWait = c.retryWaitMax
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

func (c *Client) RequestWithoutBody(method, url string) (map[string]interface{}, error) {
	request, err := http.NewRequest(method, c.url+url, nil)
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
