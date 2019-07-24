package deep

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test(t *testing.T) {
	t.Run("/a", func(t *testing.T) {
		f, err := Root.Open("/a")
		if err != nil {
			t.Fatal(err)
		}
		stat, err := f.Stat()
		if err != nil {
			t.Fatal(err)
		}
		if stat.Name() != "a" {
			t.Errorf("unexpected name: want %q, got %q", "a", stat.Name())
		}
		if stat.IsDir() {
			t.Error("want not directory, but it is")
		}
		if stat.Size() != 0 {
			t.Errorf("unexpected size: want %d, got %d", 0, stat.Size())
		}

		if stat.Mode() != 0644 {
			t.Errorf("unexpected mode: want %s, got %s", os.FileMode(0644), stat.Mode())
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		if string(b) != "" {
			t.Errorf("unexpected content: want empty, got %v", string(b))
		}
	})

	t.Run("/aa/bb/c", func(t *testing.T) {
		f, err := Root.Open("/aa/bb/c")
		if err != nil {
			t.Fatal(err)
		}
		stat, err := f.Stat()
		if err != nil {
			t.Fatal(err)
		}
		if stat.Name() != "c" {
			t.Errorf("unexpected name: want %q, got %q", "c", stat.Name())
		}
		if stat.IsDir() {
			t.Error("want not directory, but it is")
		}
		if stat.Size() != 0 {
			t.Errorf("unexpected size: want %d, got %d", 0, stat.Size())
		}

		if stat.Mode() != 0644 {
			t.Errorf("unexpected mode: want %s, got %s", os.FileMode(0644), stat.Mode())
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		if string(b) != "" {
			t.Errorf("unexpected content: want empty, got %v", string(b))
		}
	})

	t.Run("/", func(t *testing.T) {
		f, err := Root.Open("/")
		if err != nil {
			t.Fatal(err)
		}
		stat, err := f.Stat()
		if err != nil {
			t.Fatal(err)
		}
		if stat.Name() != "/" {
			t.Errorf("unexpected name: want %q, got %q", "/", stat.Name())
		}
		if !stat.IsDir() {
			t.Error("want directory, but it is not")
		}

		// group and others' permissions are depends on the environment.
		// so, we check only owner's permission.
		if stat.Mode()&0700 != 0700 {
			t.Errorf("unexpected mode: want %s, got %s", os.FileMode(0700), stat.Mode()&0700)
		}
	})

	t.Run("/aa", func(t *testing.T) {
		f, err := Root.Open("/aa")
		if err != nil {
			t.Fatal(err)
		}
		stat, err := f.Stat()
		if err != nil {
			t.Fatal(err)
		}
		if stat.Name() != "aa" {
			t.Errorf("unexpected name: want %q, got %q", "aa", stat.Name())
		}
		if !stat.IsDir() {
			t.Error("want directory, but it is not")
		}

		if stat.Mode() != 0755|os.ModeDir {
			t.Errorf("unexpected mode: want %s, got %s", 0755|os.ModeDir, stat.Mode())
		}
	})

	t.Run("/aa/bb", func(t *testing.T) {
		f, err := Root.Open("/aa/bb")
		if err != nil {
			t.Fatal(err)
		}
		stat, err := f.Stat()
		if err != nil {
			t.Fatal(err)
		}
		if stat.Name() != "bb" {
			t.Errorf("unexpected name: want %q, got %q", "bb", stat.Name())
		}
		if !stat.IsDir() {
			t.Error("want directory, but it is not")
		}

		if stat.Mode() != 0755|os.ModeDir {
			t.Errorf("unexpected mode: want %s, got %s", 0755|os.ModeDir, stat.Mode())
		}
	})
}
