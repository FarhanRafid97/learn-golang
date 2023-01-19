package test

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"

	"testing"
)

//go:embed files/a.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)

}

//go:embed files/frog.jpg
var frog []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("frog_new.png", frog, fs.ModePerm)
	if err != nil {
		panic(err)
	}

}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/frog.jpg
var files embed.FS

func TestMultipleFile(t *testing.T) {

	a, _ := files.ReadFile("files/a.txt")

	fmt.Println(string(a))
	b, _ := files.ReadFile("files/b.txt")

	fmt.Println(string(b))

	frog, _ := ioutil.ReadFile("files/frog.jpg")

	err := ioutil.WriteFile("files/frog_new_file.png", frog, fs.ModePerm)
	if err != nil {
		panic(err)
	}

}

//go:embed files/*
var path embed.FS

func TestPatchMatcher(t *testing.T) {

	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			var dirName string = entry.Name()
			if strings.HasSuffix(dirName, ".jpg") || strings.HasSuffix(dirName, ".jpeg") {
				frog, _ := ioutil.ReadFile("files/" + dirName)

				err := ioutil.WriteFile("files/"+dirName+"_new.png", frog, fs.ModePerm)
				if err != nil {
					panic(err)
				}
			} else {
				txt, _ := files.ReadFile("files/" + dirName)
				fmt.Println(string(txt))
			}

		}
	}
}
