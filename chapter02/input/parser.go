package input

import (
	"github.com/PacktPublishing/Test-Driven-Development-in-Go/chapter02/calculator"
)

type Parser struct {
	engine    *calculator.Engine
	validator *Validator
}

func (p *Parser) ProcessExpression(expr string) (*string, error) {
	// implementation code
}

// ... method declarations
