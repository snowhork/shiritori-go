package shiritoriapi

import (
	"context"
	"errors"
	"log"
	"os"
	"shiritori/gen/shiritori"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Battle(t *testing.T) {

	cases := []struct {
		name     string
		payload  *shiritori.Battlestreamingpayload
		expected *shiritori.Battlestreamingresult
	}{
		{"this is mock",
			&shiritori.Battlestreamingpayload{
				Type: "message",
				MessagePayload: &shiritori.MessagePayload{
					Message: "Hello, world",
				},
			}, &shiritori.Battlestreamingresult{
				Type: "message",
				MessagePayload: &shiritori.MessagePayload{
					Message: "Hello, world",
				},
			},
		},
	}

	for _, c := range cases {
		ctx := context.Background()
		srvc := &shiritorisrvc{logger: log.New(os.Stderr, "", 0), repo: NewMemoryRepositoryFactory()}

		stream := &MockStream{payload: c.payload}

		_ = srvc.Battle(ctx, &shiritori.BattlePayload{BattleID: "1234"}, stream)

		assert.Equal(t, c.expected.Type, stream.result.Type)
		assert.Equal(t, c.expected.MessagePayload.Message, stream.result.MessagePayload.Message)
	}
}

type MockStream struct {
	payload *shiritori.Battlestreamingpayload
	result  *shiritori.Battlestreamingresult
}

func (m *MockStream) Send(res *shiritori.Battlestreamingresult) error {
	if res.Type != "init" {
		m.result = res
	}
	return nil
}

func (m *MockStream) Recv() (*shiritori.Battlestreamingpayload, error) {
	if m.result != nil {
		return nil, errors.New("skip")
	}

	return m.payload, nil
}

func (*MockStream) Close() error {
	return nil
}
