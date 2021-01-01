package shiritoriapi

import (
	"context"
	"shiritori/values"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/stretchr/testify/assert"
)

func Test_BattleActionListener(t *testing.T) {
	cases := []struct {
		name     string
		events   []values.BattleAction
		expected []values.BattleAction
	}{{
		"single message",
		[]values.BattleAction{{Type: values.BattleActionType_Message, Timestamp: 1, MessagePayload: &values.BattleActionMessagePayload{Message: "Hello, world"}}},
		[]values.BattleAction{{Type: values.BattleActionType_Message, Timestamp: 1, MessagePayload: &values.BattleActionMessagePayload{Message: "Hello, world"}}},
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			idx := 0
			handler := &mockHandler{}

			pub := NewBattleActionListener(
				handler,
				func() (event values.BattleAction, e error) {
					if idx == len(c.events) {
						return values.BattleAction{}, EmptyBattleEventError
					}
					idx += 1
					return c.events[idx-1], nil
				})

			g, ctx := errgroup.WithContext(context.Background())
			ctx, cancel := context.WithCancel(ctx)

			g.Go(func() error {
				return pub.Listen(ctx)
			})

			g.Go(func() error {
				time.Sleep(time.Millisecond)
				cancel()
				return nil
			})

			_ = g.Wait()
			assert.Equal(t, c.expected, handler.received)
		})
	}
}

type mockHandler struct {
	received []values.BattleAction
}

func (m *mockHandler) Handle(action values.BattleAction) error {
	m.received = append(m.received, action)
	return nil
}
