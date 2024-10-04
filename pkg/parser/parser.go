package parser

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

type Parser struct {
	xpath string
}

// NewParser creates a new Parser instance with the given XPath expression.
func NewParser(xpath string) *Parser {
	return &Parser{xpath: xpath}
}

// ParseFileURL takes a URL and extracts the file URL from it by using the stored XPath expression.
// The function returns the extracted URL and an error if the XPath expression does not match any node
// or if the node does not have an href attribute.
func (p *Parser) ParseFileURL(url string) (string, error) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return "", err
	}

	node := htmlquery.FindOne(doc, p.xpath)
	if node == nil {
		return "", fmt.Errorf("node by xpath %s not found", p.xpath)
	}

	result := htmlquery.SelectAttr(node, "href")
	if result == "" {
		return "", fmt.Errorf("href attribute not found")
	}

	return result, nil
}
