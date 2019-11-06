package homedir

import (
	"fmt"
	"log"
	"os"
)

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
	homePath, e := os.UserHomeDir()
	if e != nil {
		log.Fatal("unable to get home directory")
		return "", e
	}

	fullDirName := fmt.Sprintf(`%s%s%s`, homePath, string(os.PathSeparator), dirname)

	e = os.Mkdir(fullDirName, 0666)
	if e != nil {
		return "", e
	}
	return fullDirName, nil
}
