package output

// Imported bare (the package is named error); this file declares only sentinels
// and uses no builtin error type, so the declaration reads error.Const.
import "github.com/gomatic/go-error"

// ErrUnsupportedFormat indicates a format the encoder cannot produce.
const ErrUnsupportedFormat error.Const = "unsupported output format"
