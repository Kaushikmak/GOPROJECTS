package parser

import (
	"fmt"
	"strconv"

	"github.com/kaushikmak/go-projects/TaskTracker/jsonParser/lexer"
	"github.com/kaushikmak/go-projects/TaskTracker/jsonParser/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	// Read two tokens to set both curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) Parse() interface{} {
	return p.parseValue()
}

func (p *Parser) parseValue() interface{} {
	switch p.curToken.Type {
	case token.LBRACE:
		return p.parseObject()
	case token.LBRACKET:
		return p.parseArray()
	case token.STRING:
		return p.curToken.Literal
	case token.NUMBER:
		val, err := strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			p.errors = append(p.errors, fmt.Sprintf("could not parse float: %s", p.curToken.Literal))
			return nil
		}
		return val
	case token.TRUE:
		return true
	case token.FALSE:
		return false
	case token.NULL:
		return nil
	default:
		return nil
	}
}

func (p *Parser) parseObject() map[string]interface{} {
	obj := make(map[string]interface{})

	// current is {, move to next
	p.nextToken()

	// Empty object check
	if p.curToken.Type == token.RBRACE {
		return obj
	}

	for {
		if p.curToken.Type != token.STRING {
			p.errors = append(p.errors, fmt.Sprintf("expected string key, got %s", p.curToken.Type))
			return nil
		}

		key := p.curToken.Literal
		p.nextToken()

		if p.curToken.Type != token.COLON {
			p.errors = append(p.errors, fmt.Sprintf("expected :, got %s", p.curToken.Type))
			return nil
		}
		p.nextToken()

		val := p.parseValue()
		obj[key] = val

		p.nextToken()
		if p.curToken.Type == token.RBRACE {
			break
		}

		if p.curToken.Type != token.COMMA {
			p.errors = append(p.errors, fmt.Sprintf("expected comma or }, got %s", p.curToken.Type))
			return nil
		}
		p.nextToken()
	}

	return obj
}

func (p *Parser) parseArray() []interface{} {
	arr := []interface{}{}

	// current is [, move to next
	p.nextToken()

	if p.curToken.Type == token.RBRACKET {
		return arr
	}

	for {
		val := p.parseValue()
		arr = append(arr, val)

		p.nextToken()
		if p.curToken.Type == token.RBRACKET {
			break
		}

		if p.curToken.Type != token.COMMA {
			p.errors = append(p.errors, fmt.Sprintf("expected comma or ], got %s", p.curToken.Type))
			return nil
		}
		p.nextToken()
	}
	return arr
}
