// SPDX-FileCopyrightText: 2026 Mullvad VPN AB
// SPDX-License-Identifier: GPL-3.0-only

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-carbon",
	Short: "Add & update files from other repositories.",
	Long: `git-carbon is a tool that helps manage duplicated files across repositories.

Like git-submodule but for individual files.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
