package main

import (
	"fmt"
	"strings"
	"github.com/html_link_parser/parser"
)

func main() {
	html_string1 := `
	<html>
		<body>
			<p>
				<a href="nfl.com">nfl link</a>
				<p>
					<a href="espn.com">espn link</a>
				</p>
			</p>
			<a href="yahoo.com">yahoo link</a>
		</body>
	</html>`
	html_string2 := `<p>Links:</p><ul><li><!--<a href="comment_link">Comment Link</a>--><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	html_string3 := `<a href="/dog"><span>Something in a span</span>Text not in a span<b>Bold text!</b></a>`

	a_tags, _ := parser.ProcessHTML(strings.NewReader(html_string1))
	fmt.Println(a_tags)

	a_tags, _ = parser.ProcessHTML(strings.NewReader(html_string2))
	fmt.Println(a_tags)

	a_tags, _ = parser.ProcessHTML(strings.NewReader(html_string3))
	fmt.Println(a_tags)
}