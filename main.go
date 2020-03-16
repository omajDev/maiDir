package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func sortRecursive(path string, info os.FileInfo) error {
	if info.IsDir() || filepath.HasPrefix(info.Name(), ".") {
		return filepath.SkipDir
	}

	handlDir := func(path, ext string) string {
		if ext != "" {
			return filepath.Dir(path) + "/" + ext[1:]
		}
		return filepath.Dir(path) + "/others"
	}

	ext := strings.ToLower(filepath.Ext(path))
	newDir := handlDir(path, ext)
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		os.Mkdir(newDir, 0755)
	}
	newPath := newDir + "/" + info.Name()
	os.Rename(path, newPath)
	return nil
}

func main() {

	dir := flag.String("dir", ".", "Set Directory to be sortted")
	flag.Parse()
	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		sortRecursive(*dir+"/"+f.Name(), f)
	}

}
