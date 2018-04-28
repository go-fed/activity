package pub

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

type scanState int

const (
	noScanState scanState = iota
	keyScanState
	valueScanState
	colonScanState
	commaOrEndScanState
	quotedLiteralScanState
	trueLiteralScanState
	falseLiteralScanState
	nullLiteralScanState
	numberLiteralScanState
	doneScanState
)

type scanContext int

const (
	noScanContext scanContext = iota
	keyScanContext
	valueScanContext
	objectScanContext
	arrayScanContext
)

type scanner struct {
	b                     *bytes.Buffer
	state                 scanState
	context               []scanContext
	buf                   []byte
	literal               []rune
	nextRuneEscaped       bool
	seenOneDigit          bool
	seenDecimalPoint      bool
	seenDigitAfterDecimal bool
	seenExp               bool
	seenCharAfterExp      bool
	hasComma              bool
}

// Private implementation

func newScanner(b []byte) *scanner {
	return &scanner{
		b: bytes.NewBuffer(b),
	}
}

func (s *scanner) push(c scanContext) {
	s.context = append(s.context, c)
}

func (s *scanner) pop() {
	s.context = s.context[:len(s.context)-1]
}

func (s *scanner) peek() scanContext {
	return s.context[len(s.context)-1]
}

func (s *scanner) scan() error {
	buf := make([]byte, 0, 1)
	for !utf8.FullRune(buf) {
		buf = append(buf, s.b.Next(1)[0])
	}
	r, _ := utf8.DecodeRune(buf)
	if r == utf8.RuneError {
		return fmt.Errorf("Unrecognized utf8 bytes: %v", buf)
	}
	var err error
	switch s.state {
	case noScanState:
		err = s.noScanState(r)
	case keyScanState:
		err = s.keyScanState(r)
	case valueScanState:
		err = s.valueScanState(r)
	case colonScanState:
		err = s.colonScanState(r)
	case commaOrEndScanState:
		err = s.commaOrEndScanState(r)
	case quotedLiteralScanState:
		err = s.quotedLiteralScanState(r)
	case trueLiteralScanState:
		err = s.trueLiteralScanState(r)
	case falseLiteralScanState:
		err = s.falseLiteralScanState(r)
	case nullLiteralScanState:
		err = s.nullLiteralScanState(r)
	case numberLiteralScanState:
		err = s.numberLiteralScanState(r)
	case doneScanState:
		err = s.doneScanState(r)
	}
	return err
}

func (s *scanner) noScanState(r rune) error {
	if r == '{' {
		s.state = keyScanState
		s.push(objectScanContext)
		s.hasComma = false
	} else if r == '[' {
		s.state = valueScanState
		s.push(arrayScanContext)
		s.hasComma = false
	} else if !unicode.IsSpace(r) {
		return fmt.Errorf("Disallowed JSON beginning rune: %s", r)
	}
	return nil
}

func (s *scanner) keyScanState(r rune) error {
	if r == '"' {
		s.state = quotedLiteralScanState
		s.push(keyScanContext)
		s.literal = make([]rune, 0, 1)
	} else if r == '}' {
		if s.hasComma {
			return fmt.Errorf("Disallowed } after comma")
		}
		s.pop()
		if len(s.context) == 0 {
			s.state = doneScanState
		} else {
			s.state = commaOrEndScanState
		}
	} else if !unicode.IsSpace(r) {
		return fmt.Errorf("Disallowed JSON looking for key rune: %s", r)
	}
	return nil
}

func (s *scanner) valueScanState(r rune) error {
	if r == '"' {
		s.state = quotedLiteralScanState
		s.push(valueScanContext)
		s.literal = make([]rune, 0, 1)
	} else if r == 't' { // true
		s.state = trueLiteralScanState
		s.literal = []rune{r}
	} else if r == 'f' { // false
		s.state = falseLiteralScanState
		s.literal = []rune{r}
	} else if r == 'n' { // null
		s.state = nullLiteralScanState
		s.literal = []rune{r}
	} else if r == '{' {
		s.state = keyScanState
		s.push(objectScanContext)
		s.hasComma = false
	} else if r == '[' {
		s.state = valueScanState
		s.push(arrayScanContext)
		s.hasComma = false
	} else if r == ']' {
		if s.peek() != arrayScanContext {
			return fmt.Errorf("Unexpected ] in wrong context for JSON string")
		} else if !s.hasComma {
			return fmt.Errorf("Unexpected ] after comma in JSON string")
		}
		s.pop()
		if len(s.context) == 0 {
			s.state = doneScanState
		} else {
			s.state = commaOrEndScanState
		}
	} else if r == '-' || unicode.IsDigit(r) {
		s.state = numberLiteralScanState
		s.literal = []rune{r}
		s.seenOneDigit = unicode.IsDigit(r)
		s.seenDecimalPoint = false
		s.seenDigitAfterDecimal = false
		s.seenExp = false
		s.seenCharAfterExp = false
	} else if !unicode.IsSpace(r) {
		return fmt.Errorf("Disallowed JSON looking for key rune: %s", r)
	}
	return nil
}

func (s *scanner) colonScanState(r rune) error {
	if r == ':' {
		s.state = valueScanState
	} else if !unicode.IsSpace(r) {
		return fmt.Errorf("Disallowed JSON looking for key rune: %s", r)
	}
	return nil
}

func (s *scanner) commaOrEndScanState(r rune) error {
	if r == ',' {
		s.state = keyScanState
		s.hasComma = true
		return nil
	}
	switch s.peek() {
	case objectScanContext:
		if r == '}' {
			s.pop()
			if len(s.context) == 0 {
				s.state = doneScanState
			}
			return nil
		}
	case arrayScanContext:
		if r == ']' {
			s.pop()
			if len(s.context) == 0 {
				s.state = doneScanState
			}
			return nil
		}
	default:
		return fmt.Errorf("Disallowed internal context when in commaOrEndScanState: %v", s.peek())
	}
	if !unicode.IsSpace(r) {
		return fmt.Errorf("Disallowed JSON looking for comma or end scan state: %s", r)
	}
	return nil
}

func (s *scanner) quotedLiteralScanState(r rune) error {
	done := false
	if r == '"' && !s.nextRuneEscaped {
		done = true
		s.nextRuneEscaped = false
	} else if r == '\\' {
		s.nextRuneEscaped = !s.nextRuneEscaped
	} else {
		s.nextRuneEscaped = false
	}
	if done {
		switch s.peek() {
		case keyScanContext:
			s.state = colonScanState
			s.pop()
			if len(s.context) == 0 {
				return fmt.Errorf("Quoted literal key has no parent context")
			}
		case valueScanContext:
			s.state = commaOrEndScanState
			s.pop()
			if len(s.context) == 0 {
				return fmt.Errorf("Quoted literal value has no parent context")
			}
		default:
			return fmt.Errorf("Disallowed context when parsing quoted literal: %v", s.peek())
		}
	} else {
		s.literal = append(s.literal, r)
	}
	return nil
}

func (s *scanner) trueLiteralScanState(r rune) error {
	switch len(s.literal) {
	case 1:
		if r != 'r' {
			return fmt.Errorf("Unexpected rune for 'true' literal: %v", r)
		}
	case 2:
		if r != 'u' {
			return fmt.Errorf("Unexpected rune for 'true' literal: %v", r)
		}
	case 3:
		if r != 'e' {
			return fmt.Errorf("Unexpected rune for 'true' literal: %v", r)
		}
		s.state = commaOrEndScanState
	default:
		return fmt.Errorf("Unexpected internal literal length for 'true' literal: %v", s.literal)
	}
	s.literal = append(s.literal, r)
	return nil
}

func (s *scanner) falseLiteralScanState(r rune) error {
	switch len(s.literal) {
	case 1:
		if r != 'a' {
			return fmt.Errorf("Unexpected rune for 'false' literal: %v", r)
		}
	case 2:
		if r != 'l' {
			return fmt.Errorf("Unexpected rune for 'false' literal: %v", r)
		}
	case 3:
		if r != 's' {
			return fmt.Errorf("Unexpected rune for 'false' literal: %v", r)
		}
	case 4:
		if r != 'e' {
			return fmt.Errorf("Unexpected rune for 'false' literal: %v", r)
		}
		s.state = commaOrEndScanState
	default:
		return fmt.Errorf("Unexpected internal literal length for 'false' literal: %v", s.literal)
	}
	s.literal = append(s.literal, r)
	return nil
}

func (s *scanner) nullLiteralScanState(r rune) error {
	switch len(s.literal) {
	case 1:
		if r != 'u' {
			return fmt.Errorf("Unexpected rune for 'null' literal: %v", r)
		}
	case 2:
		if r != 'l' {
			return fmt.Errorf("Unexpected rune for 'null' literal: %v", r)
		}
	case 3:
		if r != 'l' {
			return fmt.Errorf("Unexpected rune for 'null' literal: %v", r)
		}
		s.state = commaOrEndScanState
	default:
		return fmt.Errorf("Unexpected internal literal length for 'null' literal: %v", s.literal)
	}
	s.literal = append(s.literal, r)
	return nil
}

func (s *scanner) numberLiteralScanState(r rune) error {
	if unicode.IsDigit(r) {
		s.seenOneDigit = true
		if s.seenDecimalPoint {
			s.seenDigitAfterDecimal = true
		}
		if s.seenExp {
			s.seenCharAfterExp = true
		}
	} else if r == '.' {
		if !s.seenOneDigit {
			return fmt.Errorf("invalid number: decimal point before any digits")
		} else if s.seenDecimalPoint {
			return fmt.Errorf("invalid number: multiple decimal points")
		} else if s.seenExp {
			return fmt.Errorf("invalid number: cannot have decimal in exponent")
		}
		s.seenDecimalPoint = true
	} else if r == 'e' || r == 'E' {
		if !s.seenOneDigit {
			return fmt.Errorf("invalid number: e notation before digit")
		} else if s.seenDecimalPoint && !s.seenDigitAfterDecimal {
			return fmt.Errorf("invalid number: e notation before digit after a decimal point")
		} else if s.seenExp {
			return fmt.Errorf("invalid number: multiple e notations")
		}
		s.seenExp = true
	} else if r == '+' || r == '-' {
		if !s.seenExp {
			return fmt.Errorf("invalid number: cannot have +/- outside of e notation")
		} else if s.seenCharAfterExp {
			return fmt.Errorf("invalid number: +/- must follow e notation")
		}
		s.seenCharAfterExp = true
	} else {
		// Attempt to handle it as if already in next state.
		s.state = commaOrEndScanState
		return s.commaOrEndScanState(r)
	}
	s.literal = append(s.literal, r)
	return nil
}

func (s *scanner) doneScanState(r rune) error {
	if !unicode.IsSpace(r) {
		return fmt.Errorf("Disallowed character after reading JSON: %s", r)
	}
	return nil
}
