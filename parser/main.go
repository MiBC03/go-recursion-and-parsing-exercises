package main

import "fmt"

// Token representa um par (tipo, valor)
type Token struct {
	Type  string
	Value string
}

// ===== AST =====

type ASTNode interface{}

type Program struct {
	Name string
	Body ASTNode
}

type BeginEnd struct {
	Statements []ASTNode
}

type Assign struct {
	Name string
	Expr ASTNode
}

type If struct {
	Cond ASTNode
	Then ASTNode
	Else ASTNode
}

type While struct {
	Cond ASTNode
	Body ASTNode
}

type Read struct {
	Name string
}

type Write struct {
	Expr ASTNode
}

type BinaryExpr struct {
	Op    string
	Left  ASTNode
	Right ASTNode
}

type IntLiteral struct {
	Value int
}

type Identifier struct {
	Name string
}

// ===== Parser =====

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) current() Token {
	if p.pos >= len(p.tokens) {
		return Token{}
	}
	return p.tokens[p.pos]
}

func (p *Parser) consume() Token {
	t := p.current()
	p.pos++
	return t
}

func (p *Parser) parseProgram() ASTNode {
	p.consume() // "programa"
	name := p.consume().Value
	p.consume() // ";"

	body := p.parseStatement()
	p.consume() // "fim"

	return Program{Name: name, Body: body}
}

func (p *Parser) parseStatement() ASTNode {
	t := p.current()

	switch t.Value {

	case "comeco":
		p.consume()
		var stats []ASTNode
		for p.current().Value != "fim" {
			stats = append(stats, p.parseStatement())
			if p.current().Value == ";" {
				p.consume()
			}
		}
		p.consume()
		return BeginEnd{Statements: stats}

	case "leitura":
		p.consume()
		id := p.consume().Value
		return Read{Name: id}

	case "escrita":
		p.consume()
		return Write{Expr: p.parseExpression()}

	default:
		id := p.consume().Value
		p.consume() // :=
		return Assign{Name: id, Expr: p.parseExpression()}
	}
}

func (p *Parser) parseExpression() ASTNode {
	left := p.parseTerm()

	for p.current().Value == "+" || p.current().Value == "-" {
		op := p.consume().Value
		right := p.parseTerm()
		left = BinaryExpr{Op: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseTerm() ASTNode {
	left := p.parseFactor()

	for p.current().Value == "*" || p.current().Value == "/" {
		op := p.consume().Value
		right := p.parseFactor()
		left = BinaryExpr{Op: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseFactor() ASTNode {
	t := p.consume()

	if t.Type == "inteiro" {
		return IntLiteral{Value: atoi(t.Value)}
	}

	return Identifier{Name: t.Value}
}

func atoi(s string) int {
	var n int
	fmt.Sscan(s, &n)
	return n
}

func main() {
	tokens := []Token{
		{"palavra_chave", "programa"},
		{"id", "teste"},
		{";", ";"},
		{"palavra_chave", "comeco"},
		{"palavra_chave", "leitura"},
		{"id", "x"},
		{";",";"},
		{"palavra_chave", "escrita"},
		{"id", "x"},
		{"palavra_chave", "fim"},
		{"palavra_chave", "fim"},
	}

	parser := NewParser(tokens)
	ast := parser.parseProgram()
	fmt.Printf("%#v\n", ast)
}
