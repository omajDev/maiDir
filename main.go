package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func handlDir(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	if ext != "" {
		return filepath.Dir(path) + "/" + ext[1:]
	}
	return filepath.Dir(path) + "/others"
}
func createDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
}
func newPath(path string, info os.FileInfo) (string, error) {
	if info.IsDir() || filepath.HasPrefix(info.Name(), ".") {
		return "", filepath.SkipDir
	}

	return handlDir(path) + "/" + info.Name(), nil
}

func main() {

	dir := flag.String("dir", "", "Set Directory to be sortted")
	//y := flag.Bool("y", false, "")
	flag.Parse()
	if *dir == "" {
		flag.Usage()
		os.Exit(1)
	}
	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		path := *dir + "/" + f.Name()
		newPath, err := newPath(path, f)
		createDir(newPath)
		if err != nil {
			log.Println(path, ": ", err)
			continue
		}
		log.Println("moving file", path, " to ", newPath)
		os.Rename(path, newPath)

	}

}
