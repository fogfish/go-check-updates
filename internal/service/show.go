//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/go-check-updates
//

package service

import (
	"fmt"
	"io"

	"github.com/fogfish/go-check-updates/internal/types"
)

func ShowAll(w io.Writer, units []types.Unit) {
	for _, u := range units {
		w.Write(
			[]byte(
				fmt.Sprintf("\n%s\n", u.Path),
			),
		)

		Show(w, u.Mod)
	}
}

func Show(w io.Writer, mod []types.Mod) {
	for _, m := range mod {
		w.Write(
			[]byte(
				fmt.Sprintf("%-12s ⇒ %-12s\t%s\n", m.Version, m.Upgrade, m.Path),
			),
		)
	}
}

func ShowAllWithColor(w io.Writer, units []types.Unit) {
	for _, u := range units {
		w.Write(
			[]byte(
				fmt.Sprintf("\n\033[1m%s\033[0m\n", u.Path),
			),
		)

		Show(w, u.Mod)
	}
}

func ShowWithColor(w io.Writer, mod []types.Mod) {
	for _, m := range mod {
		diff := m.Version.Diff(m.Upgrade)

		var prefix, suffix string

		switch {
		case diff.Major != "":
			prefix = fmt.Sprintf("\033[31m%s\033[0m", m.Upgrade)
		case diff.Minor != "":
			prefix = m.Version.Major
			suffix = fmt.Sprintf("\033[36m%s\033[0m", diff)
		case diff.Patch != "":
			prefix = m.Version.Major
			if m.Version.Minor != "" {
				prefix = prefix + "." + m.Version.Minor
			}
			suffix = fmt.Sprintf("\033[32m%s\033[0m", diff)
		}

		w.Write(
			[]byte(
				fmt.Sprintf("%-12s ⇒ %-12s\t%s\n", m.Version, prefix+suffix, m.Path),
			),
		)
	}
}
