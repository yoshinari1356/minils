package main
import (
	"bytes"
	"io/fs"
	"testing"
	"testing/fstest"
)
func TestListDir(t *testing.T) {
	tests := []struct {
		dir  string
		want string
	}{
		{"path", "to\n"},
		{"path/to", "a.txt\nb.txt\ngo\n"},
		{"path/to/go", "c.txt\n"},
	}
	for _, tt := range tests {
		buf := new(bytes.Buffer)
		if err := listDir(testFS(), tt.dir, buf); err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		if got != tt.want {
			t.Errorf("got %s\nwant %s", got, tt.want)
		}
	}
}
func testFS() fstest.MapFS {
	fsys := fstest.MapFS{
		"path":             &fstest.MapFile{Mode: fs.ModeDir},
		"path/to":          &fstest.MapFile{Mode: fs.ModeDir},
		"path/to/a.txt":    &fstest.MapFile{Data: []byte("test\n")},
		"path/to/b.txt":    &fstest.MapFile{Data: []byte("test\n")},
		"path/to/go":       &fstest.MapFile{Mode: fs.ModeDir},
		"path/to/go/c.txt": &fstest.MapFile{Data: []byte("test\n")},
	}
	return fsys
}
