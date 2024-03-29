package homedir_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/chonla/homedir"
	"github.com/stretchr/testify/assert"
)

func TestCreatingHomeDirSuccess(t *testing.T) {
	isCalled := false
	homedir.GetHomeDir = func() (string, error) {
		return "/some/path", nil
	}
	homedir.MakeDir = func(string, os.FileMode) error {
		isCalled = true
		return nil
	}

	h, e := homedir.NewHomeDir("user")

	assert.Nil(t, e)
	assert.NotNil(t, h)
	assert.True(t, isCalled)
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

func TestCreatingHomeDirFailedFromUnableToGetUserDir(t *testing.T) {
	isCalled := false

	homedir.GetHomeDir = func() (string, error) {
		return "", errors.New("some error")
	}
	homedir.MakeDir = func(string, os.FileMode) error {
		isCalled = true
		return nil
	}

	_, e := homedir.NewHomeDir("user")

	assert.NotNil(t, e)
	assert.Equal(t, "some error", e.Error())
	assert.False(t, isCalled)
}

func TestCreatingHomeDirShouldSuccessIfHomeDirExists(t *testing.T) {
	homedir.GetHomeDir = func() (string, error) {
		return "/some/path", nil
	}
	homedir.MakeDir = func(string, os.FileMode) error {
		return os.ErrExist
	}

	h, e := homedir.NewHomeDir("user")

	assert.Nil(t, e)
	assert.Equal(t, fmt.Sprintf("/some/path%suser", string(os.PathSeparator)), h.Path())
}

func TestCreatingHomeDirFailedFromUnableToMakeDir(t *testing.T) {
	homedir.GetHomeDir = func() (string, error) {
		return "/some/path", nil
	}
	homedir.MakeDir = func(string, os.FileMode) error {
		return errors.New("sorry bro!")
	}

	_, e := homedir.NewHomeDir("user")

	assert.NotNil(t, e)
	assert.Equal(t, "sorry bro!", e.Error())
}

func TestGettingItemInsideHomeDirWithoutSlashPrefixed(t *testing.T) {
	homedir.GetHomeDir = func() (string, error) {
		return "/some/path", nil
	}
	homedir.MakeDir = func(string, os.FileMode) error {
		return nil
	}

	h, _ := homedir.NewHomeDir("user")

	p := h.With("any/path")

	assert.Equal(t, fmt.Sprintf("/some/path%suser/any/path", string(os.PathSeparator)), p)
}

func TestGettingItemInsideHomeDirWithSlashPrefixed(t *testing.T) {
	homedir.GetHomeDir = func() (string, error) {
		return "/some/path", nil
	}
	homedir.MakeDir = func(string, os.FileMode) error {
		return nil
	}

	h, _ := homedir.NewHomeDir("user")

	p := h.With("/any/path")

	assert.Equal(t, fmt.Sprintf("/some/path%suser/any/path", string(os.PathSeparator)), p)
}

func TestGettingItemInsideHomeDirWithManySlashesPrefixed(t *testing.T) {
	homedir.GetHomeDir = func() (string, error) {
		return "/some/path", nil
	}
	homedir.MakeDir = func(string, os.FileMode) error {
		return nil
	}

	h, _ := homedir.NewHomeDir("user")

	p := h.With("//////any/path")

	assert.Equal(t, fmt.Sprintf("/some/path%suser/any/path", string(os.PathSeparator)), p)
}
