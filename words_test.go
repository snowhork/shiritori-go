package shiritoriapi

import (
	"context"
	"fmt"
	"shiritori/gen/shiritori"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Words(t *testing.T) {
	cases := []struct {
		name     string
		payload  *shiritori.WordsPayload
		expected *shiritori.Wordresult
	}{
		{"When the word exists", &shiritori.WordsPayload{Word: "hoge"}, &shiritori.Wordresult{Word: "hoge", Exists: true, Hash: "hoge-true-key"}},
		{"When the word doesn't exist", &shiritori.WordsPayload{Word: "piyo"}, &shiritori.Wordresult{Word: "piyo", Exists: false, Hash: "piyo-false-key"}},
	}

	s := &shiritorisrvc{
		wordChecker: &mockWordChecker{},
		wordSigner:  &mockWordSigner{},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, _ := s.Words(context.TODO(), c.payload)

			assert.Equal(t, c.expected, actual)
		})
	}
}

type mockWordChecker struct{}

func (*mockWordChecker) Check(ctx context.Context, word string) (bool, error) {
	if word == "hoge" {
		return true, nil
	}
	return false, nil
}

type mockWordSigner struct{}

func (*mockWordSigner) Sign(word string, exists bool) string {
	return fmt.Sprintf("%s-%v-%s", word, exists, "key")
}
