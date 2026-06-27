# go-output

Encode any value to an `io.Writer` as JSON or YAML through a small format registry, with an `ErrUnsupportedFormat` sentinel for unknown formats.

## Install

```sh
go get github.com/gomatic/go-output
```

## Usage

```go
package main

import (
	"os"

	output "github.com/gomatic/go-output"
)

func main() {
	data := map[string]any{"hello": "world"}
	if err := output.Write(os.Stdout, output.FormatJSON, data); err != nil {
		panic(err)
	}
}
```

`Write` defaults to JSON when the format is empty and returns `ErrUnsupportedFormat` for any unknown format.
