package git

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp/sideband"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Client struct {
	Quiet bool
	cache map[string]*git.Repository
}

func (c *Client) clone(url, refname string) (*git.Repository, error) {
	cacheKey := fmt.Sprintf("%s#%s", url, refname)
	if repo, ok := c.cache[cacheKey]; ok {
		return repo, nil
	}
	var progress sideband.Progress
	if !c.Quiet {
		progress = os.Stderr
	}
	opts := &git.CloneOptions{
		Depth:         1,
		URL:           url,
		ReferenceName: plumbing.ReferenceName(refname),
		Progress:      progress,
	}
	err := opts.Validate()
	if err != nil {
		return nil, err
	}
	repo, err := git.Clone(memory.NewStorage(), nil, opts)
	if err != nil {
		return nil, err
	}
	if c.cache == nil {
		c.cache = make(map[string]*git.Repository, 1)
	}
	c.cache[cacheKey] = repo
	return repo, nil
}

func (c *Client) GetSourceFile(path string, url string, refname string) (io.Reader, error) {
	r, err := c.clone(url, refname)
	if err != nil {
		return nil, err
	}
	ref, err := r.Head()
	if err != nil {
		return nil, err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}
	f, err := commit.File(path)
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
