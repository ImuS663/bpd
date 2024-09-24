package parser

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

type Parser struct {
	xpath string
}

func NewParser(xpath string) *Parser {
	return &Parser{xpath: xpath}
}

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
