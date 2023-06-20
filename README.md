# ARCHIVED!

From Go 1.16, Go officially supports embedded files by using [the embed package](https://pkg.go.dev/embed).
Please use [the embed package](https://pkg.go.dev/embed) instead assets-life.
assets-life is no longer maintained.

[![Build Status](https://github.com/shogo82148/assets-life/workflows/Test/badge.svg)](https://github.com/shogo82148/assets-life/actions)

# assets-life

assets-life is a very simple embedding asset generator.
It generates an embed small in-memory file system that is served from an http.FileSystem.

## Usage

Install the command line tool first.

```
go install github.com/shogo82148/assets-life@latest
```

The assets-life command generates a package that have embed small in-memory file system.

```
assets-life /path/to/your/project/public public
```

You can access the file system by accessing a public variable `Root` of the generated package.

```go
import (
    "net/http"
    "./public" // TODO: Replace with the absolute import path
)

func main() {
    http.Handle("/", http.FileServer(public.Root))
    http.ListenAndServe(":8080", nil)
}
```

Visit http://localhost:8080/path/to/file to see your file.

The assets-life command also embed `go:generate` directive into the generated code, and assets-life itself.
It allows you to re-generate the package using go generate.

```
go generate ./public
```

The assets-life command is implemented as a [Quine (computing)](https://en.wikipedia.org/wiki/Quine_(computing))
and embedded into the generated package.
So the command is no longer needed.

## References

- [Goのバイナリに静的ファイルを埋め込むツール assets-life を書いた (written in Japanese)](https://shogo82148.github.io/blog/2019/07/24/assets-life/)
- [the embed package](https://pkg.go.dev/embed)
- [Quine (computing)](https://en.wikipedia.org/wiki/Quine_(computing))
