package shiritoriapi

import (
	"context"
	"io"
	"shiritori/entity"
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
	stop := make(chan bool)

	msgReciever, ctx := errgroup.WithContext(ctx)
	msgReciever.Go(func() error {
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

				if err := s.process(ctx, battleId, payload); err != nil {
					return errors.Wrap(err, "Message Process Error")
				}
			}
		}
	})

	go func() {
		lastTime := time.Now().Unix()
		for {
			select {
			case <-stop:
				break
			case <-ticker.C:
				nextTime, err := s.poll(ctx, battleId, lastTime, stream)
				if err != nil {
					s.logger.Println(errors.Wrap(err, "Message Poll Error"))
					continue
				}

				lastTime = nextTime
			}
		}
	}()

	if err := msgReciever.Wait(); err != nil {
		s.logger.Print(err)
	}

	stop <- true

	return stream.Close()
}

func (s *shiritorisrvc) process(ctx context.Context, battleId string, payload *shiritori.Battlestreamingpayload) error {
	var data *entity.BattleEvent

	switch payload.Type {
	case "message":
		data = &entity.BattleEvent{
			Timestamp: time.Now().Unix(),
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

func (s *shiritorisrvc) poll(ctx context.Context, battleId string, current int64, stream shiritori.BattleServerStream) (int64, error) {
	events, err := s.repo.BattleEvent.GetNewer(battleId, current)
	if err != nil {
		return 0, errors.Wrap(err, "Get BattleEvent Error")
	}

	for _, event := range events {
		var streamingResult *shiritori.Battlestreamingresult

		switch event.Type {
		case entity.EventType_Message:
			if event.MessagePayload == nil {
				s.logger.Println(errors.New("Message Payload must not be empty"))
				continue
			}

			streamingResult = &shiritori.Battlestreamingresult{
				Timestamp: event.Timestamp,
				Type:      "message",
				MessagePayload: &shiritori.MessagePayload{
					Message: event.MessagePayload.Message,
				},
			}
		}

		if event.Timestamp > current {
			current = event.Timestamp
		}

		if err := stream.Send(streamingResult); err != nil {
			return 0, errors.Wrap(err, "Message Send Error")
		}
	}

	return current, nil
}
