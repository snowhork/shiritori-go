package shiritoriapi

import (
	"context"
	"shiritori/gen/shiritori"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func (s *shiritorisrvc) Battle(ctx context.Context, p *shiritori.BattlePayload, stream shiritori.BattleServerStream) error {
	battleId := p.BattleID

	if err := stream.Send(&shiritori.Battlestreamingresult{Type: "init"}); err != nil {
		return stream.Close()
	}

	ticker := time.NewTicker(50 * time.Millisecond)
	stop := make(chan struct{})

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return s.listen(ctx, battleId, stream, func() int64 {
			return time.Now().Unix()
		})
	})

	go func() {
		if err := s.poll(ctx, battleId, stop, ticker.C, stream, time.Now().Unix()); err != nil {
			s.logger.Print(errors.Wrap(err, "Event Poll Error"))
		}
	}()

	if err := eg.Wait(); err != nil {
		s.logger.Print(err)
	}

	stop <- struct{}{}

	return stream.Close()
}
