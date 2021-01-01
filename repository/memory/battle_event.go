package memory

import "shiritori/values"

type BattleEventRepository struct {
	store map[values.BattleID][]values.BattleEvent
}

func NewBattleEventRepository() *BattleEventRepository {
	return &BattleEventRepository{
		store: map[values.BattleID][]values.BattleEvent{},
	}
}

func (repo *BattleEventRepository) Insert(event values.BattleEvent) error {
	repo.store[event.BattleID] = append(repo.store[event.BattleID], event)
	return nil
}

func (repo *BattleEventRepository) GetNewer(battleId values.BattleID, timestamp values.BattleEventTimestamp) ([]values.BattleEvent, error) {
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
