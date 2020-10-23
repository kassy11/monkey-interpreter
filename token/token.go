// トークンを定義

package token

type TokenType string

type Token struct {
	Type    TokenType // 識別子
	Literal string    // トークンのリテラル値
}

// トークンの種類を定義
const (
	ILLEGAL = "ILLEGAL" // 未知なトークン
	EOF     = "EOF"     // ファイル終端

	// 識別子＋リテラル
	IDENT = "IDENT"
	INT   = "INT"

	// 演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// keywordsテーブルを確認して、渡された識別子がキーワードに当たるかを返す
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
