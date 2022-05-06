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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:                   "update [--all] FILE...",
	DisableFlagsInUseLine: true,
	Short:                 "Update carbon copies from their respective repository.",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.Load()
		die(err)
		for _, p := range args {
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

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolP("all", "a", false, "Update all files git-carbon knows about.")
}
