package config

import (
	"errors"
	"io/fs"
	"os"

	format "github.com/go-git/go-git/v5/plumbing/format/config"
)

var (
	FileName = ".gitcarbon"
	section  = "cc"
)

type Config struct {
	raw *format.Config
}

type CC struct {
	Path             string
	SourceRepository string
}

func New() *Config {
	return &Config{
		raw: format.New(),
	}
}

func Load() (*Config, error) {
	config := New()
	f, err := os.Open(FileName)
	if errors.Is(err, fs.ErrNotExist) {
		return config, nil
	} else if err != nil {
		return nil, err
	}
	defer f.Close()
	err = format.NewDecoder(f).Decode(config.raw)
	return config, err
}

func (c *Config) Save() error {
	f, err := os.Create(FileName)
	if err != nil {
		return err
	}
	defer f.Close()
	err = format.NewEncoder(f).Encode(c.raw)
	return err
}

func (c *Config) Add(path string, sourceRepo string) {
	c.raw.AddOption(section, path, "sourceRepository", sourceRepo)
}

func (c *Config) Get(path string) *CC {
	sub := c.raw.Section(section).Subsection(path)
	return &CC{path, sub.Option("sourceRepository")}
}
