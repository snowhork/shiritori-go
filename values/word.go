package values

import (
	"errors"
	"unicode"
)

type WordChar string
type WordBody string
type WordBodyHash string

var InvalidWordBody = errors.New("invalid word body")

func NewWordBody(word string) (WordBody, error) {
	for _, r := range word {
		if !(unicode.In(r, unicode.Hiragana) || r == '\u30FC') { // ー（長音）
			return WordBody(""), InvalidWordBody
		}
	}
	return WordBody(word), nil
}

func (w WordBody) TopChar() WordChar {
	r := []rune(w)
	return WordChar(r[:1])
}

func (w WordBody) LastChar() WordChar {
	r := []rune(w)
	return WordChar(r[len(r)-1:])
}

func (w WordBody) Length() int {
	return len([]rune(w))
}
