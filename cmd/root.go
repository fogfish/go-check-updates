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
	"strings"

	"github.com/fogfish/go-check-updates/internal/service"
	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		e := err.Error()
		fmt.Println(strings.ToUpper(e[:1]) + e[1:])
		os.Exit(1)
	}
}

var (
	rootUpdate bool
	rootPath   string

	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	rootCmd.Flags().BoolVarP(&rootUpdate, "update", "u", false, "update go.mod")
	rootCmd.Flags().StringVar(&rootPath, "path", ".", "path to module")
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
	RunE:    root,
	Version: fmt.Sprintf("go-check-updates/%s (%s), %s", version, commit, date),
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

	mod, err := service.Check(rootPath)
	bar.Finish()

	if err != nil {
		return err
	}

	if len(mod) == 0 {
		os.Stdout.WriteString("\n✅ go.mod is up to date.\n")
		return nil
	}

	service.ShowWithColor(os.Stdout, mod)
	os.Stdout.WriteString("\n run `go-check-updates -u` to upgrade go.mod\n")

	return nil
}

func update(cmd *cobra.Command, args []string) error {
	bar := progress()
	bar.Describe("checking go.mod")

	mod, err := service.Check(rootPath)
	if err != nil {
		return err
	}

	bar.Finish()

	if len(mod) == 0 {
		os.Stdout.WriteString("\n✅ go.mod is up to date.\n")
		return nil
	}

	for _, m := range mod {
		err := service.Update(rootPath, m)
		if err != nil {
			return err
		}
	}

	os.Stdout.WriteString("\n run `go mod tidy`\n")

	return nil
}
