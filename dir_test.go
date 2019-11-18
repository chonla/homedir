package homedir_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/chonla/homedir"
	"github.com/stretchr/testify/assert"
)

func TestCreatingHomeDirSuccess(t *testing.T) {
	homedir.GetHomeDir = func() (string, error) {
		return "/some/path", nil
	}
	homedir.MakeDir = func(string, os.FileMode) error {
		return nil
	}

	h, e := homedir.NewHomeDir("user")

	assert.Nil(t, e)
	assert.NotNil(t, h)
}

func TestGettingPathFromHomeDir(t *testing.T) {
	homedir.GetHomeDir = func() (string, error) {
		return "/some/path", nil
	}
	homedir.MakeDir = func(string, os.FileMode) error {
		return nil
	}

	h, _ := homedir.NewHomeDir("user")

	assert.Equal(t, fmt.Sprintf("/some/path%suser", string(os.PathSeparator)), h.Path())
}
