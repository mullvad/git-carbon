/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/gregoire-mullvad/git-carbon/config"
	"github.com/spf13/cobra"
)

var allFlag *bool

func init() {
	rootCmd.AddCommand(updateCmd)
	allFlag = updateCmd.Flags().BoolP("all", "a", false, "Update all files git-carbon knows about.")
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:                   "update [--all] FILE...",
	DisableFlagsInUseLine: true,
	Short:                 "Update carbon copies from their respective repository.",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.Load()
		die(err)
		paths := args
		if *allFlag {
			paths = make([]string, 0, len(conf.CCs))
			for name := range conf.CCs {
				paths = append(paths, name)
			}
			sort.Strings(paths)
		}
		for _, p := range paths {
			cc := conf.CCs[p]
			fmt.Fprintf(os.Stderr, "Updating %s from %s\n", p, cc.SourceRepository)
			src, err := getSourceFile(p, cc.SourceRepository)
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
