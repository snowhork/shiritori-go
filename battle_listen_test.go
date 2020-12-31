package shiritoriapi

import (
	"context"
	"errors"
	"log"
	"os"
	"shiritori/entity"
	"shiritori/gen/shiritori"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_listen(t *testing.T) {
	mainBattleID := "1234"

	cases := []struct {
		name     string
		payload  []shiritori.Battlestreamingpayload
		expected []entity.BattleEvent
	}{{
		"single message",
		[]shiritori.Battlestreamingpayload{{Type: "message", MessagePayload: &shiritori.MessagePayload{Message: "Hello, world"}}},
		[]entity.BattleEvent{{Type: entity.EventType_Message, Timestamp: 1, BattleID: mainBattleID, MessagePayload: &entity.MessagePayload{Message: "Hello, world"}}},
	}}

	for _, c := range cases {
		srvc := &shiritorisrvc{logger: log.New(os.Stderr, "", 0), repo: NewMemoryRepositoryFactory()}

		stream := &mockReceivableStream{payloads: c.payload}

		_ = srvc.listen(context.Background(), mainBattleID, stream, func() int64 {
			return 1
		})

		actual, _ := srvc.repo.BattleEvent.GetNewer(mainBattleID, 0)
		assert.Equal(t, c.expected, actual)
	}
}

type mockReceivableStream struct {
	payloads []shiritori.Battlestreamingpayload
	it       int
}

func (m *mockReceivableStream) Recv() (*shiritori.Battlestreamingpayload, error) {
	if m.it >= len(m.payloads) {
		return nil, errors.New("skip")
	}

	m.it += 1
	return &m.payloads[m.it-1], nil
}
