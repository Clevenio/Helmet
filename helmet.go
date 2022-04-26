// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"github.com/clevenio/helmet/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	cmd.Version = version
	cmd.Commit = commit
	cmd.Date = date
	cmd.BuiltBy = builtBy

	cmd.Execute()
}
