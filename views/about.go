package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func About() g.Node {
	return Div(
		Headline("About"),
		P(
			TextLink("https://github.com/maragudk/gomponents", "gomponents"),
			g.Text(" is a project by "),
			TextLink("https://www.maragu.dk", "maragu"),
			g.Text("."),
		),
	)
}
