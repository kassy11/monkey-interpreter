package ast

import "github.com/kassy11/monkey-interpreter/token"

// ASTの全てのノードはNodeを実装する
type Node interface {
	TokenLiteral() string // ノードが関連付けられているトークンのリテラル値を返す
}

type Statement interface {
	Node
	statementNode() // ダミーメソッド（Goコンパイラに情報を与えるため）
}

type Expression interface {
	Node
	expressionNode() // ダミーメソッド（Goコンパイラに情報を与えるため）
}

// ASTのルートノード
type Program struct {
	Statements []Statement // 文の集まりはここに格納される
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral()  { return ls.Token.Literal }

// 識別子を保持する
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
