package goembed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed gambar.jpg
var logo []byte

func TestBytes(t *testing.T) {
	err := ioutil.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	var filesArray []string = []string{
		"a", "b", "c",
	}

	for _, v := range filesArray {
		x, _ := files.ReadFile(fmt.Sprintf("files/%s.txt", v))
		fmt.Println(string(x))
	}
}
