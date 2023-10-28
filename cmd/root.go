//
// Copyright (C) 2023 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/go-check-updates
//

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogfish/go-check-updates/internal/service"
	"github.com/fogfish/go-check-updates/internal/types"
	"github.com/spf13/cobra"
)

func Execute(vsn string) {
	rootCmd.Version = vsn

	if err := rootCmd.Execute(); err != nil {
		e := err.Error()
		fmt.Println(strings.ToUpper(e[:1]) + e[1:])
		os.Exit(1)
	}
}

var (
	rootUpdate    bool
	rootRecursive bool
	rootGitPush   string
	rootPath      string
)

func init() {
	rootCmd.Flags().BoolVarP(&rootUpdate, "update", "u", false, "update go.mod")
	rootCmd.Flags().BoolVarP(&rootRecursive, "recursive", "r", false, "update go.mod recursively")
	rootCmd.Flags().StringVar(&rootGitPush, "push", "", "push go.mod changes to git repository in branch "+types.UniqueBranchName)
	rootCmd.Flags().StringVar(&rootPath, "path", "."+string(filepath.Separator), "path to module")
}

var rootCmd = &cobra.Command{
	Use:   "go-check-updates",
	Short: "upgrades your go.mod dependencies to the latest versions",
	Long: `
Upgrades your go.mod dependencies to the latest versions.

The utility is wrapper for Golang's list command:

	go list -u \
	  -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' \
	  -m all

Check status of go.mod dependencies:

	go-check-updates

Upgrades your go.mod dependencies to the latest versions:

	go-check-updates -u

See more info https://github.com/fogfish/go-check-updates
	`,
	Example: `
	go-check-updates
	go-check-updates -u
	go-check-updates -u --push github
	`,
	SilenceUsage: true,
	RunE:         root,
}

func root(cmd *cobra.Command, args []string) error {
	if rootUpdate {
		return update(cmd, args)
	}

	return check(cmd, args)
}

func check(cmd *cobra.Command, args []string) error {
	bar := progress()
	bar.Describe("checking go.mod")

	mod, err := service.CheckAll(rootPath, rootRecursive)
	bar.Finish()

	if err != nil {
		return err
	}

	if len(mod) == 0 {
		os.Stdout.WriteString("\n✅ go.mod is up to date.\n")
		return nil
	}

	service.ShowAllWithColor(os.Stdout, mod)
	os.Stdout.WriteString("\nrun `go-check-updates -u --push origin` to upgrade go.mod and push changes to upstream repo\n")

	return nil
}

func update(cmd *cobra.Command, args []string) error {
	bar := progress()
	bar.Describe("checking go.mod")

	units, err := service.CheckAll(rootPath, rootRecursive)
	if err != nil {
		return err
	}

	bar.Finish()

	if len(units) == 0 {
		os.Stdout.WriteString("\n✅ go.mod is up to date.\n")
		return nil
	}

	if rootGitPush != "" {
		err = service.GitBranch(rootPath)
		if err != nil {
			return err
		}
	}

	for _, unit := range units {
		for _, m := range unit.Mod {
			err := service.Update(unit.Path, m)
			if err != nil {
				return err
			}
		}

		err = service.GoModTidy(unit.Path)
		if err != nil {
			return err
		}

		if rootGitPush != "" {
			err = service.GitAdd(unit.Path)
			if err != nil {
				return err
			}
		}
	}

	if rootGitPush != "" {
		err = service.GitCommit(rootPath, units)
		if err != nil {
			return err
		}

		err = service.GitPush(rootPath, rootGitPush)
		if err != nil {
			return err
		}

		service.GitUnBranch(rootPath)
	}

	return nil
}
