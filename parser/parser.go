package parser

import (
	"github.com/alanfoster/assembler/ast"
	"github.com/alanfoster/assembler/lexer"
	"github.com/alanfoster/assembler/token"
	"bytes"
)

type Parser struct {
	lexer   *lexer.Lexer
	current token.Token
	peek    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
	}
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (p *Parser) HasMoreInstructions() bool {
	return !p.isCurrent(token.EOF)
}

func (p *Parser) Advance() ast.Node {
	switch p.current.Type {
	case token.AT:
		return p.parseAInstruction()
	default:
		return p.parseCInstruction()
	}
}

// @value
//
// Where value is a symbol or number
func (p *Parser) parseAInstruction() *ast.AInstruction {
	p.advance(token.AT)

	value := p.current.Lexeme
	p.advance(token.VALUE)

	return &ast.AInstruction{
		Value: value,
	}
}

// C ->
// Dest = Comp; Jump
// | Dest = Comp
// | Comp; Jump
// | Comp
func (p *Parser) parseCInstruction() *ast.CInstruction {
	instr := &ast.CInstruction{}

	if p.isPeek(token.EQUALS) {
		instr.Destination = p.parseDest()
		p.advance(token.EQUALS)
	}

	instr.Command = p.parseCommand()

	if p.isCurrent(token.SEMICOLON) {
		p.advance(token.SEMICOLON)
		instr.Jump = p.parseJump()
	}

	return instr
}

// Parses the destination.
// This could ensure a valid recipient, but it does not.
func (p *Parser) parseDest() *ast.Value {
	current := p.current
	p.advance(token.VALUE)
	return &ast.Value{Value: current.Lexeme}
}

// Parses the jump location.
// This could ensure a valid Jump location, but it does not.
func (p *Parser) parseJump() *ast.Value {
	current := p.current
	p.advance(token.VALUE)
	return &ast.Value{Value: current.Lexeme}
}

// Representation is somewhat cheaty, to make lookup easier.
//
// Command ->
// 	operator Value
// 	| Value operator Value
// 	| Value
func (p *Parser) parseCommand() ast.Command {
	var out bytes.Buffer

	// prefix operator value
	if p.isCurrent(token.OPERATOR) {
		out.WriteString(p.current.Lexeme)
		p.advance(token.OPERATOR)
		out.WriteString(p.current.Lexeme)
		p.advance(token.VALUE)
		return ast.Command{Value: out.String()}
	}

	// Reading the value
	out.WriteString(p.current.Lexeme)
	p.advance(token.VALUE)

	// Handle Infix
	if (p.isCurrent(token.OPERATOR)) {
		out.WriteString(p.current.Lexeme)
		p.advance(token.OPERATOR)
		out.WriteString(p.current.Lexeme)
		p.advance(token.VALUE)
	}

	return ast.Command{Value: out.String()}
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.lexer.Advance()
}

func (p *Parser) isCurrent(tokenType token.Type) bool {
	return p.current.Type == tokenType
}

func (p *Parser) isPeek(tokenType token.Type) bool {
	return p.peek.Type == tokenType
}

func (p *Parser) advance(tokenType token.Type) {
	if !p.isCurrent(tokenType) {
		panic("Unexpected token")
	}

	p.nextToken()
}
