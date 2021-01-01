package values

type BattleAction struct {
	Type      BattleActionType
	Timestamp BattleEventTimestamp

	MessagePayload *BattleActionMessagePayload
}

type BattleActionType string

const BattleActionType_Message = BattleActionType("message")

type BattleActionMessagePayload struct {
	Message string
}

func NewMessageBattleAction(timestamp BattleEventTimestamp, message string) BattleAction {
	return BattleAction{
		Type:      BattleActionType_Message,
		Timestamp: timestamp,
		MessagePayload: &BattleActionMessagePayload{
			Message: message,
		},
	}
}
