// Code generated by templ@(devel) DO NOT EDIT.

package testahref

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func render() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<a href=\"javascript:alert(&#39;unaffected&#39;);\">")
		if err != nil {
			return err
		}
		var_2 := `Ignored`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><a href=\"")
		if err != nil {
			return err
		}
		var var_3 templ.SafeURL = templ.URL("javascript:alert('should be sanitized')")
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_3)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_4 := `Sanitized`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><a href=\"")
		if err != nil {
			return err
		}
		var var_5 templ.SafeURL = templ.SafeURL("javascript:alert('should not be sanitized')")
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_5)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_6 := `Unsanitized`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
