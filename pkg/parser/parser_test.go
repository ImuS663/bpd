package parser

import (
	"testing"
)

const xpath = "//*[@id=\"example\"]/div[1]"

func TestNewParserNotNil(t *testing.T) {
	parser := NewParser(xpath)

	if parser == nil {
		t.Error("parser is nil")
	}
}

func TestNewParserXpath(t *testing.T) {
	parser := NewParser(xpath)

	if parser.xpath != xpath {
		t.Errorf("expected %s, got %s", xpath, parser.xpath)
	}
}

func TestParseFileURLNoError(t *testing.T) {
	parser := NewParser("/html/body/div/p[2]/a")

	_, err := parser.ParseFileURL("https://example.com")
	if err != nil {
		t.Error(err)
	}
}

func TestParseFileURLWithError(t *testing.T) {
	parser := NewParser("/not/found/xpath")

	_, err := parser.ParseFileURL("https://example.com")
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestParseFileURL(t *testing.T) {
	parser := NewParser("/html/body/div/p[2]/a")
	expectedResult := "https://www.iana.org/domains/example"

	result, _ := parser.ParseFileURL("https://example.com")

	if result != expectedResult {
		t.Errorf("expected %s, got %s", expectedResult, result)
	}
}
