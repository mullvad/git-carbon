package config_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/mullvad/git-carbon/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadSave(t *testing.T) {
	tests := []struct {
		desc    string
		content string
		ccs     map[string]config.CC
	}{
		{
			"Minimal",
			`[carbon "foo.txt"]
	sourceRepository = git@example.com:test.git
`,
			map[string]config.CC{"foo.txt": {SourceRepository: "git@example.com:test.git"}},
		},
		{
			"WithSourcePath",
			`[carbon "foo.txt"]
	sourceRepository = git@example.com:test.git
	sourcePath = bar.txt
`,
			map[string]config.CC{"foo.txt": {SourceRepository: "git@example.com:test.git", SourcePath: "bar.txt"}},
		},
		{
			"WithSourceRef",
			`[carbon "foo.txt"]
	sourceRepository = git@example.com:test.git
	sourceRef = refs/tags/v1.2.3
`,
			map[string]config.CC{"foo.txt": {SourceRepository: "git@example.com:test.git", SourceRef: "refs/tags/v1.2.3"}},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Load/%s", test.desc), func(t *testing.T) {
			src := strings.NewReader(test.content)
			actual, err := config.Load(src)
			assert.NoError(t, err)
			assert.Equal(t, test.ccs, actual.CCs)
		})
		t.Run(fmt.Sprintf("Save/%s", test.desc), func(t *testing.T) {
			conf := config.New()
			conf.CCs = test.ccs
			buf := bytes.NewBufferString("")
			err := conf.Save(buf)
			assert.NoError(t, err)
			assert.Equal(t, test.content, buf.String())
		})
	}
}
