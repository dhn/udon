// Copyright (c) 2023 dhn. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import "github.com/projectdiscovery/gologger"

const banner = `

  88   88 8888b.   dP"Yb  88b 88 
  88   88  8I  Yb dP   Yb 88Yb88 
  Y8   8P  8I  dY Yb   dP 88 Y88 
   YbodP' 8888Y"   YbodP  88  Y8                                                   

`

// Version is the current version of dnsx
const Version = `0.0.1`

// showBanner is used to show the banner to the user
func ShowBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\t\tmade with <3\n\n")
}
