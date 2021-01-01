package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewWordBody(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected error
	}{
		{"ひらがなの場合", "りんご", nil},
		{"カタカナの場合", "リンゴ", InvalidWordBody},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := NewWordBody(c.input)
			assert.Equal(t, c.expected, err)
		})
	}
}

func Test_WordBody_LastChar(t *testing.T) {
	cases := []struct {
		name     string
		wordBody WordBody
		expected WordChar
	}{{"ひらがなの場合", WordBody("りんご"), WordChar("ご")}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, c.wordBody.LastChar())
		})
	}
}

func Test_WordBody_Length(t *testing.T) {
	cases := []struct {
		name     string
		wordBody WordBody
		expected int
	}{{"ひらがなの場合", WordBody("りんご"), 3}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, c.wordBody.Length())
		})
	}
}
