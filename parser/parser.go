package parser

import (
	"github.com/alanfoster/assembler/ast"
	"github.com/alanfoster/assembler/lexer"
	"github.com/alanfoster/assembler/token"
	"bytes"
	"fmt"
	"strconv"
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

func (p *Parser) ParseProgram() ast.Program {
	program := ast.Program{
		Instructions: []ast.Instruction{},
	}

	for p.HasMoreInstructions() {
		switch p.current.Type {
		case token.AT:
			instr := p.parseAInstruction()
			program.Instructions = append(program.Instructions, instr)
		case token.LEFT_BRACKET:
			instr := p.parseLInstruction()
			program.Instructions = append(program.Instructions, instr)
		default:
			instr := p.parseCInstruction()
			program.Instructions = append(program.Instructions, instr)
		}
	}

	return program
}

// @value
//
// Where value is a symbol or number
func (p *Parser) parseAInstruction() ast.Instruction {
	p.advance(token.AT)
	var value ast.AInstructionValue

	if p.isCurrent(token.NUMBER) {
		number, err := strconv.ParseInt(p.current.Lexeme, 10, 16)
		if err != nil {
			panic(err)
		}
		p.advance(token.NUMBER)
		value = &ast.Number{Value: int(number)}
	} else if p.isCurrent(token.VALUE) {
		value = &ast.Variable{Name: p.current.Lexeme}
		p.advance(token.VALUE)
	}

	return &ast.AInstruction{
		Value: value,
	}
}

// LInstruction -> LeftBrace Value RightBrace
func (p *Parser) parseLInstruction() ast.Instruction {
	p.advance(token.LEFT_BRACKET)
	value := p.current
	p.advance(token.VALUE)
	p.advance(token.RIGHT_BRACKET)

	return &ast.LInstruction{
		Value: value.Lexeme,
	}
}

// CInstruction ->
// Dest = Comp; Jump
// | Dest = Comp
// | Comp; Jump
// | Comp
func (p *Parser) parseCInstruction() ast.Instruction {
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
	p.advance(token.JUMP)
	return &ast.Value{Value: current.Lexeme}
}

// Representation is somewhat cheaty, to make lookup easier later.
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
		out.WriteString(p.parseNumberOrValue().Lexeme)
		return ast.Command{Value: out.String()}
	}

	// Reading the value
	out.WriteString(p.parseNumberOrValue().Lexeme)

	// Handle Infix
	if p.isCurrent(token.OPERATOR) {
		out.WriteString(p.current.Lexeme)
		p.advance(token.OPERATOR)
		out.WriteString(p.parseNumberOrValue().Lexeme)
	}

	return ast.Command{Value: out.String()}
}

func (p *Parser) parseNumberOrValue() token.Token {
	s := p.current
	if p.isCurrent(token.NUMBER) {
		p.advance(token.NUMBER)
	} else {
		p.advance(token.VALUE)
	}

	return s
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
		panic(fmt.Errorf("expected token type %s, instead got: %+v", tokenType, p.current))
	}

	p.nextToken()
}
