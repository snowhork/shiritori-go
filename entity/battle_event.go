package entity

type BattleEvent struct {
	Type      BattleEventType
	Timestamp int64
	BattleID  string

	MessagePayload *MessagePayload
}

type BattleEventType string

type MessagePayload struct {
	Message string
}

const EventType_Message = BattleEventType("message")
