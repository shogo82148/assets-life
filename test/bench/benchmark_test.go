package bench

import "testing"

func BenchmarkOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Root.Open("/00/00.txt")
	}
}

func BenchmarkReaddir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dir, _ := Root.Open("/")
		dir.Readdir(0)
	}
}
