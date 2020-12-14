package views

import (
	"github.com/go-chi/chi"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Home(routes []chi.Route) g.Node {
	return Div(
		Headline("Home"),
		P(
			g.Text("This project shows how to use "),
			TextLink("https://github.com/maragudk/gomponents", "gomponents"),
			g.Text(" together with "),
			TextLink("https://tailwindcss.com", "TailwindCSS"),
			g.Text(". Check out the introduction blog post: "),
			TextLink("https://www.maragu.dk/blog/building-view-components-with-gomponents-and-tailwindcss-in-go/",
				"Building view components with gomponents and TailwindCSS in Go"),
			g.Text("."),
		),

		SubHeadline("Routes declared in this app"),
		Ul(
			g.Map(len(routes), func(i int) g.Node {
				return Li(g.Text(routes[i].Pattern))
			})...,
		),

		SubHeadline("Buttons 😎"),
		Div(Class("max-w-lg flex space-x-8"),
			NiceButton("Click me!", true),
			NiceButton("Please don't click me…", false),
		),
	)
}

func NiceButton(text string, primary bool) g.Node {
	return Button(g.Text(text), c.Classes{
		"flex items-center justify-center px-4 py-3 rounded-md": true,
		"text-white bg-indigo-600 hover:bg-indigo-700":          primary,
		"text-indigo-600 bg-gray-50 hover:bg-gray-200":          !primary,
	})
}
