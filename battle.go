package shiritoriapi

import (
	"context"
	"io"
	"shiritori/gen/shiritori"
	"shiritori/values"
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

	pub := NewBattleEventPublisher(
		s.repo.BattleEvent,
		func() (values.BattleEvent, error) {
			payload, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					return values.BattleEvent{}, errors.Wrap(err, "Message Receive Error")
				}
			}

			return parseStreamingPayloadToEntity(battleId, payload, time.Now().Unix())
		})

	sub := NewBattleEventSubscriber(ticker.C, time.Now().Unix(), battleId, s.repo.BattleEvent,
		func(event values.BattleEvent) error {
			res, err := convertEntityToStreamingResult(event)
			if err != nil {
				return err
			}

			return stream.Send(res)
		})

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := pub.Publish(ctx); err != nil {
			return errors.Wrap(err, "Publish Error")
		}

		return nil
	})
	eg.Go(func() error {
		if err := sub.Subscribe(ctx); err != nil {
			return errors.Wrap(err, "Subscribe Error")
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		s.logger.Print(err)
	}

	return stream.Close()
}

func parseStreamingPayloadToEntity(battleId string, payload *shiritori.Battlestreamingpayload, timestamp int64) (values.BattleEvent, error) {
	switch payload.Type {
	case "message":
		return values.NewBattleEventMessage(battleId, timestamp, payload.MessagePayload.Message), nil
	case "close":
		return values.BattleEvent{}, EmptyBattleEventError
	}

	return values.BattleEvent{}, errors.New("UnKnown payload Type")
}

func convertEntityToStreamingResult(event values.BattleEvent) (*shiritori.Battlestreamingresult, error) {
	switch event.Type {
	case values.EventType_Message:
		if event.MessagePayload == nil {
			return nil, errors.New("Message Payload must not be empty")
		}

		return &shiritori.Battlestreamingresult{
			Timestamp: event.Timestamp,
			Type:      "message",
			MessagePayload: &shiritori.MessagePayload{
				Message: event.MessagePayload.Message,
			},
		}, nil
	default:
		return nil, errors.Errorf("Unknown event.Type: %s", event.Type)
	}
}
