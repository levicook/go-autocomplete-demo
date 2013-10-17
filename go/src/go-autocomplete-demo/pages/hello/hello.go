package hello

import (
	"bytes"
	//"fmt"
	"go-autocomplete-demo/layout"
	//"go-autocomplete-demo/steal"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	// specify our dataPool
	var dataPool struct{ X string }

	// populate dataPool
	dataPool.X = "x"

	// build head
	var head bytes.Buffer
	must(layout.WriteCssLinkTags(&head,
		//steal.CssFor("app-js/pages/hello"),
	))

	// build main
	var main bytes.Buffer
	main.WriteString(`<h1>Hello</h1>`)

	// build tail
	var tail bytes.Buffer
	must(layout.WriteDataPool(&tail, dataPool))
	must(layout.WriteScriptTags(&tail,
		//steal.JsFor("app-js/pages/hello"),
	))

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
