package readdir

import (
	"io"
	"testing"
)

func Test(t *testing.T) {
	// Readdir will return all children
	t.Run("Readdir(-1)", func(t *testing.T) {
		dir, err := Root.Open("/")
		if err != nil {
			t.Errorf("fs.Open(/) = %v", err)
			return
		}
		fis, err := dir.Readdir(-1)
		if err != nil {
			t.Errorf("dir.Readdir(-1) = %v", err)
			return
		}
		if len(fis) != 3 {
			t.Errorf("got: %d, expect: 3", len(fis))
		}
	})

	// Readdir will return all children
	t.Run("Readdir(0)", func(t *testing.T) {
		dir, err := Root.Open("/")
		if err != nil {
			t.Errorf("fs.Open(/) = %v", err)
			return
		}
		fis, err := dir.Readdir(0)
		if err != nil {
			t.Errorf("dir.Readdir(0) = %v", err)
			return
		}
		if len(fis) != 3 {
			t.Errorf("got: %d, expect: 3", len(fis))
		}
	})

	t.Run("Readdir(>0)", func(t *testing.T) {
		dir, err := Root.Open("/")
		if err != nil {
			t.Errorf("fs.Open(/) = %v", err)
			return
		}

		// Readdir will return 1 item
		fis, err := dir.Readdir(1)
		if err != nil {
			t.Errorf("dir.Readdir(1) = %v", err)
			return
		}
		if len(fis) != 1 {
			t.Errorf("got: %d, expect: 1", len(fis))
		}
		if fis[0].Name() != "aa" {
			t.Errorf("got: %s, expect: aa", fis[0].Name())
		}

		// Readdir will return 1 item
		fis, err = dir.Readdir(1)
		if err != nil {
			t.Errorf("dir.Readdir(1) = %v", err)
			return
		}
		if len(fis) != 1 {
			t.Errorf("got: %d, expect: 1", len(fis))
		}
		if fis[0].Name() != "bb" {
			t.Errorf("got: %s, expect: bb", fis[0].Name())
		}

		// take rest entries
		fis, err = dir.Readdir(-1)
		if err != nil {
			t.Errorf("dir.Readdir(1) = %v", err)
			return
		}
		if len(fis) != 1 {
			t.Errorf("got: %d, expect: 1", len(fis))
		}
		if fis[0].Name() != "cc" {
			t.Errorf("got: %s, expect: cc", fis[0].Name())
		}

		// try to take rest entries, but no entry will return
		// because all entries is already read
		fis, err = dir.Readdir(-1)
		if err != nil {
			t.Errorf("dir.Readdir(1) = %v", err)
			return
		}
		if len(fis) != 0 {
			t.Errorf("got: %d, expect: 0", len(fis))
		}

		// reach EOF
		fis, err = dir.Readdir(1)
		if err != io.EOF {
			t.Errorf("error should be io.EOF, but: %s", err)
			return
		}
		if len(fis) != 0 {
			t.Errorf("got: %d, expect: 0", len(fis))
		}
	})
}
