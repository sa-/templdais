// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package templdais

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type PaginationAttrs struct {
	Total       int
	PagesToShow int
	CurrentPage int
	LastPage    int
	NextPageUrl string
	PrevPageUrl string
}

func showPage(i, currentPage, total, pagesToShow int) bool {
	if i == 1 || i == total {
		return true
	}

	if i >= currentPage-pagesToShow && i <= currentPage+pagesToShow {
		return true
	}

	return false

}

func Pagination(pg PaginationAttrs, btn ButtonAttrs) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"join\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := 1; i <= pg.Total; i++ {
			if i == pg.CurrentPage {
				templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
					templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
					if !templ_7745c5c3_IsBuffer {
						templ_7745c5c3_Buffer = templ.GetBuffer()
						defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("i")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					if !templ_7745c5c3_IsBuffer {
						_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
					}
					return templ_7745c5c3_Err
				})
				templ_7745c5c3_Err = Button(btn.SetActive(), templ.Attributes{}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				if showPage(i, pg.CurrentPage, pg.Total, pg.PagesToShow) {
					templ_7745c5c3_Var3 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
						templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
						if !templ_7745c5c3_IsBuffer {
							templ_7745c5c3_Buffer = templ.GetBuffer()
							defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("i")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						if !templ_7745c5c3_IsBuffer {
							_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
						}
						return templ_7745c5c3_Err
					})
					templ_7745c5c3_Err = Button(btn, templ.Attributes{}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
