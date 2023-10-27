//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/go-check-updates
//

package service

import (
	"os"
	"os/exec"

	"github.com/fogfish/go-check-updates/internal/types"
)

func Update(dir string, mod types.Mod) error {
	gcu := exec.Command(
		"go", "get",
		"-d", mod.Path+"@"+mod.Upgrade.String(),
	)
	gcu.Dir = dir
	gcu.Stderr, gcu.Stdout = os.Stderr, os.Stdout

	err := gcu.Run()
	if err != nil {
		return err
	}

	return nil
}
