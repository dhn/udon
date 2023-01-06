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

// GetSOData function returns all domains based on the given UA id name
func GetSOData(ua string) <-chan utils.Result {
	results := make(chan utils.Result)
	ua = strings.Replace(strings.ToUpper(ua), "UA", "", -1)

	go func() {
		getSOData(fmt.Sprintf("http://site-overview.com/website-report-search/analytics-account-id/%s",
			url.QueryEscape(ua)), ua, results)
		close(results)
	}()

	return results
}

// Send a HTTP request and parse the JSON response
func getSOData(sourceURL string, ua string, results chan utils.Result) {
	resp := utils.GetHTTPRequest(sourceURL, map[string]string{})

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return
	}

	doc.Find("u").Each(func(i int, domain *goquery.Selection) {
		for _, result := range strings.Split(domain.Text(), "\n") {
			if len(strings.TrimSpace(result)) != 0 && !strings.Contains(result, ua) {
				results <- utils.Result{Value: result, Source: "site-overview"}
			}
		}
	})
}
