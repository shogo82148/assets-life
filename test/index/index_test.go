package index

import "testing"

func Test(t *testing.T) {
	dir, err := Root.Open("/")
	if err != nil {
		t.Fatal(err)
	}
	fis, err := dir.Readdir(0)
	if err != nil {
		t.Fatal(err)
	}

	if len(fis) != 2 {
		t.Errorf("want %d, got %d", 2, len(fis))
	}
	if fis[0].Name() != "index.html" {
		t.Errorf("want %q, got %q", "index.html", fis[0].Name())
	}
	if fis[1].Name() != "sub_dir" {
		t.Errorf("want %q, got %q", "sub_dir", fis[0].Name())
	}
}

func BenchmarkOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Root.Open("/index.html")
	}
}

func BenchmarkReaddir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dir, _ := Root.Open("/")
		dir.Readdir(0)
	}
}
