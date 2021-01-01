package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BattleState_Validate(t *testing.T) {
	cases := []struct {
		name     string
		number   ThemeNumber
		char     WordChar
		word     WordBody
		expected bool
	}{
		{"Valid ThemeNumber, Valid ThemeChar, pass", NewThemeNumber(5, false), WordChar("か"), WordBody("かたたたき"), true},
		{"Invalid ThemeNumber, Valid ThemeChar, fail", NewThemeNumber(5, false), WordChar("か"), WordBody("かどかわ"), false},
		{"Valid ThemeNumber, Invalid ThemeChar, fail", NewThemeNumber(5, false), WordChar("か"), WordBody("せたがやく"), false},
		{"Invalid ThemeNumber, Invalid ThemeChar, fail", NewThemeNumber(5, false), WordChar("か"), WordBody("せたがや"), false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := BattleState{ThemeNumber: c.number, ThemeChar: c.char}
			assert.Equal(t, c.expected, s.Validate(c.word))
		})
	}

}

func Test_ThemeNumber_Validate(t *testing.T) {
	cases := []struct {
		name     string
		number   ThemeNumber
		word     WordBody
		expected bool
	}{
		{"By 5, when length is under 5, fail", NewThemeNumber(5, false), WordBody("かどかわ"), false},
		{"By 5, when length is exactly 5, fail", NewThemeNumber(5, false), WordBody("かたたたき"), true},
		{"By 5, when length is over 5, fail", NewThemeNumber(5, false), WordBody("かたやきそば"), false},

		{"By 5+, when length is under 5, fail.", NewThemeNumber(5, true), WordBody("かどかわ"), false},
		{"By 5+, when length is exactly 5, pass.", NewThemeNumber(5, true), WordBody("かたたたき"), true},
		{"By 5+, when lenght is over 5, pass.", NewThemeNumber(5, true), WordBody("かたやきそば"), true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, c.number.Validate(c.word))
		})
	}

}
