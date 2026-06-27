# go-output

CLI-agnostic result encoding (package `output`): `Format`, the JSON/YAML encoder registry, and `Write(w, format, data)`. It takes a writer, a format, and a value, so a daemon, a CLI, or a test all reuse the same encoding. It owns the one error intrinsic to its machinery, `ErrUnsupportedFormat`, declared on `go-error`.

- Depends on `go-error` and `gopkg.in/yaml.v3` only. Must not import a CLI framework.
- Quality gate: gofumpt, `go vet`, staticcheck, govulncheck, gocognit ≤ 7, **100% coverage**.
