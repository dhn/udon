// Copyright (c) 2022 dhn. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sources

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dhn/udon/utils"
)

// GetHDData function returns all domains based on the given UA id name
func GetHDData(ua string) <-chan utils.Result {
	results := make(chan utils.Result)

	go func() {
		getHDData(fmt.Sprintf("https://api.hackertarget.com/analyticslookup/?q=%s",
			url.QueryEscape(ua)), results)
		close(results)
	}()

	return results
}

// Send a HTTP request and parse the JSON response
func getHDData(sourceURL string, results chan utils.Result) {
	resp := utils.GetHTTPRequest(sourceURL, map[string]string{})

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return
	}

	doc.Find("*").Each(func(i int, domain *goquery.Selection) {
		for _, result := range strings.Split(domain.Text(), "\n") {
			if len(strings.TrimSpace(result)) != 0 && !strings.Contains(result, "error getting results") {
				results <- utils.Result{Value: result, Source: "hackertarget"}
			}
		}
	})
}
