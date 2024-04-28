// This will only support ASCII and not full Unicode.
// If we wanted full Unicode then we need to change
// the ch byte to a rune(int32 to represent Unicode)
// and adjust how we read the next characters since
// they could be multiple bytes long.

package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	// l := &Lexer{input: input}
	// return l
	return &Lexer{input: input}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII for NUL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
