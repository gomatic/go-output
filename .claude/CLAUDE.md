# go-output

CLI-agnostic result encoding (package `output`): `Format`, the JSON/YAML encoder registry, and `Write(w, format, data)`. Owns `ErrUnsupportedFormat`, declared in `errors.go` on `go-error`'s `error.Const` (a dedicated sentinel file so `output.go` keeps using the builtin `error`). Generic — lives in `gomatic`.

- Depends on `go-error` and `gopkg.in/yaml.v3` only. Must not import a CLI framework.
- Gate: gofumpt, vet, staticcheck, govulncheck, gocognit ≤ 7, 100% coverage.
