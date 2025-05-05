package goembed

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed files/version.txt
var version string

//go:embed favicon.ico
var logo []byte

//go:embed files/*.txt
var files embed.FS

func main() {
	fmt.Println("Version:", version)
	fmt.Println("Logo size:", len(logo))

	paths, _ := files.ReadDir("files")
	for _, path := range paths {
		fmt.Println("File:", path.Name())
	}
}
