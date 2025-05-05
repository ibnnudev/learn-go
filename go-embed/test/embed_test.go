package test

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

// go:embed files/version.txt
var version string

func TestEmbed(t *testing.T) {
	version, err := files.ReadFile("files/version.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Embedded version:", string(version))
}

//go:embed files/version.txt files/version2.txt
var files embed.FS

func TestEmbedFS(t *testing.T) {
	version, err := files.ReadFile("files/version.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Embedded version:", string(version))

	version2, err := files.ReadFile("files/version2.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Embedded version2:", string(version2))
}
