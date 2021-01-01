package values

type BattleEvent struct {
	Type      BattleEventType
	Timestamp BattleEventTimestamp
	BattleID  BattleID

	MessagePayload *BattleEventMessagePayload
}

type BattleEventType string
type BattleEventTimestamp int64

type BattleEventMessagePayload struct {
	Message string
}

const EventType_Message = BattleEventType("message")

func NewBattleEventMessage(battleId BattleID, timestamp BattleEventTimestamp, message string) BattleEvent {
	return BattleEvent{
		Type:      EventType_Message,
		BattleID:  battleId,
		Timestamp: timestamp,
		MessagePayload: &BattleEventMessagePayload{
			Message: message,
		},
	}
}
