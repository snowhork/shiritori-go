package shiritoriapi

import (
	"context"
	"shiritori/gen/shiritori"

	"github.com/pkg/errors"
)

func (s *shiritorisrvc) Words(ctx context.Context, p *shiritori.WordsPayload) (*shiritori.Wordresult, error) {
	exists, err := s.wordChecker.Check(ctx, p.Word)

	if err != nil {
		return nil, errors.Wrap(err, "WordCheck Error")
	}

	hash := s.wordSigner.Sign(p.Word, exists)

	return &shiritori.Wordresult{
		Word:   p.Word,
		Exists: exists,
		Hash:   hash,
	}, nil
}
