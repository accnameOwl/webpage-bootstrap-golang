package renderHTML

import (
	"net/http"
	"text/template"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}
	rnd = renderer.New(opts)
}

var renderHTMLOpts = renderer.Options{ParseGlobPattern: "./tpl/*.html"}

//RenderHTMLCardGithub => renders html code that translates to bootstrap > card
func RenderHTMLCardGithub(response http.ResponseWriter, request *http.Request) {
	rnd.HTML(response, http.StatusOK, "card", nil)
}

//CardGithub holds data relative to
type Card struct {
	HeaderText    string
	SubheaderText string
	ParagraphText string
	LinkOneText   string
	LinkOneURL    string
	LinkTwoText   string
	LinkTwoURL    string
}

// generate github card as template
// card, err := template.ParseFiles("./tpl/card.html")

//RenderCard => output datapoint from a card, as html template
func (card *Card) RenderCard() {
	tmpl := template.Must(template.ParseFiles("tpl/card.html"))

}
