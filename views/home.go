package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Home() g.Node {
	return Div(
		Headline("Home"),
		P(
			g.Text("This project shows how to use "),
			TextLink("https://github.com/maragudk/gomponents", "gomponents"),
			g.Text(" together with "),
			TextLink("https://tailwindcss.com", "TailwindCSS"),
			g.Text(". Check out the introduction blog post: "),
			TextLink("https://www.maragu.dk/blog/using-gomponents-with-tailwindcss-in-go/", "Using gomponents with TailwindCSS in Go"),
			g.Text("."),
		),
	)
}
