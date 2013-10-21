package hello

import (
	"app/layout"
	"bytes"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	// specify our dataPool
	var dataPool struct{ X string }

	// populate dataPool
	dataPool.X = "x"

	// build head
	var head bytes.Buffer
	must(layout.WriteCssLinkTags(&head)) // TODO cssFor("pages/hello"),

	// build main
	var main bytes.Buffer
	main.WriteString(`<h1>Hello</h1>`)

	// build tail
	var tail bytes.Buffer
	must(layout.WriteDataPool(&tail, dataPool))
	must(layout.WriteScriptTags(&tail)) // TODO jsFor("pages/hello"),

	// build page
	page := layout.BareBones{
		Title: "Hello",
		Head:  head,
		Main:  main,
		Tail:  tail,
	}

	// render page
	must(page.Render(w))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
