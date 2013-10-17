package layout

import "io"

func WriteDataPool(w io.Writer, v interface{}) error {
	return WriteIsland(w, "data-pool", v)
}
