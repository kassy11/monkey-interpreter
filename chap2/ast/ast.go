package ast

import (
	"github.com/kassy11/monkey-interpreter/chap2/token"
)

type Node interface{
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// ルートノード
type Program struct {
	Statements []Statement // プログラムはここに格納される
}

// ノードが関連付けられているトークン値を返す
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LETトークン
	Name *Identifier // 変数の名前
	Value Expression
}

func (ls *LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal}

type Identifier struct {
	Token token.Token // token.IDENTトークン
	Value string
}

func (i *Identifier) expressionNode(){}
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}

