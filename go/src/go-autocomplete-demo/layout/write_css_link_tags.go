package layout

import "io"

func WriteCssLinkTags(w io.Writer, hrefs ...string) (err error) {
	for _, href := range hrefs {
		if len(href) > 0 {
			if _, err = fprintf(w, `<link rel="stylesheet" type="text/css" href=%q></link>`, href); err != nil {
				return
			}
		}
	}
	return
}
