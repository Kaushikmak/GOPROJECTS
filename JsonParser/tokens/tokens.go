package tokens

type TokenType int

// list of constants
const (
	Illegal        TokenType = iota //"ILLEGAL"
	BeginArray                      //"["
	BeginObject                     // "{"
	EndArray                        // "]"
	EndObject                       // "}"
	NameSeparator                   // ":"
	ValueSeparator                  // ","
	EndOfFile                       // "EOF"
	String                          // "STRING"
	Number                          // "NUMBER"
	True                            // "TRUE"
	False                           // "FALSE"
	Null                            // "NULL"
)

// struct that hold token type and literal string
type Token struct {
	Type    TokenType
	Literal string
}

// constructor function for token
func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// function that print token to string
func (t TokenType) String() string {
	switch t {
	case Illegal:
		return "ILLEGAL"
	case BeginArray:
		return "["
	case BeginObject:
		return "{"
	case EndArray:
		return "]"
	case EndObject:
		return "}"
	case NameSeparator:
		return ":"
	case ValueSeparator:
		return ","
	case EndOfFile:
		return "EOF"
	case String:
		return "STRING"
	case Number:
		return "NUMBER"
	case True:
		return "TRUE"
	case False:
		return "FALSE"
	case Null:
		return "NULL"
	default:
		return "UNKNOWN"
	}
}
