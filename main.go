//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/go-check-updates
//

package main

import (
	"fmt"

	"github.com/fogfish/go-check-updates/cmd"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	cmd.Execute(fmt.Sprintf("go-check-updates/%s (%s), %s", version, commit[:7], date))
}
