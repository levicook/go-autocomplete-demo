package layout

import (
	"encoding/json"
	"io"
)

func WriteIsland(w io.Writer, name string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = fprintf(w, `<script id=%q type="application/json">%s</script>`, name, data)

	return err
}
