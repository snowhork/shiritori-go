package shiritoriapi

import "shiritori/values"

type ActionHandler struct {
	battleID values.BattleID
	repo     *RepositoryFactory
}

func NewActionHandler(battleID values.BattleID, repo *RepositoryFactory) *ActionHandler {
	return &ActionHandler{
		battleID: battleID,
		repo:     repo,
	}
}

func (h *ActionHandler) Handle(action values.BattleAction) error {
	switch action.Type {
	case values.BattleActionType_Message:
		return h.handleMessage(action)
	}

	return nil
}

func (h *ActionHandler) handleMessage(action values.BattleAction) error {
	return h.repo.BattleEvent.Insert(values.NewBattleEventMessage(h.battleID, action.Timestamp, action.MessagePayload.Message))
}
