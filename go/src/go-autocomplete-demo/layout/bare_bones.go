package layout

import (
	"bytes"
	"go-autocomplete-demo/fail"
	"html/template"
	"io"
)

const bareBonesHtml = `<!DOCTYPE html>
<html lang="en"><head>
<meta charset="utf-8">
<title>{{.Title}}</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
{{.HeadHtml}}
</head>
<body>
<div id="main">{{.MainHtml}}</div>
<div id="tail" style="position: absolute; top: -10000px; height: 0px; width: 0px;">
{{.TailHtml}}
</div>
</body>
</html>`

var bareBonesTemplate *template.Template

func init() {
	t, e := template.New("layout/BareBones").Parse(bareBonesHtml)
	fail.If(e)
	bareBonesTemplate = t
}

type BareBones struct {
	Title string
	Head  bytes.Buffer
	Main  bytes.Buffer
	Tail  bytes.Buffer
}

func (bb BareBones) HeadHtml() template.HTML {
	return template.HTML(bb.Head.String())
}

func (bb BareBones) MainHtml() template.HTML {
	return template.HTML(bb.Main.String())
}

func (bb BareBones) TailHtml() template.HTML {
	return template.HTML(bb.Tail.String())
}

func (bb BareBones) Render(w io.Writer) error {
	return bareBonesTemplate.Execute(w, bb)
}
