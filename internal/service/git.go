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
	"os"
	"os/exec"
	"strings"

	"github.com/fogfish/go-check-updates/internal/types"
)

func GitBranch(dir string) error {
	git := exec.Command(
		"git", "checkout", "-b", types.UniqueBranchName,
	)
	git.Dir = dir
	git.Stderr, git.Stdout = os.Stderr, os.Stdout

	err := git.Run()
	if err != nil {
		return err
	}

	return nil
}

func GitUnBranch(dir string) error {
	git := exec.Command(
		"git", "checkout", "main",
	)
	git.Dir = dir
	git.Stderr, git.Stdout = os.Stderr, os.Stdout

	err := git.Run()
	if err != nil {
		return err
	}

	git = exec.Command(
		"git", "branch", "-D", types.UniqueBranchName,
	)
	git.Dir = dir
	git.Stderr, git.Stdout = os.Stderr, os.Stdout

	err = git.Run()
	if err != nil {
		return err
	}

	return nil
}

func GitAdd(dir string) error {
	git := exec.Command(
		"git", "add", "go.mod", "go.sum",
	)
	git.Dir = dir
	git.Stderr, git.Stdout = os.Stderr, os.Stdout

	err := git.Run()
	if err != nil {
		return err
	}

	return nil
}

func GitCommit(dir string, units []types.Unit) error {
	//
	//
	acc := 0
	for _, unit := range units {
		acc += len(unit.Mod)
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Update go.mod (%d modules)\n\n", acc))

	sb.WriteString("```")
	for _, unit := range units {
		sb.WriteString(
			fmt.Sprintf("\n%s\n", unit.Path),
		)
		for _, m := range unit.Mod {
			sb.WriteString(
				fmt.Sprintf("%s %s ⇒ %s\n", m.Path, m.Version, m.Upgrade),
			)
		}
	}
	sb.WriteString("```")

	git := exec.Command(
		"git", "commit", "-m", sb.String(),
	)
	git.Dir = dir
	git.Stderr, git.Stdout = os.Stderr, os.Stdout

	err := git.Run()
	if err != nil {
		return err
	}

	return nil
}

func GitPush(dir string, origin string) error {
	git := exec.Command(
		"git", "push", origin, "-u", types.UniqueBranchName,
	)
	git.Dir = dir
	git.Stderr, git.Stdout = os.Stderr, os.Stdout

	err := git.Run()
	if err != nil {
		return err
	}

	return nil
}
