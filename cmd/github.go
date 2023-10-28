package cmd

import (
	"os"

	"github.com/fogfish/go-check-updates/internal/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(githubCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate workflows",
	Run:   generate,
}

func generate(cmd *cobra.Command, args []string) {
	cmd.Help()
}

var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "generate GitHub Action",
	Long: `
Generates GitHub Action to automate pull request creation when go-update-deps branch is created.
	`,
	Example: `
	go-check-updates generate github > .github/workflow/update-deps.yml
	`,
	SilenceUsage: true,
	RunE:         github,
}

func github(cmd *cobra.Command, args []string) error {
	_, err := os.Stdout.Write([]byte(githubAction()))
	return err
}

func githubAction() string {
	return `
##
## Open Pull Request: Update dependency for Go modules
##
## Note: add explicit following "push" trigger to other workflows that checks quality.
##       GitHub would not trigger actions automatically, see https://github.com/peter-evans/create-pull-request/blob/main/docs/concepts-guidelines.md#workarounds-to-trigger-further-workflow-runs
##
## Copy the following lines:
##  on:
##    push:
##      branches:
##        - ` + types.UniqueBranchName + `
##        - /refs/heads/` + types.UniqueBranchName + `
##
##
name: deps
on:
  push:
    branches:
      - ` + types.UniqueBranchName + `
      - /refs/heads/` + types.UniqueBranchName + `

jobs:
  deps:
    permissions:
      contents: read
      pull-requests: write

    runs-on: ubuntu-latest
    steps:

      - uses: actions/checkout@v4

      - uses: diillson/auto-pull-request@v1.0.1
        id: open-pull-request
        with:
          destination_branch: "main"
          source_branch: "` + types.UniqueBranchName + `"

      - uses: actions/labeler@v4
        with:
          pr-number: ${{steps.open-pull-request.outputs.pr_number}}
`
}
