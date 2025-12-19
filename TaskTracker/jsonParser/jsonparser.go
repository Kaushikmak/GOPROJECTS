package jsonparser

import (
	"log"

	"github.com/kaushikmak/go-projects/TaskTracker/jsonParser/lexer"
	"github.com/kaushikmak/go-projects/TaskTracker/jsonParser/parser"
)

// JsonParser takes a raw JSON string and returns the parsed data as an interface{}.
func JsonParser(input string) any {
	l := lexer.New(input)
	p := parser.New(l)
	result := p.Parse()

	if len(p.Errors()) != 0 {
		log.Fatalf("Parser errors: %v", p.Errors())
	}
	return result
}
