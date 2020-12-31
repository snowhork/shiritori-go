package shiritoriapi

import (
	"context"
	"io"
	"shiritori/entity"
	"shiritori/gen/shiritori"

	"github.com/pkg/errors"
)

type receivable interface {
	Recv() (*shiritori.Battlestreamingpayload, error)
}

func (s *shiritorisrvc) listen(ctx context.Context, battleId string, stream receivable, timeFunc func() int64) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			payload, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					return errors.Wrap(err, "Message Receive Error")
				}
			}

			if payload.Type == "close" {
				return nil
			}

			if err := s.processPayload(ctx, battleId, payload, timeFunc()); err != nil {
				return errors.Wrap(err, "Message Process Error")
			}
		}
	}
}

func (s *shiritorisrvc) processPayload(ctx context.Context, battleId string, payload *shiritori.Battlestreamingpayload, timestamp int64) error {
	var data entity.BattleEvent

	switch payload.Type {
	case "message":
		data = entity.BattleEvent{
			Timestamp: timestamp,
			BattleID:  battleId,
			Type:      entity.EventType_Message,
			MessagePayload: &entity.MessagePayload{
				Message: payload.MessagePayload.Message,
			},
		}
	case "close":
	}

	return s.repo.BattleEvent.Insert(data)
}
