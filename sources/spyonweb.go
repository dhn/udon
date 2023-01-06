// Copyright (c) 2023 dhn. All rights reserved.
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

// GetSPData function returns all domains based on the given UA id name
func GetSPData(ua string) <-chan utils.Result {
	results := make(chan utils.Result)

	go func() {
		getSPData(fmt.Sprintf("https://spyonweb.com/%s",
			url.QueryEscape(ua)), ua, results)
		close(results)
	}()

	return results
}

// Send a HTTP request and parse the JSON response
func getSPData(sourceURL string, ua string, results chan utils.Result) {
	resp := utils.GetHTTPRequest(sourceURL, map[string]string{})

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return
	}

	doc.Find("a").Each(func(i int, domain *goquery.Selection) {
		for _, result := range strings.Split(domain.Text(), "\n") {
			value, _ := domain.Attr("href")
			if strings.Contains(value, "/go/") {
				results <- utils.Result{Value: result, Source: "spyonweb"}
			}
		}
	})
}
