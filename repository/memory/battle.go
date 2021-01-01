package memory

import (
	"shiritori/entity"
	"shiritori/values"
)

type BattleRepository struct {
	store map[values.BattleID]*entity.Battle
}

func NewBattleRepository() *BattleRepository {
	return &BattleRepository{
		store: map[values.BattleID]*entity.Battle{},
	}
}

func (repo *BattleRepository) Get(id values.BattleID) (*entity.Battle, error) {
	if battle, ok := repo.store[id]; ok {
		return battle, nil
	} else {
		return nil, nil
	}
}

func (repo *BattleRepository) Upsert(entity *entity.Battle) error {
	repo.store[entity.ID] = entity
	return nil
}
