package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gregoire-mullvad/git-carbon/config"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:                   "remove FILE",
	DisableFlagsInUseLine: true,
	Short:                 "Remove a carbon-copy from the current repository.",
	Long:                  `Remove the given file from .gitcarbon and the current repository.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := config.LoadFile(".gitcarbon")
		die(err)
		if len(args) != 1 {
			return errors.New("Missing FILE")
		}
		path := args[0]
		_, ok := conf.CCs[path]
		if !ok {
			die(fmt.Errorf("the following file is not a carbon copy:\n\t%s", path))
		}
		if !*removeFlags.keep && exists(path) {
			if !*removeFlags.force && isDirty(path) {
				die(fmt.Errorf(
					"the following file has local modifications:\n\t%s\n(use -k to keep the file, or -f to force removal)",
					path))
			}
			die(os.Remove(path))
			die(stage(path))
		}
		delete(conf.CCs, path)
		if len(conf.CCs) > 0 {
			die(conf.SaveFile(config.FileName))
		} else {
			die(os.Remove(config.FileName))
		}
		die(stage(config.FileName))
		return nil
	},
}

var removeFlags struct {
	force *bool
	keep  *bool
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeFlags.force = removeCmd.Flags().BoolP("force", "f", false, "Remove the file even if there are local modifications")
	removeFlags.keep = removeCmd.Flags().BoolP("keep", "k", false, "Remove the file from .gitcarbon but keep the local copy")
}
