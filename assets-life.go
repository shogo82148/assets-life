// Copyright (C) 2019 Ichinose Shogo All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// This program generates filesystem.go and assets-life.go itself.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	pkgpath "path"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) <= 2 {
		log.Println("Usage:")
		log.Println(os.Args[0] + " INPUT_DIR OUTPUT_DIR [PACKAGE_NAME]")
		os.Exit(2)
	}
	in, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	out, err := filepath.Abs(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	var name string
	if len(os.Args) > 3 {
		name = os.Args[3]
	}
	if name == "" {
		name = filepath.Base(out)
	}
	if err := build(in, out, name); err != nil {
		log.Fatal(err)
	}
}

func build(in, out, name string) error {
	filename := "assets-life.go"
	rel, err := filepath.Rel(out, in)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(out, 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(out, "filesystem.go"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	header := `// Copyright (C) 2019 Ichinose Shogo All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
	
// Code generated by go run %s. DO NOT EDIT.

//%s

package %s

import (
	"io"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

// Root is the root of the file system.
var Root http.FileSystem = fileSystem{
`
	fmt.Fprintf(f, header, filename, "go:generate go run "+filename+" \""+rel+"\" . "+name, name)
	err = filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
		// ignore hidden files
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		fmt.Fprintf(f, "\tfile{\n")
		if info.IsDir() {
			fmt.Fprintln(f, "\t\tcontent: \"\",")
		} else {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			fmt.Fprintf(f, "\t\tcontent: %q,\n", string(b))
		}
		rel, err := filepath.Rel(in, path)
		if err != nil {
			return err
		}
		fmt.Fprintf(f, "\t\tname:    %q,\n", pkgpath.Clean("/"+filepath.ToSlash(rel)))
		mode := info.Mode()
		fmt.Fprintf(f, "\t\tmode:    0%03o", int(mode.Perm()))
		if mode.IsDir() {
			fmt.Fprintf(f, " | os.ModeDir")
		}
		fmt.Fprint(f, ",\n")
		fmt.Fprint(f, "\t},\n")
		return nil
	})
	if err != nil {
		return err
	}
	footer := `}

type fileSystem []file

func (fs fileSystem) Open(name string) (http.File, error) {
	i := sort.Search(len(fs), func(i int) bool { return fs[i].name >= name })
	if i >= len(fs) || fs[i].name != name {
		return nil, os.ErrNotExist
	}
	f := &fs[i]
	return &httpFile{
		Reader: strings.NewReader(f.content),
		file:   f,
		fs:     fs,
		idx:    i,
		dirIdx: i + 1,
	}, nil
}

type file struct {
	name    string
	content string
	mode    os.FileMode
}

var _ os.FileInfo = (*file)(nil)

func (f *file) Name() string {
	return path.Base(f.name)
}

func (f *file) Size() int64 {
	return int64(len(f.content))
}

func (f *file) Mode() os.FileMode {
	return f.mode
}

var zeroTime time.Time

func (f *file) ModTime() time.Time {
	return zeroTime
}

func (f *file) IsDir() bool {
	return f.Mode().IsDir()
}

func (f *file) Sys() interface{} {
	return nil
}

type httpFile struct {
	*strings.Reader
	file   *file
	fs     fileSystem
	idx    int
	dirIdx int
}

var _ http.File = (*httpFile)(nil)

func (f *httpFile) Stat() (os.FileInfo, error) {
	return f.file, nil
}

func (f *httpFile) Readdir(count int) ([]os.FileInfo, error) {
	ret := []os.FileInfo{}
	if !f.file.IsDir() {
		return ret, nil
	}

	prefix := f.file.name + "/"
	if count <= 0 {
		for ; f.dirIdx < len(f.fs); f.dirIdx++ {
			name := f.fs[f.dirIdx].name
			if !strings.HasPrefix(name, prefix) {
				break
			}
			if idx := strings.IndexRune(name[len(prefix):], '/'); idx >= 0 {
				continue
			}
			ret = append(ret, &f.fs[f.dirIdx])
		}
		return ret, nil
	}

	ret = make([]os.FileInfo, 0, count)
	for ; f.dirIdx < len(f.fs); f.dirIdx++ {
		name := f.fs[f.dirIdx].name
		if !strings.HasPrefix(name, prefix) {
			return ret, io.EOF
		}
		if idx := strings.IndexRune(name[len(prefix):], '/'); idx >= 0 {
			continue
		}
		ret = append(ret, &f.fs[f.dirIdx])
		if len(ret) == count {
			return ret, nil
		}
	}
	return ret, io.EOF
}

func (f *httpFile) Close() error {
	return nil
}`
	fmt.Fprintln(f, footer)
	if err := f.Close(); err != nil {
		return err
	}

	f, err = os.OpenFile(filepath.Join(out, filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	format := `// Copyright (C) 2019 Ichinose Shogo All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build ignore

// This program generates filesystem.go and assets-life.go itself.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	pkgpath "path"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) <= 2 {
		log.Println("Usage:")
		log.Println(os.Args[0] + " INPUT_DIR OUTPUT_DIR [PACKAGE_NAME]")
		os.Exit(2)
	}
	in, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	out, err := filepath.Abs(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	var name string
	if len(os.Args) > 3 {
		name = os.Args[3]
	}
	if name == "" {
		name = filepath.Base(out)
	}
	if err := build(in, out, name); err != nil {
		log.Fatal(err)
	}
}

func build(in, out, name string) error {
	filename := "assets-life.go"
	rel, err := filepath.Rel(out, in)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(out, 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(out, "filesystem.go"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	header := %c%s%c
	fmt.Fprintf(f, header, filename, "go:generate go run "+filename+" \""+rel+"\" . "+name, name)
	err = filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
		// ignore hidden files
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		fmt.Fprintf(f, "\tfile{\n")
		if info.IsDir() {
			fmt.Fprintln(f, "\t\tcontent: \"\",")
		} else {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			fmt.Fprintf(f, "\t\tcontent: %%q,\n", string(b))
		}
		rel, err := filepath.Rel(in, path)
		if err != nil {
			return err
		}
		fmt.Fprintf(f, "\t\tname:    %%q,\n", pkgpath.Clean("/"+filepath.ToSlash(rel)))
		mode := info.Mode()
		fmt.Fprintf(f, "\t\tmode:    0%%03o", int(mode.Perm()))
		if mode.IsDir() {
			fmt.Fprintf(f, " | os.ModeDir")
		}
		fmt.Fprint(f, ",\n")
		fmt.Fprint(f, "\t},\n")
		return nil
	})
	if err != nil {
		return err
	}
	footer := %c%s%c
	fmt.Fprintln(f, footer)
	if err := f.Close(); err != nil {
		return err
	}

	f, err = os.OpenFile(filepath.Join(out, filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	format := %c%s%c
	fmt.Fprintf(f, format, 96, header, 96, 96, footer, 96, 96, format, 96)
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
`
	fmt.Fprintf(f, format, 96, header, 96, 96, footer, 96, 96, format, 96)
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
