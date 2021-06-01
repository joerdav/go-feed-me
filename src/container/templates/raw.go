package templates

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/yosssi/gohtml"
)

func Raw(text string) (t templ.Component) {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, text)
		return err
	})
}

func RawTemplate(c templ.Component) (t templ.Component) {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		var buf bytes.Buffer
		c.Render(ctx, &buf)
		s := gohtml.Format(buf.String())
		_, err = io.WriteString(w, templ.EscapeString(s))
		return err
	})
}
