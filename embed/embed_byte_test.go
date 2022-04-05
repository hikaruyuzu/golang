package embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// embed file kedalam bentuk []byte , bisa digunakan untuk embed data gambar, audio bahkan video

//go:embed kaffu.png
var chino []byte

func TestEmbedByte(t *testing.T) {
	err := ioutil.WriteFile("chino.png", chino, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(chino)
}
