package views

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func NavbarLink(path, text, currentPath string) g.Node {
	active := path == currentPath

	return A(Href(path),
		c.Classes{
			"p-4 text-sm font-medium focus:text-white": true,
			"text-white":                     active,
			"text-gray-300 hover:text-white": !active,
		},
		g.Text(text),
	)
}

func Headline(text string) g.Node {
	return H1(text, Class("text-3xl font-bold text-gray-900 my-6"))
}

func TextLink(href, text string) g.Node {
	return A(Href(href), g.Text(text), Class("text-blue-600 hover:text-blue-900 transition duration-150 ease-in-out"))
}
