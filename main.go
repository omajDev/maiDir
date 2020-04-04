package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	maidir "maidir/maider"
	"os"
)

func prompt(msg string) string {
	var c string
	for c != "y" && c != "n" {
		fmt.Print(msg)
		fmt.Scanf("%s", &c)
	}
	return c
}
func main() {
	var (
		dir = flag.String("dir", ".", "Set Directory to be sorted")
		y   = flag.Bool("y", false, "answer y for every prompt")
	)
	flag.Parse()

	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		path := *dir + "/" + f.Name()
		newDir, err := maidir.NewPath(path, f)
		maidir.CreateDir(newDir)
		newPath := newDir + "/" + f.Name()
		if err != nil {
			continue
		}
		if *y {
			fmt.Println("moving file", path, " to ", newPath)
			os.Rename(path, newPath)
			continue
		}
		if prompt(fmt.Sprint("are you sure you want to move ", path, " to ", newPath, " (y/n)? ")) == "y" {
			fmt.Println("moving file", path, " to ", newPath)
			os.Rename(path, newPath)
			continue
		}
	}

}
