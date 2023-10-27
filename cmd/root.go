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

var rootUpdate bool

func init() {
	rootCmd.Flags().BoolVarP(&rootUpdate, "update", "u", false, "update go.mod")
}

var rootCmd = &cobra.Command{
	Use:   "go-check-updates",
	Short: "xxx",
	Long: `
xxx
	`,
	RunE:    root,
	Version: "go-check-updates/v0.0.0",
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

	mod, err := service.Check("")
	if err != nil {
		return err
	}

	bar.Finish()

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

	mod, err := service.Check("")
	if err != nil {
		return err
	}

	bar.Finish()

	if len(mod) == 0 {
		os.Stdout.WriteString("\n✅ go.mod is up to date.\n")
		return nil
	}

	for _, m := range mod {
		err := service.Update("", m)
		if err != nil {
			return err
		}
	}

	os.Stdout.WriteString("\n run `go mod tidy`\n")

	return nil
}
