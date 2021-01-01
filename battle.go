package shiritoriapi

import (
	"context"
	"io"
	"shiritori/gen/shiritori"
	"shiritori/values"
	"time"

	"github.com/gorilla/websocket"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func (s *shiritorisrvc) Battle(ctx context.Context, p *shiritori.BattlePayload, stream shiritori.BattleServerStream) error {
	battleId := values.BattleID(p.BattleID)

	if err := stream.Send(&shiritori.Battlestreamingresult{Type: "init"}); err != nil {
		return stream.Close()
	}

	ticker := time.NewTicker(50 * time.Millisecond)
	handler := NewActionHandler(battleId, s.repo)
	ctx, cancel := context.WithCancel(ctx)

	lis := NewBattleActionListener(
		handler,
		func() (values.BattleAction, error) {
			payload, err := stream.Recv()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					s.logger.Printf("close websocket: %v", err)
					cancel()
					return values.BattleAction{}, nil
				}

				if err != io.EOF {
					return values.BattleAction{}, errors.Wrap(err, "Message Receive Error")
				}
				return values.BattleAction{}, err
			}

			s.logger.Printf("read paylod. Type: %v", payload.Type)

			if payload.MessagePayload.Message == "close" {
				cancel()
			}

			return parseStreamingPayloadToEntity(payload, time.Now().Unix())
		})

	sub := NewBattleEventSubscriber(ticker.C, values.BattleEventTimestamp(time.Now().Unix()), values.BattleID(battleId), s.repo.BattleEvent,
		func(event values.BattleEvent) error {
			res, err := convertEntityToStreamingResult(event)
			if err != nil {
				return err
			}

			return stream.Send(res)
		})

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := lis.Listen(ctx); err != nil {
			return errors.Wrap(err, "Listen Error")
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

func parseStreamingPayloadToEntity(payload *shiritori.Battlestreamingpayload, timestamp int64) (values.BattleAction, error) {
	switch payload.Type {
	case "message":
		return values.NewMessageBattleAction(values.BattleEventTimestamp(timestamp), payload.MessagePayload.Message), nil
	case "close":
		return values.BattleAction{}, EmptyBattleEventError
	}

	return values.BattleAction{}, errors.New("UnKnown payload Type")
}

func convertEntityToStreamingResult(event values.BattleEvent) (*shiritori.Battlestreamingresult, error) {
	switch event.Type {
	case values.EventType_Message:
		if event.MessagePayload == nil {
			return nil, errors.New("Message Payload must not be empty")
		}

		return &shiritori.Battlestreamingresult{
			Timestamp: int64(event.Timestamp),
			Type:      "message",
			MessagePayload: &shiritori.MessagePayload{
				Message: event.MessagePayload.Message,
			},
		}, nil
	default:
		return nil, errors.Errorf("Unknown event.Type: %s", event.Type)
	}
}
