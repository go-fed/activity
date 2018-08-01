package pub

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	mediaMissingErr     = errors.New("media type missing /")
	mediaSubtypeMissing = errors.New("media subtype missing after /")
	missingEqualsError  = errors.New("media parameter missing equals sign for token")
	tooManyQuotesError  = errors.New("too many quotes for optional parameter")
)

var _ decoderState = &mediaRangeState{}
var _ decoderState = &mediaRangeOptionalParameterState{}

type mediaRangeState struct {
	isEndOfRange bool
	next         decoderState
}

func (s *mediaRangeState) Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	err = mediaMissingErr
	s.isEndOfRange = false
	s.next = s
	// Skip any leading optional whitespace
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	// Scan until a ';' or ','. Ensure we have seen a '/'.
	for runeWidth, width, i := 0, 0, start; i < len(data); i, width = i+runeWidth, width+runeWidth {
		var r rune
		r, runeWidth = utf8.DecodeRune(data[i:])
		if r == ';' {
			s.next = &mediaRangeOptionalParameterState{}
			return i, data[start:width], err
		} else if r == ',' {
			s.isEndOfRange = true
			return i, data[start:width], err
		} else if r == '/' {
			err = mediaSubtypeMissing
		} else if !unicode.IsSpace(r) && err == mediaSubtypeMissing {
			err = nil
		}
	}
	if atEOF && len(data) > start {
		s.isEndOfRange = true
		return len(data), data[start:], err
	}
	// Need more data
	return 0, nil, nil
}

func (s *mediaRangeState) Result(token string, r *mediaTypeHeaderResult) bool {
	r.mediaRange = token
	return s.isEndOfRange
}

func (s *mediaRangeState) Next() decoderState {
	return s.next
}

type mediaRangeOptionalParameterState struct {
	isEndOfRange bool
	isQuoted     bool
	next         decoderState
}

func (s *mediaRangeOptionalParameterState) Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	err = missingEqualsError
	s.isEndOfRange = false
	s.next = s
	// Skip any leading optional whitespace
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	s.isQuoted = false
	const (
		tokenState = iota
		tokenOrQuoteState
		quotedStringState
		endState
	)
	processingState := tokenState
	processingFirstToken := true
	// Scan until an unquoted ',' or ';'. Make sure we have seen '='
	for runeWidth, width, i := 0, 0, 0; i < len(data); i, width = i+runeWidth, width+runeWidth {
		var r rune
		r, runeWidth = utf8.DecodeRune(data[i:])
		switch processingState {
		case tokenState:
			if r == '=' {
				if processingFirstToken {
					err = nil
					processingFirstToken = false
					processingState = tokenOrQuoteState
				} else {
					return 0, nil, fmt.Errorf("illegal value for token: %s", r)
				}
			} else if r == ')' ||
				r == '(' ||
				r == '/' ||
				r == ':' ||
				r == '<' ||
				r == '>' ||
				r == '?' ||
				r == '@' ||
				r == '[' ||
				r == ']' ||
				r == '\\' ||
				r == '{' ||
				r == '}' ||
				r == '"' {
				return 0, nil, fmt.Errorf("illegal value for token: %s", r)
			} else if (r == ',' || r == ';') && processingFirstToken {
				return 0, nil, fmt.Errorf("illegal value for token: %s", r)
			} else if r == ',' {
				s.next = &mediaRangeState{}
				s.isEndOfRange = true
				return i, data[start:width], err
			} else if r == ';' {
				s.next = &mediaRangeOptionalParameterState{}
				return i, data[start:width], err
			}
		case quotedStringState:
			if r == '"' {
				processingState = endState
			}
		case tokenOrQuoteState:
			if r == '"' {
				processingState = quotedStringState
				s.isQuoted = true
			} else {
				processingState = tokenState
			}
		case endState:
			if r == ',' {
				s.next = &mediaRangeState{}
				s.isEndOfRange = true
				return i, data[start:width], err
			} else if r == ';' {
				s.next = &mediaRangeOptionalParameterState{}
				return i, data[start:width], err
			} else {
				return 0, nil, fmt.Errorf("illegal value after quoted string: %s", r)
			}
		}
	}
	if atEOF && len(data) > start {
		s.next = &mediaRangeState{}
		s.isEndOfRange = true
		return len(data), data[start:], err
	}
	// Need more data
	return 0, nil, nil
}

func (s *mediaRangeOptionalParameterState) Result(token string, r *mediaTypeHeaderResult) bool {
	// Token is already well-formed, so split on first equals and strip any literal quotes
	pair := strings.SplitAfterN(token, "=", 1)
	tok := strings.TrimSpace(pair[0])
	value := strings.TrimSpace(pair[1])
	if s.isQuoted {
		value = strings.Trim(value, "\"")
	}
	r.mediaRangeParameters = append(r.mediaRangeParameters, parameterResult{
		token: tok,
		value: value,
	})
	return s.isEndOfRange
}

func (s *mediaRangeOptionalParameterState) Next() decoderState {
	return s.next
}

type parameterResult struct {
	token string
	value string
}

type mediaTypeHeaderResult struct {
	mediaRange           string
	mediaRangeParameters []parameterResult
}

type decoderState interface {
	// Split splits the data into a next token.
	Split(data []byte, atEOF bool) (advance int, token []byte, err error)
	// Result applies the token returned by this state to the given result.
	// Returns true if the result is completed and a brand new one should
	// begin.
	Result(token string, r *mediaTypeHeaderResult) bool
	// Next is always called after Split, to obtain the next decoder state.
	Next() decoderState
}

type mediaHeaderDecoder struct {
	state decoderState
}

func (m *mediaHeaderDecoder) splitAcceptBoundaries(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	advance, token, err = m.state.Split(data, atEOF)
	return
}

func (m *mediaHeaderDecoder) Parse(s string) []*mediaTypeHeaderResult {
	m.state = &mediaRangeState{}
	scanner := bufio.NewScanner(bytes.NewBufferString(s))
	scanner.Split(m.splitAcceptBoundaries)
	var results []*mediaTypeHeaderResult
	current := &mediaTypeHeaderResult{}
	for scanner.Scan() {
		token := scanner.Text()
		if m.state.Result(token, current) {
			results = append(results, current)
			current = &mediaTypeHeaderResult{}
		}
		m.state = m.state.Next()
	}
	return results
}
