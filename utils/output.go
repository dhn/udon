// Copyright (c) 2023 dhn. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/projectdiscovery/gologger"
)

// JSON object domain
type jsonDomain struct {
	Domain string `json:"domain"`
	Source string `json:"source"`
}

// Print results as JSON
func WriteJSON(results <-chan Result) {
	encoder := jsoniter.NewEncoder(os.Stdout)
	var domain jsonDomain

	for result := range results {
		domain.Domain = result.Value
		domain.Source = result.Source
		err := encoder.Encode(&domain)
		if err != nil {
			gologger.Fatal().Msgf(err.Error())
		}
	}
}
