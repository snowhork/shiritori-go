package shiritoriapi

import (
	"context"
	"shiritori/gen/shiritori"
	"shiritori/values"
	"sort"
	"time"

	"github.com/pkg/errors"
)

type sendable interface {
	Send(*shiritori.Battlestreamingresult) error
}

func (s *shiritorisrvc) poll(ctx context.Context, battleId string, stopper <-chan struct{}, ticker <-chan time.Time, stream sendable, startTime int64) error {
	lastTime := startTime

	for {
		select {
		case <-stopper:
			break
		case <-ticker:
			events, err := s.checkEvent(ctx, battleId, lastTime)
			if err != nil {
				s.logger.Println(errors.Wrap(err, "Message Poll Error"))
				continue
			}

			sort.Slice(events, func(i, j int) bool {
				return events[i].Timestamp < events[j].Timestamp
			})

			for _, event := range events {
				if lastTime < event.Timestamp {
					result, err := parseEntityToStreamingResult(&event)
					if err != nil {
						return err
					}

					if err := stream.Send(result); err != nil {
						return errors.Wrap(err, "Message Send Error")
					}

					lastTime = result.Timestamp
				}
			}
		}
	}
}

func (s *shiritorisrvc) checkEvent(ctx context.Context, battleId string, current int64) ([]values.BattleEvent, error) {
	events, err := s.repo.BattleEvent.GetNewer(battleId, current)
	if err != nil {
		return nil, errors.Wrap(err, "Get BattleEvent Error")
	}

	return events, nil
}

func parseEntityToStreamingResult(event *values.BattleEvent) (*shiritori.Battlestreamingresult, error) {
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
