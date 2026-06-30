package output_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

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
