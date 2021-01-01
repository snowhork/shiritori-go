package values

type BattleAction struct {
	Type      BattleActionType
	Timestamp BattleEventTimestamp
	UserID    UserID

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
}

func NewMessageBattleAction(timestamp BattleEventTimestamp, userID UserID, message string) BattleAction {
	return BattleAction{
		Type:      BattleActionType_Message,
		Timestamp: timestamp,
		UserID:    userID,
		MessagePayload: &BattleActionMessagePayload{
			Message: message,
		},
	}
}

func NewMessageBattleActionPostWord(timestamp BattleEventTimestamp, userID UserID, body WordBody, hash WordBodyHash, exists bool) BattleAction {
	return BattleAction{
		Type:      BattleActionType_PostWord,
		Timestamp: timestamp,
		UserID:    userID,
		PostWordPayload: &BattleActionPostWordPayload{
			WordBody:     body,
			WordBodyHash: hash,
			WordExists:   exists,
		},
	}
}
