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
	userId := values.UserID(p.UserID)

	if err := stream.Send(&shiritori.Battlestreamingresult{Type: "init"}); err != nil {
		return stream.Close()
	}

	ticker := time.NewTicker(50 * time.Millisecond)
	handler := NewActionHandler(battleId, s.repo)
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		ticker.Stop()
	}()

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

			s.logger.Printf("read paylod. Type: %s", payload.Type)

			if payload.MessagePayload.Message == "close" {
				cancel()
			}

			timestamp := values.BattleEventTimestamp(time.Now().Unix())
			res, err := parseStreamingPayloadToEntity(payload, timestamp, userId)
			if err != nil {
				s.logger.Printf("parse payload error: %v", err)
				return values.BattleAction{}, EmptyBattleEventError
			}

			return res, nil
		})

	sub := NewBattleEventSubscriber(ticker.C, values.BattleEventTimestamp(time.Now().Unix()), values.BattleID(battleId), s.repo.BattleEvent,
		func(event values.BattleEvent) error {
			res, err := convertEntityToStreamingResult(event)
			if err != nil {
				s.logger.Printf("convert entity error: %v", err)
				return err
			}

			s.logger.Printf("sent Result paylod. Type: %s", res.Type)
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

func parseStreamingPayloadToEntity(p *shiritori.Battlestreamingpayload, timestamp values.BattleEventTimestamp, userID values.UserID) (values.BattleAction, error) {
	switch p.Type {
	case "message":
		if p.MessagePayload == nil {
			return values.BattleAction{}, errors.New("MessagePayload must not be nil")
		}

		return values.NewMessageBattleAction(timestamp, userID, p.MessagePayload.Message), nil
	case "post_word":
		if p.PostWordPayload == nil {
			return values.BattleAction{}, errors.New("PostWordPayload must not be nil")
		}

		word, err := values.NewWordBody(p.PostWordPayload.Word)
		if err != nil {
			return values.BattleAction{}, err
		}
		return values.NewMessageBattleActionPostWord(timestamp, userID, word, values.WordBodyHash(p.PostWordPayload.Hash), p.PostWordPayload.Exists), nil
	case "close":
		return values.BattleAction{}, EmptyBattleEventError
	}

	return values.BattleAction{}, errors.New("Unknown payload Type")
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
