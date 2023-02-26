package project

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
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

    splitPath := strings.Split(path, "/")
    newPath := strings.Join(splitPath[1:], "/")
    newPath = fmt.Sprintf("%s/%s", target, newPath)

    if d.IsDir() {
        return os.Mkdir(newPath, os.ModePerm)
    }

    file, err := fs.ReadFile(example, path)

    if err != nil {
        return err
    }

    return os.WriteFile(newPath, file, 0644)
}
