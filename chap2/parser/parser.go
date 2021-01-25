package parser

import (
	"github.com/kassy11/monkey-interpreter/chap2/ast"
	"github.com/kassy11/monkey-interpreter/chap2/token"
	"github.com/kassy11/monkey-interpreter/chap2/lexer"
)

type Parser struct {
	l *lexer.Lexer // 字句解析インスタンス
	// lexerでのposition, readPositionのトークンver
	curToken token.Token // 現在のトークン
	peekToken token.Token // 次のトークン
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// 2つのトークンを読み込んで、curTokenとpeekTokenがセットされる
	p.nextToken()
	p.nextToken()

	return p
}

// トークンを進める
func (p *Parser) nextToken(){
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
	// ASTのルートノードの作成
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// EOFになるまで、トークンを読み進める
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil{
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement{
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetSatetment()
	default:
		return nil
	}
}

func ( p *Parser ) parseLetSatetment() *ast.LetStatement{
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT){
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN){
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON){
		p.nextToken()
	}

	return stmt
}


func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIS(t){
		p.nextToken()
		return true
	}else{
		return false
	}
}

func (p *Parser) peekTokenIS(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}