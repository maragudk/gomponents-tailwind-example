package views

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type Props struct {
	Title   string
	Path    string
	Content g.Node
}

func Page(p Props) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    p.Title,
		Language: "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("/static/styles/app.css"), Type("text/css")),
		},
		Body: []g.Node{
			Nav(Class("bg-gray-700"),
				Container(
					Div(Class("flex items-baseline space-x-4"),
						NavbarLink("/", "Home", p.Path),
						NavbarLink("/about", "About", p.Path),
					),
				),
			),
			Container(
				Main(p.Content),
			),
		},
	})
}

// Container restricts the width of the children, centers, and adds some padding.
func Container(children ...g.Node) g.Node {
	return Div(Class("max-w-7xl mx-auto px-4"), g.Group(children))
}
