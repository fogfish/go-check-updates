//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/go-check-updates
//

package service

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fogfish/go-check-updates/internal/types"
)

func Check(dir string) ([]types.Mod, error) {
	buf := &bytes.Buffer{}

	_, err := os.Stat(filepath.Join(dir, "go.mod"))
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("go.mod do not exists at %s", dir)
	}

	gcu := exec.Command(
		"go", "list",
		"-u",
		"-f", "{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}:{{.Version}}:{{.Update.Version}}{{end}}",
		"-m", "all",
	)
	gcu.Dir = dir
	gcu.Stderr, gcu.Stdout = os.Stderr, buf

	err = gcu.Run()
	if err != nil {
		return nil, err
	}

	mod := make([]types.Mod, 0)
	seq := strings.Split(buf.String(), "\n")
	for _, pkg := range seq {
		vsn := strings.Split(pkg, ":")
		if len(vsn) == 3 {
			path, vsn, upg := vsn[0], types.Version(vsn[1]), types.Version(vsn[2])
			mod = append(mod, types.Mod{
				Path:    path,
				Version: vsn,
				Upgrade: upg,
			})
		}
	}

	return mod, nil
}
