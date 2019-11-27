package homedir

import (
	"fmt"
	"os"
	"strings"
)

// GetHomeDir is alias for getting user home directory
var GetHomeDir = os.UserHomeDir

// MakeDir is alias for making directory
var MakeDir = os.Mkdir

// HomeWrapper is interface of HomeDir
type HomeWrapper interface {
	Path() string
	With(string) string
}

// HomeDir is home directory struct
type HomeDir struct {
	path string
}

// NewHomeDir creates homedir instance
func NewHomeDir(dirname string) (*HomeDir, error) {
	fullPath, e := ensure(dirname)
	if e != nil {
		return nil, e
	}
	return &HomeDir{
		path: fullPath,
	}, e
}

// ensure is called to ensure path existence in homedir. return full path
func ensure(dirname string) (string, error) {
	homePath, e := GetHomeDir()
	if e != nil {
		return "", e
	}

	fullDirName := fmt.Sprintf(`%s%s%s`, homePath, string(os.PathSeparator), dirname)

	e = MakeDir(fullDirName, 0755)
	if e != nil && !os.IsExist(e) {
		return "", e
	}
	return fullDirName, nil
}

// Path returns full path of user defined directory in home directory
func (h *HomeDir) Path() string {
	return h.path
}

// With returns sub path inside current home directory with user defined dir
func (h *HomeDir) With(sub string) string {
	sub = strings.TrimLeft(sub, string(os.PathSeparator))
	return fmt.Sprintf("%s%s%s", h.path, string(os.PathSeparator), sub)
}
