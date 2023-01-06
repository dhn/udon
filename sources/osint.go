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

// GetOSData function returns all domains based on the given UA id name
func GetOSData(ua string) <-chan utils.Result {
	results := make(chan utils.Result)

	go func() {
		getOSData("https://osint.sh/analytics/", ua, results)
		close(results)
	}()

	return results
}

// Send a HTTP request and parse the JSON response
func getOSData(sourceURL string, ua string, results chan utils.Result) {
	data := []byte(fmt.Sprintf("code=%s", url.QueryEscape(ua)))
	resp := utils.PostHTTPRequest(sourceURL, data)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return
	}

	doc.Find("td").Each(func(i int, domain *goquery.Selection) {
		for _, result := range strings.Split(domain.Text(), "\n") {
			value, _ := domain.Attr("data-th")
			if value == "Domain" && len(strings.TrimSpace(result)) != 0 {
				results <- utils.Result{Value: strings.TrimSpace(result), Source: "osint.sh"}
			}
		}
	})
}
