package config

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"sort"

	format "github.com/go-git/go-git/v5/plumbing/format/config"
)

var (
	FileName            = ".gitcarbon"
	carbonSection       = "carbon"
	sourceRepositoryKey = "sourceRepository"
	sourceRefKey        = "sourceRef"
	sourcePathKey       = "sourcePath"
)

type Config struct {
	CCs map[string]CC
	raw *format.Config
}

type CC struct {
	Path             string
	SourceRepository string
	SourceRef        string
	SourcePath       string
}

func New() *Config {
	return &Config{
		CCs: make(map[string]CC),
		raw: format.New(),
	}
}

func LoadFile(name string) (*Config, error) {
	config := New()
	f, err := os.Open(FileName)
	if errors.Is(err, fs.ErrNotExist) {
		return config, nil
	} else if err != nil {
		return nil, err
	}
	defer f.Close()
	return Load(f)
}

func Load(r io.Reader) (*Config, error) {
	config := New()
	err := format.NewDecoder(r).Decode(config.raw)
	for _, ss := range config.raw.Section(carbonSection).Subsections {
		config.CCs[ss.Name] = CC{
			SourceRepository: ss.Option(sourceRepositoryKey),
			SourceRef:        ss.Option(sourceRefKey),
			SourcePath:       ss.Option(sourcePathKey),
		}
	}
	return config, err
}

func (c *Config) SaveFile(name string) error {
	f, err := os.Create(FileName)
	if err != nil {
		return err
	}
	defer f.Close()
	return c.Save(f)
}

func (c *Config) Save(w io.Writer) error {
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
		cc := c.CCs[name]
		ss.AddOption(sourceRepositoryKey, cc.SourceRepository)
		if cc.SourcePath != "" {
			ss.AddOption(sourcePathKey, cc.SourcePath)
		}
		if cc.SourceRef != "" {
			ss.AddOption(sourceRefKey, cc.SourceRef)
		}
	}
	s.Subsections = subsections
	return format.NewEncoder(w).Encode(c.raw)
}
