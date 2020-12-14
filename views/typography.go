package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Headline(text string) g.Node {
	return H1(text, Class("text-3xl font-bold text-gray-900 my-6"))
}

func SubHeadline(text string) g.Node {
	return H2(text, Class("text-2xl font-bold text-gray-900 my-6"))
}

func TextLink(href, text string) g.Node {
	return A(Href(href), g.Text(text), Class("text-blue-600 hover:text-blue-900 transition duration-150 ease-in-out"))
}
