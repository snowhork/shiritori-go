package shiritoriapi

import (
	"context"
	"fmt"
	"shiritori/gen/shiritori"
	"shiritori/values"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Words(t *testing.T) {
	cases := []struct {
		name     string
		payload  *shiritori.WordsPayload
		expected *shiritori.Wordresult
	}{
		{"When the word exists", &shiritori.WordsPayload{Word: "りんご"}, &shiritori.Wordresult{Word: "りんご", Exists: true, Hash: "りんご-true-key"}},
		{"When the word doesn't exist", &shiritori.WordsPayload{Word: "ばなな"}, &shiritori.Wordresult{Word: "ばなな", Exists: false, Hash: "ばなな-false-key"}},
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

func (*mockWordChecker) Check(ctx context.Context, word values.WordBody) (bool, error) {
	if word == "りんご" {
		return true, nil
	}
	return false, nil
}

type mockWordSigner struct{}

func (*mockWordSigner) Sign(word values.WordBody, exists bool) values.WordBodyHash {
	return values.WordBodyHash(fmt.Sprintf("%s-%v-%s", word, exists, "key"))
}
