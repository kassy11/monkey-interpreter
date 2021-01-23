// 字句解析器
// ソースコードを受け取り、そのソースコードのトークン列を受け取る

package lexer

import "github.com/kassy11/monkey-interpreter/chap1/token"

// 字句解析対象の文字列と
type Lexer struct {
	input string // // 字句解析対象の文字列
	// 以下2つの位置はいずれも、入力中の文字にアクセスするためのインデックスとして用いる
	position     int  // 現在検査中の文字chの位置を指す
	readPosition int  // 現在の文字位置の次（これから読み込む位置）
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // 構造体のinputに引数で受け取る文字列を格納する
	l.readChar()
	return l
}

// 一文字（バイト）単位で読み進める
// 次の１文字を読み込んで現在のinput文字位置を一つ進める
func (l *Lexer) readChar() {
	// 最後まで読み込んだとき
	if l.readPosition >= len(l.input) {
		l.ch = 0 // asciiコードの『ファイル終端』を表す0にしておく
	} else {
		// 次の文字を読み出す
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// トークン単位で読み進める
// inputから現在検査中の文字のトークン構造体を一つ生成して返す
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' { // =の次が=のとき(==のとき)
			ch := l.ch
			// 一文字読み進める
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' { // 次が=のとき(!=のとき)
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // 次にNewToken()を読んだときに、l.chが更新されているようにする
	return tok
}

// 新しいトークン構造体を生成して返す
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 識別子を読んで、非英字に到達するまで字句解析の位置を進める
func (l *Lexer) readIdentifier() string {
	position := l.position
	// 文字でいる間ずっと読み込む
	for isLetter(l.ch) {
		l.readChar()
	}
	// 文字列を返す
	return l.input[position:l.position]
}

// 現在読んでいる文字が英字かどうか
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// 次の入力を覗き見する
// readChar()とほぼ同じで、インクリメントしないところは異なる→覗き見
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
