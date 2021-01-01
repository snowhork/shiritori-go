package shiritoriapi

import (
	"context"
	"shiritori/gen/shiritori"
	"shiritori/values"

	"github.com/pkg/errors"
)

func (s *shiritorisrvc) Words(ctx context.Context, p *shiritori.WordsPayload) (*shiritori.Wordresult, error) {
	wordBody, err := values.NewWordBody(p.Word)
	if err != nil {
		return nil, err
	}

	exists, err := s.wordChecker.Check(ctx, wordBody)

	if err != nil {
		return nil, errors.Wrap(err, "WordCheck Error")
	}

	hash := s.wordSigner.Sign(wordBody, exists)

	return &shiritori.Wordresult{
		Word:   string(wordBody),
		Exists: exists,
		Hash:   string(hash),
	}, nil
}
