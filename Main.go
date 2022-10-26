package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed gambar.jpg
var logo []byte

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	var filesArray []string = []string{
		"a", "b", "c",
	}

	for _, v := range filesArray {
		x, _ := files.ReadFile(fmt.Sprintf("files/%s.txt", v))
		fmt.Println(string(x))
	}
}
