package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"

	"strings"
)

//go:embed version.txt
var version string

//go:embed test/files/*
var path embed.FS

func main() {
	fmt.Println(version)
	dirEntries, _ := path.ReadDir("test/files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			var dirName string = entry.Name()
			if strings.HasSuffix(dirName, ".jpg") || strings.HasSuffix(dirName, ".jpeg") {
				frog, _ := ioutil.ReadFile("test/files/" + dirName)

				err := ioutil.WriteFile("newfolder/"+dirName+"_new.png", frog, fs.ModePerm)
				if err != nil {
					panic(err)
				}
			} else {
				path.ReadFile("test/files/" + dirName)
				fmt.Println(string(dirName))
			}

		}
	}
}
