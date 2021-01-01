package shiritoriapi

import (
	"context"
	"io"
	"shiritori/values"

	"github.com/pkg/errors"
)

type BattleEventPublisher struct {
	repo    BattleEventRepository
	receive BattleEventReceive
}

type BattleEventReceive func() (values.BattleEvent, error)

var EmptyBattleEventError = errors.New("Empty Event")

func NewBattleEventPublisher(repo BattleEventRepository, receive BattleEventReceive) *BattleEventPublisher {
	return &BattleEventPublisher{
		repo:    repo,
		receive: receive,
	}
}

func (pub *BattleEventPublisher) Publish(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			event, err := pub.receive()
			if err != nil {
				if err == EmptyBattleEventError {
					continue
				}

				if err != io.EOF {
					return errors.Wrap(err, "Message Receive Error")
				}
			}

			if err := pub.repo.Insert(event); err != nil {
				return errors.Wrap(err, "Event Insert Error")
			}
		}
	}
}
