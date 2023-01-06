// Copyright (c) 2023 dhn. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

// Options
type Options struct {
	SearchString string // UA id is the word to find for
	Silent       bool   // Silent suppresses any extra text and only writes the output to screen
	Stdin        bool   // Stdin specifies whether stdin input was given to the process
	Version      bool   // Version specifies if we should just show version and exit
	JSON         bool   // JSON output
	Output       io.Writer
}

// ParseOptions parses the command line flags provided by a user
func ParseOptions() *Options {
	options := &Options{}

	flag.StringVar(&options.SearchString, "s", "", "UA ID to find domains for")
	flag.BoolVar(&options.Silent, "silent", false, "Show only domains in output")
	flag.BoolVar(&options.Version, "version", false, "Show version of udon")
	flag.BoolVar(&options.JSON, "json", false, "Print results as JSON")
	flag.Parse()

	options.Output = os.Stdout
	options.Stdin = hasStdin()
	options.configureOutput()

	ShowBanner()

	if options.Version {
		gologger.Info().Msgf("Current Version: %s\n", Version)
		os.Exit(0)
	}

	// Validate the options passed by the user and if any
	// invalid options have been used, exit.
	err := options.validateOptions()
	if err != nil {
		gologger.Fatal().Msgf("Program exiting: %s\n", err)
	}

	return options
}

func hasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	isPipedFromChrDev := (stat.Mode() & os.ModeCharDevice) == 0
	isPipedFromFIFO := (stat.Mode() & os.ModeNamedPipe) != 0

	return isPipedFromChrDev || isPipedFromFIFO
}

// configureOutput configures the output on the screen
func (options *Options) configureOutput() {
	if options.Silent {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	}
}

// validateOptions validates the configuration options passed
func (options *Options) validateOptions() error {
	if len(strings.TrimSpace(options.SearchString)) == 0 {
		return errors.New("please set the UA ID with '-s'")
	}

	return nil
}
