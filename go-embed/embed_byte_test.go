package goembed

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed favicon.ico
var favicon []byte

func TestEmbedByte(t *testing.T) {
	err := ioutil.WriteFile("favicon_new.ico", favicon, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
