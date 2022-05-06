package config

import (
	"errors"
	"io/fs"
	"os"
	"sort"

	format "github.com/go-git/go-git/v5/plumbing/format/config"
)

var (
	FileName            = ".gitcarbon"
	carbonSection       = "carbon"
	sourceRepositoryKey = "sourceRepository"
)

type Config struct {
	CCs map[string]CC
	raw *format.Config
}

type CC struct {
	Path             string
	SourceRepository string
}

func New() *Config {
	return &Config{
		CCs: make(map[string]CC),
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
	for _, ss := range config.raw.Section(carbonSection).Subsections {
		config.CCs[ss.Name] = CC{SourceRepository: ss.Option(sourceRepositoryKey)}
	}
	return config, err
}

func (c *Config) Save() error {
	s := c.raw.Section(carbonSection)
	subsections := make(format.Subsections, 0, len(c.CCs))

	// Sort subsections by name so marshalling is deterministic
	names := make([]string, 0, len(c.CCs))
	for name := range c.CCs {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		ss := &format.Subsection{Name: name}
		subsections = append(subsections, ss)
		ss.AddOption(sourceRepositoryKey, c.CCs[name].SourceRepository)
	}
	s.Subsections = subsections
	f, err := os.Create(FileName)
	if err != nil {
		return err
	}
	defer f.Close()
	err = format.NewEncoder(f).Encode(c.raw)
	return err
}
