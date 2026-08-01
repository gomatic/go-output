package output_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	output "github.com/gomatic/go-output"
)

type sample struct {
	Name  string `json:"name"  yaml:"name"`
	Count int    `json:"count" yaml:"count"`
}

func TestWrite(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		format   output.Format
		contains []string
	}{
		{"json", output.FormatJSON, []string{`"name": "bucket"`, `"count": 3`}},
		{"yaml", output.FormatYAML, []string{"name: bucket", "count: 3"}},
		{"empty defaults to json", "", []string{`"name": "bucket"`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			want := assert.New(t)
			var buf bytes.Buffer
			want.NoError(output.Write(&buf, tt.format, sample{Name: "bucket", Count: 3}))
			for _, sub := range tt.contains {
				want.Contains(buf.String(), sub)
			}
		})
	}
}

func TestWriteJSONNoHTMLEscape(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	assert.New(t).NoError(output.Write(&buf, output.FormatJSON, map[string]string{"url": "a&b<c"}))
	assert.New(t).Contains(buf.String(), "a&b<c")
}

func TestWriteUnsupportedFormat(t *testing.T) {
	t.Parallel()
	want := assert.New(t)
	err := output.Write(&bytes.Buffer{}, "xml", sample{})
	want.ErrorIs(err, output.ErrUnsupportedFormat)
	want.True(strings.Contains(err.Error(), "xml"))
}

// TestErrUnsupportedFormatIsMatchableAndDistinct names ErrUnsupportedFormat's
// claim. It is declared as an errs.Const so a caller matches it with errors.Is
// rather than by message text — a caller branching on the string breaks the
// moment the wording improves, silently, with nothing to warn them. An
// unsupported format is a caller mistake they can recover from (fall back to a
// default format); an encoding failure is not, so the two must stay
// distinguishable.
func TestErrUnsupportedFormatIsMatchableAndDistinct(t *testing.T) {
	t.Parallel()

	err := output.Write(io.Discard, output.Format("no-such-format"), map[string]string{"a": "b"})

	require.ErrorIs(t, err, output.ErrUnsupportedFormat)
	assert.NotErrorIs(t, errors.New("some other failure"), output.ErrUnsupportedFormat,
		"an unrelated error must not be mistaken for an unsupported format")

	wrapped := output.ErrUnsupportedFormat.With(errors.New("cause"), "yaml")
	assert.ErrorIs(t, wrapped, output.ErrUnsupportedFormat,
		"wrapping must preserve the sentinel, or a caller's recovery path stops firing")
}
