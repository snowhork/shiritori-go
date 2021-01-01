package shiritoriapi

import (
	"context"
	"shiritori/values"

	"github.com/pkg/errors"
)

type BattleActionListener struct {
	receive BattleActionReceive
	handler BattleActionHandler
}

type BattleActionReceive func() (values.BattleAction, error)

type BattleActionHandler interface {
	Handle(action values.BattleAction) error
}

var EmptyBattleEventError = errors.New("Empty Event")

func NewBattleActionListener(handler BattleActionHandler, receive BattleActionReceive) *BattleActionListener {
	return &BattleActionListener{
		handler: handler,
		receive: receive,
	}
}

func (lis *BattleActionListener) Listen(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			action, err := lis.receive()
			if err != nil {
				if err == EmptyBattleEventError {
					continue
				}

				return err
			}

			if err := lis.handler.Handle(action); err != nil {
				return errors.Wrap(err, "Action Handle Error")
			}
		}
	}
}
