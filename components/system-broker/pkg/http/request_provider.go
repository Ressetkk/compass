/*
 * Copyright 2020 The Compass Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/kyma-incubator/compass/components/system-broker/pkg/log"
	"io"
	"net/http"
)

func NewRequestProvider(uidsrv UUIDService) *RequestProvider {
	return &RequestProvider{
		uuidService: uidsrv,
	}
}

type UUIDService interface {
	Generate() string
}

type RequestInput struct {
	Method     string
	URL        string
	Parameters map[string]string
	Body       interface{}
	Headers    map[string]string
}

type RequestProvider struct {
	uuidService UUIDService
}

func (rp *RequestProvider) Provide(ctx context.Context, input RequestInput) (*http.Request, error) {
	var bodyReader io.Reader
	if input.Body != nil {
		bodyBytes, err := json.Marshal(input.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	request, err := http.NewRequest(input.Method, input.URL, bodyReader)
	if err != nil {
		return nil, err
	}

	if len(input.Headers) != 0 {
		for key, value := range input.Headers {
			request.Header.Add(key, value)
		}
	}

	if len(input.Parameters) != 0 {
		q := request.URL.Query()
		for k, v := range input.Parameters {
			q.Set(k, v)
		}
		request.URL.RawQuery = q.Encode()
	}

	log.C(ctx).Infof("Sending request %s %s", request.Method, request.URL)

	return request, nil
}
