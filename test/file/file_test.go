package file

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test(t *testing.T) {
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
	if stat.Mode() != 0755 {
		t.Errorf("unexpected mode: want %s, got %s", os.FileMode(0755), stat.Mode())
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "" {
		t.Errorf("unexpected content: want empty, got %v", string(b))
	}
}
