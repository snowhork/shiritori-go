package shiritoriapi

import (
	"context"
	"shiritori/values"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/stretchr/testify/assert"
)

func Test_BattleEventPublisher(t *testing.T) {
	mainBattleID := values.BattleID("1234")

	cases := []struct {
		name     string
		events   []values.BattleEvent
		expected []values.BattleEvent
	}{{
		"single message",
		[]values.BattleEvent{{Type: values.EventType_Message, Timestamp: 1, BattleID: mainBattleID, MessagePayload: &values.MessagePayload{Message: "Hello, world"}}},
		[]values.BattleEvent{{Type: values.EventType_Message, Timestamp: 1, BattleID: mainBattleID, MessagePayload: &values.MessagePayload{Message: "Hello, world"}}},
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			idx := 0

			pub := NewBattleEventPublisher(
				NewMemoryRepositoryFactory().BattleEvent,
				func() (event values.BattleEvent, e error) {
					if idx == len(c.events) {
						return values.BattleEvent{}, EmptyBattleEventError
					}
					idx += 1
					return c.events[idx-1], nil
				})

			g, ctx := errgroup.WithContext(context.Background())
			ctx, cancel := context.WithCancel(ctx)

			g.Go(func() error {
				return pub.Publish(ctx)
			})

			g.Go(func() error {
				time.Sleep(time.Millisecond)
				cancel()
				return nil
			})

			_ = g.Wait()

			actual, _ := pub.repo.GetNewer(mainBattleID, 0)
			assert.Equal(t, c.expected, actual)
		})

	}
}
