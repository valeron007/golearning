package main

import (
	"fmt"
	"golang.org/x/net/html"
	"math"
	"net/http"
	"os"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

type Node struct {
	Type                     NodeType
	Data                     string
	Attr                     []Attribute
	FirstChild, NextSibiling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

//func Parse(r io.Reader) (*Node, error)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	url := "https://www.goland.org/"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))

	//doc, err := html.Parse(os.Stdin)
	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}
