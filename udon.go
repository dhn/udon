// Copyright (c) 2023 dhn. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/dhn/udon/runner"
	"github.com/dhn/udon/utils"
)

func main() {
	options := utils.ParseOptions()

	if options.SearchString != "" {
		runner.UA(options.SearchString, *options)
	}
}
