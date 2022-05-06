/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/gregoire-mullvad/git-carbon/config"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:                   "add [--force] FILE REPOSITORY",
	DisableFlagsInUseLine: true,
	Short:                 "Add a carbon-copy to the current repository.",
	Long: `Copy the given file from a git remote to the current repository.

The path and repository are recorded in .gitcarbon so the file content can
easily be updated later.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.Load()
		die(err)
		dstp := args[0]
		url := args[1]
		if _, err := os.Stat(dstp); err == nil {
			// File exists
			if !*forceFlag {
				fmt.Fprintf(os.Stderr, "Error: %s already exists\n", dstp)
				os.Exit(1)
			}
		}
		src, err := getSourceFile(dstp, url)
		die(err)
		dst, err := os.Create(dstp)
		die(err)
		defer dst.Close()
		io.Copy(dst, src)
		conf.CCs[dstp] = config.CC{SourceRepository: url}
		err = conf.Save()
		die(err)
		err = stage(dstp)
		die(err)
		err = stage(config.FileName)
		die(err)
	},
}

var forceFlag *bool

func init() {
	rootCmd.AddCommand(addCmd)
	forceFlag = addCmd.Flags().BoolP("force", "f", false, "Add file even if it already exist in the repository")
}
