package entity

type BattleEvent struct {
	Timestamp int64
	Type      BattleEventType
	BattleID  string

	MessagePayload *MessagePayload
}

type BattleEventType string

type MessagePayload struct {
	Message string
}

const EventType_Message = BattleEventType("message")
