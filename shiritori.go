package shiritoriapi

import (
	"context"
	"io"
	"log"
	"shiritori/gen/shiritori"
	"shiritori/pkg/wordchecker"
	"shiritori/pkg/wordsigner"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type shiritorisrvc struct {
	logger      *log.Logger
	wordChecker WordChecker
	wordSigner  WordSigner
}

type WordChecker interface {
	Check(ctx context.Context, word string) (bool, error)
}

type WordSigner interface {
	Sign(word string, exists bool) string
}

// NewShiritori returns the shiritori service implementation.
func NewShiritori(logger *log.Logger) shiritori.Service {
	return &shiritorisrvc{logger: logger, wordChecker: wordchecker.NewWordChecker(), wordSigner: wordsigner.NewWordSigner("123456789")}
}

// Add implements add.
func (s *shiritorisrvc) Add(ctx context.Context, p *shiritori.AddPayload) (res int, err error) {
	s.logger.Print("shiritori.add")
	return
}

// Battle implements battle.
func (s *shiritorisrvc) Battle(ctx context.Context, p *shiritori.BattlePayload, stream shiritori.BattleServerStream) error {
	if err := stream.Send(&shiritori.Battleevent{BattleID: &p.BattleID}); err != nil {
		return stream.Close()
	}

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		for {
			data, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					return errors.Wrap(err, "Message Receive Error")
				}
			}

			if data.Msg != nil && *data.Msg == "close" {
				return nil
			}

			if err := stream.Send(&shiritori.Battleevent{BattleID: data.Msg}); err != nil {
				return errors.Wrap(err, "Message Send Error")
			}
		}
	})

	if err := g.Wait(); err != nil {
		s.logger.Print(err)
	}

	return stream.Close()
}
