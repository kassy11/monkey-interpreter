package parser

import (
	"fmt"
	"github.com/kassy11/monkey-interpreter/chap2/ast"
	"github.com/kassy11/monkey-interpreter/chap2/lexer"
	"github.com/kassy11/monkey-interpreter/chap2/token"
)

type Parser struct {
	l *lexer.Lexer // 字句解析インスタンス
	// lexerでのposition, readPositionのトークンver
	curToken  token.Token // 現在のトークン
	peekToken token.Token // 次のトークン
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// 2つのトークンを読み込んで、curTokenとpeekTokenがセットされる
	p.nextToken()
	p.nextToken()

	return p
}

// トークンを進める
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	// ASTのルートノードの作成
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// EOFになるまで、トークンを読み進める
	// !p.curTokenIs(token.EOF) でも良い
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetSatetment()
	case token.RETURN:
		return p.parseReturSatetment()
	default:
		return nil
	}
}

func (p *Parser) parseLetSatetment() *ast.LetStatement {
	// let x = 5; の形式になるように
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturSatetment() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON){
		p.nextToken()
	}
	return stmt
}


// 次のトークンが予想通りか
// 次のトークンが正しい場合に限って、トークンを進める
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// 現在のトークンが予想通りか
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}