package maidir

import (
	"os"
	"path/filepath"
	"strings"
)

func HandlDir(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	if ext != "" {
		return filepath.Dir(path) + "/" + ext[1:]
	}
	return filepath.Dir(path) + "/others"
}
func CreateDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
}
func NewPath(path string, info os.FileInfo) (string, error) {
	if info.IsDir() || filepath.HasPrefix(info.Name(), ".") {
		return "", filepath.SkipDir
	}

	return HandlDir(path), nil
}
