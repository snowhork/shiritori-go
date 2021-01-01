package shiritoriapi

import (
	"context"
	"errors"
	"log"
	"os"
	"shiritori/gen/shiritori"
	"shiritori/values"
	"testing"
	"time"

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

		stream := &MockStream{payload: c.payload, start: time.Now()}

		_ = srvc.Battle(ctx, &shiritori.BattlePayload{BattleID: "1234"}, stream)

		assert.Equal(t, c.expected.Type, stream.result.Type)
		assert.Equal(t, c.expected.MessagePayload.Message, stream.result.MessagePayload.Message)
	}
}

type MockStream struct {
	payload *shiritori.Battlestreamingpayload
	result  *shiritori.Battlestreamingresult
	start   time.Time
}

func (m *MockStream) Send(res *shiritori.Battlestreamingresult) error {
	if res.Type != "init" {
		m.result = res
	}
	return nil
}

func (m *MockStream) Recv() (*shiritori.Battlestreamingpayload, error) {
	time.Sleep(time.Millisecond * 200)

	if time.Since(m.start) >= time.Millisecond*1000 {
		return nil, errors.New("skip")
	}

	return m.payload, nil
}

func (*MockStream) Close() error {
	return nil
}

func Test_parseStreamingPayloadToEntityNoError(t *testing.T) {
	timestamp := values.BattleEventTimestamp(123)
	userID := values.UserID("123")
	cases := []struct {
		name     string
		payload  *shiritori.Battlestreamingpayload
		expected values.BattleAction
	}{
		{"message Type", &shiritori.Battlestreamingpayload{Type: "message", MessagePayload: &shiritori.MessagePayload{Message: "body"}}, values.BattleAction{
			Type:           values.BattleActionType_Message,
			Timestamp:      timestamp,
			UserID:         userID,
			MessagePayload: &values.BattleActionMessagePayload{Message: "body"}},
		},
		{"post_word Type", &shiritori.Battlestreamingpayload{Type: "post_word", PostWordPayload: &shiritori.PostWordPayload{Word: "りんご", Hash: "", Exists: true}}, values.BattleAction{
			Type:            values.BattleActionType_PostWord,
			Timestamp:       timestamp,
			UserID:          userID,
			PostWordPayload: &values.BattleActionPostWordPayload{WordBody: "りんご", WordBodyHash: "", WordExists: true}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := parseStreamingPayloadToEntity(c.payload, timestamp, userID)
			assert.NoError(t, err)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func Test_parseStreamingPayloadToEntityError(t *testing.T) {
	timestamp := values.BattleEventTimestamp(123)
	userID := values.UserID("123")
	cases := []struct {
		name    string
		payload *shiritori.Battlestreamingpayload
	}{
		{"message Type but MessagePayload is nil", &shiritori.Battlestreamingpayload{Type: "message", MessagePayload: nil}},
		{"post_word Type but PostWordPayload is nil", &shiritori.Battlestreamingpayload{Type: "post_word", PostWordPayload: nil}},
		{"post_word Type but Word is invalid", &shiritori.Battlestreamingpayload{Type: "post_word", PostWordPayload: &shiritori.PostWordPayload{Word: "カタカナ", Hash: "", Exists: true}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := parseStreamingPayloadToEntity(c.payload, timestamp, userID)
			assert.Error(t, err)
		})
	}
}
