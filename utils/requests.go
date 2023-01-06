// Copyright (c) 2023 dhn. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"net"
	"time"

	"github.com/corpix/uarand"
	"github.com/projectdiscovery/gologger"
	"github.com/valyala/fasthttp"
)

// GET HTTP wrapper
func GetHTTPRequest(url string, headers map[string]string) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.Add("User-Agent", uarand.GetRandom())

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}

	err := client.Do(req, resp)
	if err != nil {
		gologger.Error().Msgf(err.Error())
	}

	return resp
}

// POST HTTP wrapper
func PostHTTPRequest(url string, data []byte) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.SetBody(data)

	req.Header.SetMethod("POST")
	req.Header.Add("User-Agent", uarand.GetRandom())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second*10)
		},
	}

	err := client.Do(req, resp)
	if err != nil {
		gologger.Error().Msgf(err.Error())
	}

	return resp
}
