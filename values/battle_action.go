package values

type BattleAction struct {
	Type      BattleActionType
	Timestamp BattleEventTimestamp

	MessagePayload  *BattleActionMessagePayload
	PostWordPayload *BattleActionPostWordPayload
}

type BattleActionType string

const BattleActionType_Message = BattleActionType("message")
const BattleActionType_PostWord = BattleActionType("post_word")

type BattleActionMessagePayload struct {
	Message string
}

type BattleActionPostWordPayload struct {
	WordBody     WordBody
	WordBodyHash WordBodyHash
	WordExists   bool
	UserID       UserID
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
