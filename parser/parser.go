package parser

import (
	"io"
	"strings"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

type Links []Link

 func ProcessHTML(r io.Reader) (Links, error) {
	 doc, err := html.Parse(r)
	 if err != nil {
		 return nil, err
	 }

	 return search(doc), nil
 }

// type Node struct {
	// Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	// Type      NodeType
	// DataAtom  atom.Atom
	// Data      string
	// Namespace string
	// Attr      []Attribute
// }

func search(n *html.Node) Links {
	var links Links	
	stack := []*html.Node{n}

	counter := 0
	var nd *html.Node

	for len(stack) > 0 {
		stack, nd = popStack(stack)

		if nd.Type == html.ElementNode {
			if nd.Data == "a" {
				newLink := Link{}
				newLink.Href = grabHref(nd)
				newLink.Text = grabText(nd)
				links = append(links, newLink)
			}
		}
		for c := nd.LastChild; c != nil; c = c.PrevSibling {
			stack = append(stack, c)
		}

		counter += 1
	}

	return links
}

func grabHref(nd *html.Node) string {
	for _, a := range nd.Attr {
		if a.Key == "href" {
			return a.Val
		}
	}

	return ""
}

func grabText(n *html.Node) string {
	var sb strings.Builder

	a_stack := []*html.Node{n}
	var nd *html.Node

	for len(a_stack) > 0 {
		a_stack, nd = popStack(a_stack)
		
		if nd.Type == html.TextNode {
			sb.WriteString(nd.Data)
			sb.WriteString(" ")
		}
		for ac := nd.LastChild; ac != nil; ac = ac.PrevSibling {
			a_stack = append(a_stack, ac)
		}
	}

	return strings.TrimRight(sb.String(), " ")
}

func popStack(stack []*html.Node) ([]*html.Node, *html.Node) {
	if len(stack) == 0 {
		return stack, nil
	}

	popped := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	return stack, popped
}