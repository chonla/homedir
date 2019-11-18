package homedir

import (
	"fmt"
	"os"
)

// GetHomeDir is alias for getting user home directory
var GetHomeDir = os.UserHomeDir

// MakeDir is alias for making directory
var MakeDir = os.Mkdir

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

	e = MakeDir(fullDirName, 0666)
	if e != nil {
		return "", e
	}
	return fullDirName, nil
}

// Path returns full path of user defined directory in home directory
func (h *HomeDir) Path() string {
	return h.path
}
