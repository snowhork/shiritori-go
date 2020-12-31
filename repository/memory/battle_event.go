package memory

import "shiritori/entity"

type BattleEventRepository struct {
	store map[string][]*entity.BattleEvent
}

func NewBattleEventRepository() *BattleEventRepository {
	return &BattleEventRepository{
		store: map[string][]*entity.BattleEvent{},
	}
}

func (repo *BattleEventRepository) Insert(event *entity.BattleEvent) error {
	repo.store[event.BattleID] = append(repo.store[event.BattleID], event)
	return nil
}

func (repo *BattleEventRepository) GetNewer(battleId string, timestamp int64) ([]entity.BattleEvent, error) {
	var res []entity.BattleEvent

	if events, ok := repo.store[battleId]; ok {
		for _, e := range events {
			if e.Timestamp > timestamp {
				res = append(res, *e)
			}
		}
	}

	return res, nil
}
