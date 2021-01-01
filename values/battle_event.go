package values

type BattleEvent struct {
	Type      BattleEventType
	Timestamp int64
	BattleID  string

	MessagePayload *MessagePayload
}

type BattleEventType string
type BattleEventTimestamp int64

type MessagePayload struct {
	Message string
}

const EventType_Message = BattleEventType("message")

func NewBattleEventMessage(battleId string, timestamp int64, message string) BattleEvent {
	return BattleEvent{
		Type:      EventType_Message,
		BattleID:  battleId,
		Timestamp: timestamp,
		MessagePayload: &MessagePayload{
			Message: message,
		},
	}
}
