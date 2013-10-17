package layout

import "io"

func WriteScriptTags(w io.Writer, srcs ...string) (err error) {
	for _, src := range srcs {
		if len(src) > 0 {
			if _, err = fprintf(w, `<script src=%q></script>`, src); err != nil {
				return
			}
		}
	}
	return
}
