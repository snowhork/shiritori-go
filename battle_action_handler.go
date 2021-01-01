package shiritoriapi

import "shiritori/values"

type ActionHandler struct {
	battleID      values.BattleID
	repo          *RepositoryFactory
	signValidator WordSigner
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
		return h.handleMessage(action.Timestamp, action.MessagePayload)
	case values.BattleActionType_PostWord:
		return h.handlePostWord(action.Timestamp, action.PostWordPayload)
	}

	return nil
}

func (h *ActionHandler) handleMessage(timestamp values.BattleEventTimestamp, p *values.BattleActionMessagePayload) error {
	return h.repo.BattleEvent.Insert(values.NewBattleEventMessage(h.battleID, timestamp, p.Message))
}

func (h *ActionHandler) handlePostWord(timestamp values.BattleEventTimestamp, p *values.BattleActionPostWordPayload) error {
	if p.WordExists && h.signValidator.Sign(p.WordBody, p.WordExists) == p.WordBodyHash {
		battle, err := h.repo.Battle.Get(h.battleID)
		if err != nil {
			return err
		}

		if battle.ChangeStateByPostWord(*p) {
			if err := h.repo.Battle.Upsert(battle); err != nil {
				return err
			}
		}
	}

	return nil
}
