// Copyright (c) 2023 dhn. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runner

import (
	"github.com/dhn/udon/sources"
	"github.com/dhn/udon/utils"
)

// Query several sources to reverse search
func UA(id string, options utils.Options) {
	ua := utils.MergeChannels(
		sources.GetHDData(id),
		sources.GetSOData(id),
		sources.GetOSData(id),
		sources.GetSPData(id),
	)
	ua = utils.RemoveDuplicates(ua)
	utils.PrintResults(options.JSON, ua)
}
