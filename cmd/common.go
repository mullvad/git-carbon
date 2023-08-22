package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	git "github.com/go-git/go-git/v5"
)

func stage(path string) error {
	r, err := git.PlainOpen(".")
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	_, err = w.Add(path)
	return err
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}

func isDirty(path string) bool {
	r, err := git.PlainOpen(".")
	die(err)
	w, err := r.Worktree()
	die(err)
	status, err := w.Status()
	die(err)
	fileStatus, ok := status[path]
	if !ok {
		// File is not in the status map, it's either clean or doesn't exists.
		return false
	}
	return fileStatus.Staging != git.Unmodified || fileStatus.Worktree != git.Unmodified
}

func die(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
}
