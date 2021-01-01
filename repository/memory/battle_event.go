package memory

import "shiritori/values"

type BattleEventRepository struct {
	store map[string][]values.BattleEvent
}

func NewBattleEventRepository() *BattleEventRepository {
	return &BattleEventRepository{
		store: map[string][]values.BattleEvent{},
	}
}

func (repo *BattleEventRepository) Insert(event values.BattleEvent) error {
	repo.store[event.BattleID] = append(repo.store[event.BattleID], event)
	return nil
}

func (repo *BattleEventRepository) GetNewer(battleId string, timestamp int64) ([]values.BattleEvent, error) {
	var res []values.BattleEvent

	if events, ok := repo.store[battleId]; ok {
		for _, e := range events {
			if e.Timestamp > timestamp {
				res = append(res, e)
			}
		}
	}

	return res, nil
}
