package goembed

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	files, _ := path.ReadDir("files")

	for _, entry := range files {
		if !entry.IsDir() {
			fmt.Print(entry.Name())
			file, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				t.Errorf("Error reading file: %v", err)
			}
			fmt.Println(string(file))
		}
	}
}
