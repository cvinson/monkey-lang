package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  // Specials
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifiers and literals
  IDENT = "IDENT"
  INT = "INT"
  STR = "STR"

  // Operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  SLASH = "/"
  ASTERISK = "*"

  // Comparisons
  LT = "<"
  GT = ">"
  EQ = "=="
  NOTEQ = "!="

  // Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"
)

var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }

  return IDENT
}
