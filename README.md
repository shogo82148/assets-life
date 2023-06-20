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

The assets-life command is no longer needed because it is embedded into the generated package.
