package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

var debug = false

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	if len(os.Args) > 2 {
		debug = (os.Args[2] == "--debug" || os.Args[2] == "-d")
	}
	abs := path.Join(os.Getenv("GOPATH"), "src")
	if name != "" {
		match := &FileMatch{name: name, path: abs}
		err := filepath.Walk(abs, match.Find)
		if err != nil && err.Error() != "done" {
			fmt.Fprintf(os.Stderr, "** %s", err)
		}
		abs = match.path
	}
	fmt.Println(abs)
}

type FileMatch struct {
	name string
	path string
}

func (f *FileMatch) Find(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fname := info.Name()
	if fname == "_workspace" || fname == ".git" || fname == ".hg" || fname == ".bundle" {
		return filepath.SkipDir
	}
	m := info.Mode()
	if !m.IsDir() && m&os.ModeSymlink == 0 {
		return nil
	}
	if fname == f.name {
		f.path = path
		return fmt.Errorf("done")
	}
	if debug {
		fmt.Println(path + "/" + info.Name())
	}

	return nil
}
