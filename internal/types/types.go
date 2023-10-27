//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/go-check-updates
//

package types

import "strings"

// Version of package
type Vsn struct {
	Major, Minor, Patch string
}

func Version(s string) Vsn {
	vsn := Vsn{}
	seq := strings.Split(s, ".")

	if len(seq) >= 1 {
		vsn.Major = seq[0]
	}
	if len(seq) >= 2 {
		vsn.Minor = seq[1]
	}
	if len(seq) >= 3 {
		vsn.Patch = seq[2]
	}

	return vsn
}

func (vsn Vsn) String() string {
	sb := strings.Builder{}
	sb.WriteString(vsn.Major)

	if vsn.Minor != "" {
		sb.WriteRune('.')
		sb.WriteString(vsn.Minor)
	}

	if vsn.Patch != "" {
		sb.WriteRune('.')
		sb.WriteString(vsn.Patch)
	}

	return sb.String()
}

func (vsn Vsn) Diff(v Vsn) Vsn {
	if vsn.Major != v.Major {
		return v
	}

	if vsn.Minor != v.Minor {
		return Vsn{Minor: v.Minor, Patch: v.Patch}
	}

	if vsn.Patch != v.Patch {
		return Vsn{Patch: v.Patch}
	}

	return Vsn{}
}

type Mod struct {
	Path    string
	Version Vsn
	Upgrade Vsn
}
