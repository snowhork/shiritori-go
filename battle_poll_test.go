package shiritoriapi

import (
	"context"
	"log"
	"os"
	"shiritori/entity"
	"shiritori/gen/shiritori"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_poll(t *testing.T) {
	mainBattleId := "1234"

	cases := []struct {
		name     string
		entities []entity.BattleEvent
		expected []shiritori.Battlestreamingresult
	}{
		{"pass single event",
			[]entity.BattleEvent{{Type: entity.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &entity.MessagePayload{Message: "Hello, world"}}},
			[]shiritori.Battlestreamingresult{{Type: "message", Timestamp: 2, MessagePayload: &shiritori.MessagePayload{Message: "Hello, world"}}},
		},
		{"sort multi events",
			[]entity.BattleEvent{
				{Type: entity.EventType_Message, Timestamp: 3, BattleID: mainBattleId, MessagePayload: &entity.MessagePayload{Message: "Hello, world"}},
				{Type: entity.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &entity.MessagePayload{Message: "Hello, world"}},
			},
			[]shiritori.Battlestreamingresult{
				{Type: "message", Timestamp: 2, MessagePayload: &shiritori.MessagePayload{Message: "Hello, world"}},
				{Type: "message", Timestamp: 3, MessagePayload: &shiritori.MessagePayload{Message: "Hello, world"}},
			},
		},
		{"filter old events",
			[]entity.BattleEvent{
				{Type: entity.EventType_Message, Timestamp: -1, BattleID: mainBattleId, MessagePayload: &entity.MessagePayload{Message: "Hello, world"}},
				{Type: entity.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &entity.MessagePayload{Message: "Hello, world"}},
			},
			[]shiritori.Battlestreamingresult{
				{Type: "message", Timestamp: 2, MessagePayload: &shiritori.MessagePayload{Message: "Hello, world"}},
			},
		},
		{"ignore other battle events",
			[]entity.BattleEvent{
				{Type: entity.EventType_Message, Timestamp: 2, BattleID: mainBattleId, MessagePayload: &entity.MessagePayload{Message: "Hello, world"}},
				{Type: entity.EventType_Message, Timestamp: 3, BattleID: "unknown", MessagePayload: &entity.MessagePayload{Message: "Hello, world"}},
			},
			[]shiritori.Battlestreamingresult{
				{Type: "message", Timestamp: 2, MessagePayload: &shiritori.MessagePayload{Message: "Hello, world"}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			srvc := &shiritorisrvc{logger: log.New(os.Stderr, "", 0), repo: NewMemoryRepositoryFactory()}

			for i := range c.entities {
				_ = srvc.repo.BattleEvent.Insert(c.entities[i])
			}

			ticker := make(chan time.Time)
			stopper := make(chan struct{})
			stream := &mockSendableStream{}

			go func() {
				_ = srvc.poll(ctx, mainBattleId, stopper, ticker, stream, 1)
			}()
			ticker <- time.Now()
			stopper <- struct{}{}

			assert.Equal(t, c.expected, stream.sent)
		})
	}
}

type mockSendableStream struct {
	sent []shiritori.Battlestreamingresult
}

func (m *mockSendableStream) Send(res *shiritori.Battlestreamingresult) error {
	m.sent = append(m.sent, *res)
	return nil
}
