/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gregoire-mullvad/git-carbon/config"
	"github.com/gregoire-mullvad/git-carbon/git"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:                   "add [--force] [--ref REF] REPOSITORY FILE [DESTINATION]",
	DisableFlagsInUseLine: true,
	Short:                 "Add a carbon-copy to the current repository.",
	Long: `Copy the given file from a git remote to the current repository.

The path and repository are recorded in .gitcarbon so the file content can
easily be updated later.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.LoadFile(".gitcarbon")
		die(err)
		url := args[0]
		srcp := args[1]
		var dstp string
		if len(args) == 3 {
			dstp = args[2]
		} else if len(args) == 2 {
			dstp = srcp
		} else {
			die(errors.New("Wrong number of arguments"))
		}
		if _, err := os.Stat(dstp); err == nil {
			// File exists
			if !*addFlags.force {
				fmt.Fprintf(os.Stderr, "Error: %s already exists\n", dstp)
				os.Exit(1)
			}
		}
		src, err := (&git.Client{Quiet: *addFlags.quiet}).GetSourceFile(srcp, url, *addFlags.ref)
		die(err)
		dst, err := os.Create(dstp)
		die(err)
		defer dst.Close()
		io.Copy(dst, src)
		cc := config.CC{SourceRepository: url}
		if srcp != dstp {
			cc.SourcePath = srcp
		}
		if *addFlags.ref != "" {
			cc.SourceRef = *addFlags.ref
		}
		conf.CCs[dstp] = cc
		err = conf.SaveFile(".gitcarbon")
		die(err)
		err = stage(dstp)
		die(err)
		err = stage(config.FileName)
		die(err)
	},
}

var addFlags struct {
	force *bool
	quiet *bool
	ref   *string
}

func init() {
	rootCmd.AddCommand(addCmd)
	addFlags.force = addCmd.Flags().BoolP("force", "f", false, "Add file even if it already exist in the repository")
	addFlags.quiet = addCmd.Flags().BoolP("quiet", "q", false, "Suppress output")
	addFlags.ref = addCmd.Flags().StringP("ref", "r", "", "Ref of source repository")
}
