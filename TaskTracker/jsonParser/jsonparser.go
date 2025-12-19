package jsonparser

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kaushikmak/go-projects/TaskTracker/jsonParser/lexer"
	"github.com/kaushikmak/go-projects/TaskTracker/jsonParser/parser"
)

func JsonParser() {
	// A complex JSON string for testing
	input := `[{
		"name": "Nit Ki Gali",
		"users": 1050,
		"active": true,
		"features": ["chat", "anon", "video"],
		"meta": {
			"version": 1.2,
			"server": "vnit-local"
		},
		"bugs": null
	},
	{
		"name": "kaushik",
		"users": 1050,
		"active": true,
		"features": ["chat", "anon", "video"],
		"meta": {
			"version": 1.2,
			"server": "vnit-local"
		},
		"bugs": null
	}
	]`

	fmt.Println("--- Raw Input ---")
	fmt.Println(input)

	// 1. Lexical Analysis
	l := lexer.New(input)

	// 2. Syntactic Analysis
	p := parser.New(l)
	result := p.Parse()

	// Check for errors
	if len(p.Errors()) != 0 {
		log.Fatalf("Parser errors: %v", p.Errors())
	}

	fmt.Println("\n--- Parsed Output (Go Map) ---")
	fmt.Printf("%+v\n", result)

	// Verify by marshaling back to JSON using standard library
	fmt.Println("\n--- Verified Output (Standard Lib Marshal) ---")
	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
}
