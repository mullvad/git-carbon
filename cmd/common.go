package cmd

import (
	"io"
	"log"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
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

func getSourceFile(path string, url string) (io.Reader, error) {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		Depth: 1,
		URL:   url,
	})
	if err != nil {
		return nil, err
	}
	ref, err := r.Head()
	if err != nil {
		return nil, err
	}
	c, err := r.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}
	f, err := c.File(path)
	if err != nil {
		return nil, err
	}
	return f.Blob.Reader()
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
