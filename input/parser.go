package input

import "github.com/ranog/tdd-in-go/calculator"

type Parser struct {
	engine    *calculator.Engine
	validator *Validator
}

func (p *Parser) ProcessExpression(expr string) (*string, error) {
	// implementation code
}

// ... method declarations
