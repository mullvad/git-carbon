/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/mullvad/git-carbon/config"
	"github.com/mullvad/git-carbon/git"
	"github.com/spf13/cobra"
)

var updateFlags struct {
	all    *bool
	quiet  *bool
	branch *string
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateFlags.all = updateCmd.Flags().BoolP("all", "a", false, "Update all files git-carbon knows about.")
	updateFlags.quiet = updateCmd.Flags().BoolP("quiet", "q", false, "Suppress output.")
	updateFlags.branch = updateCmd.Flags().String("branch", "", "Override the branch to update from.")
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:                   "update [--all] FILE...",
	DisableFlagsInUseLine: true,
	Short:                 "Update carbon copies from their respective repository.",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.LoadFile(".gitcarbon")
		die(err)
		paths := args
		if *updateFlags.all {
			paths = make([]string, 0, len(conf.CCs))
			for name := range conf.CCs {
				paths = append(paths, name)
			}
			sort.Strings(paths)
		}
		gitClient := &git.Client{Quiet: *updateFlags.quiet}
		for _, p := range paths {
			cc := conf.CCs[p]
			if !*updateFlags.quiet {
				fmt.Fprintf(os.Stderr, "Updating %s from %s\n", p, cc.SourceRepository)
			}
			srcp := cc.SourcePath
			if srcp == "" {
				srcp = p
			}
			branch := *updateFlags.branch
			if branch == "" {
				branch = cc.SourceRef
			}
			src, err := gitClient.GetSourceFile(srcp, cc.SourceRepository, branch)
			die(err)
			dst, err := os.Create(p)
			die(err)
			defer dst.Close()
			_, err = io.Copy(dst, src)
			die(err)
			err = stage(p)
			die(err)
		}
	},
}
