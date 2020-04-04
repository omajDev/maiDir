package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		dir = flag.String("dir", "", "Set Directory to be sorted")
		y   = flag.Bool("y", false, "")
	)
	flag.Parse()
	if *dir == "" {
		flag.Usage()
		os.Exit(0)
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
		newPath, err := maidir.NewPath(path, f)
		maidir.CreateDir(newPath)
		if err != nil {
			log.Println(path, ": ", err)
			continue
		}
		if *y {
			log.Println("moving file", path, " to ", newPath)
			os.Rename(path, newPath)
			continue
		}

	}

}
