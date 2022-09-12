package cmd

import (
	"log"

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

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
