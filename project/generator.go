package project

import (
	"embed"
	"io/fs"
	"os"
	"regexp"
)

//go:embed example
var example embed.FS

var target string

func Generate(name string) error {
	target = name

	return fs.WalkDir(example, ".", createFiles)
}

func createFiles(path string, d fs.DirEntry, err error) error {
	if path == "." {
		return nil
	}

	newPath := regexp.MustCompile("^example").ReplaceAllString(path, target)

	if d.IsDir() {
		return os.Mkdir(newPath, os.ModePerm)
	}

	file, err := fs.ReadFile(example, path)

	if err != nil {
		return err
	}

	return os.WriteFile(newPath, file, 0644)
}
