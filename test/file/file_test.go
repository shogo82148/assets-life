package file

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	f, err := Root.Open("/file.txt")
	if err != nil {
		t.Fatal(err)
	}
	stat, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if stat.Name() != "file.txt" {
		t.Errorf("unexpected name: want %q, got %q", "file.txt", stat.Name())
	}
	if stat.IsDir() {
		t.Error("want not directory, but it is")
	}
	if stat.Size() != 0 {
		t.Errorf("unexpected size: want %d, got %d", 0, stat.Size())
	}

	// group and others' permissions are depends on the environment.
	// so, we check only owner's permission.
	if stat.Mode()&0700 != 0700 {
		t.Errorf("unexpected mode: want %s, got %s", os.FileMode(0700), stat.Mode()&0700)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "" {
		t.Errorf("unexpected content: want empty, got %v", string(b))
	}
}

func TestHiddenFile(t *testing.T) {
	_, err := Root.Open("/.hidden_file")
	if !os.IsNotExist(err) {
		t.Errorf("hidden file will be not exist, but %v", err)
	}
}
