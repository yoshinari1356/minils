package main
import (
	"fmt"
	"io"
	"io/fs"
	"os"
)
type osFS struct{}
func (fsys *osFS) Open(name string) (fs.File, error) {
	f, err := os.Open(name)
	if f == nil {
		return nil, err
	}
	return f, err
}
func (fsys *osFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}
func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func run() error {
	fsys := new(osFS)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	return listDir(fsys, wd, os.Stdout)
}
func listDir(fsys fs.ReadDirFS, dir string, out io.Writer) error {
	entries, err := fsys.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		_, err := out.Write([]byte(fmt.Sprintf("%s\n", e.Name())))
		if err != nil {
			return err
		}
	}
	return nil
}
