package shiritoriapi

import (
	"context"
	"shiritori/values"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_BattleEventSubscriber(t *testing.T) {
	mainBattleId := "1234"

	cases := []struct {
		name     string
		entities []values.BattleEvent
		expected []values.BattleEvent
	}{
		{"pass single event",
			[]values.BattleEvent{{Type: values.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}}},
			[]values.BattleEvent{{Type: values.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}}},
		},
		{"sort multi events",
			[]values.BattleEvent{
				{Type: values.EventType_Message, Timestamp: 3, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
				{Type: values.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
			},
			[]values.BattleEvent{
				{Type: values.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
				{Type: values.EventType_Message, Timestamp: 3, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
			},
		},
		{"filter old events",
			[]values.BattleEvent{
				{Type: values.EventType_Message, Timestamp: -1, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
				{Type: values.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
			},
			[]values.BattleEvent{
				{Type: values.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
			},
		},
		{"filter other battle events",
			[]values.BattleEvent{
				{Type: values.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
				{Type: values.EventType_Message, Timestamp: 3, BattleID: "unknown", MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
			},
			[]values.BattleEvent{
				{Type: values.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &values.MessagePayload{Message: "Hello, world"}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			ctx, cancel := context.WithCancel(context.Background())
			ticker := make(chan time.Time)

			var actual []values.BattleEvent

			sub := NewBattleEventSubscriber(ticker, 1, mainBattleId, NewMemoryRepositoryFactory().BattleEvent,
				func(event values.BattleEvent) error {
					actual = append(actual, event)
					return nil
				})

			for _, e := range c.entities {
				_ = sub.repo.Insert(e)
			}

			go func() {
				_ = sub.Subscribe(ctx)
			}()
			ticker <- time.Now()
			cancel()

			assert.Equal(t, c.expected, actual)
		})
	}
}
