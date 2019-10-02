package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	a, err := ioutil.ReadFile("assets-life.go")
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadFile("test/file/assets-life.go")
	if err != nil {
		t.Fatal(err)
	}

	// normalize line feed marks
	a = bytes.Replace(a, []byte("\r\n"), []byte("\n"), -1)
	b = bytes.Replace(b, []byte("\r\n"), []byte("\n"), -1)

	// the first generation does not have the Build Constraints for "go get"
	b = bytes.Replace(b, []byte("\n// +build ignore\n"), []byte(""), 1)

	if string(a) != string(b) {
		t.Error("do not match", string(b))
	}
}
