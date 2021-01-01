package shiritoriapi

import (
	"errors"
	"shiritori/values"
)

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

func (h *ActionHandler) Handle(a values.BattleAction) error {
	switch a.Type {
	case values.BattleActionType_Message:
		return h.handleMessage(a.Timestamp, a.UserID, a.MessagePayload)
	case values.BattleActionType_PostWord:
		return h.handlePostWord(a.Timestamp, a.UserID, a.PostWordPayload)
	}

	return nil
}

func (h *ActionHandler) handleMessage(timestamp values.BattleEventTimestamp, userID values.UserID, p *values.BattleActionMessagePayload) error {
	return h.repo.BattleEvent.Insert(values.NewBattleEventMessage(h.battleID, timestamp, p.Message))
}

func (h *ActionHandler) handlePostWord(timestamp values.BattleEventTimestamp, userID values.UserID, p *values.BattleActionPostWordPayload) error {
	if p.WordExists && h.signValidator.Sign(p.WordBody, p.WordExists) == p.WordBodyHash {
		battle, err := h.repo.Battle.Get(h.battleID)
		if err != nil {
			return err
		}

		if battle.State.CurrentBattlePlayer.UserID == userID {
			return errors.New("invalid user turn")
		}

		if err := battle.ChangeStateByPostWord(*p); err != nil {
			return err
		}

		if err := h.repo.Battle.Upsert(battle); err != nil {
			return err

		}
	}

	return nil
}
