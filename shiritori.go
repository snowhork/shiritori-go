package shiritoriapi

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"log"
	shiritori "shiritori/gen/shiritori"
	"golang.org/x/sync/errgroup"
)

// shiritori service example implementation.
// The example methods log the requests and return zero values.
type shiritorisrvc struct {
	logger *log.Logger
}

// NewShiritori returns the shiritori service implementation.
func NewShiritori(logger *log.Logger) shiritori.Service {
	return &shiritorisrvc{logger}
}

// Add implements add.
func (s *shiritorisrvc) Add(ctx context.Context, p *shiritori.AddPayload) (res int, err error) {
	s.logger.Print("shiritori.add")
	return
}

// Battle implements battle.
func (s *shiritorisrvc) Battle(ctx context.Context, p *shiritori.BattlePayload, stream shiritori.BattleServerStream) (error) {
	if err := stream.Send(&shiritori.Battleevent{BattleID: &p.BattleID}); err != nil {
		return stream.Close()
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		for {
			data, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					return errors.Wrap(err, "Message Receive Error")
				}
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
