package entity

import "shiritori/values"

type Battle struct {
	ID      values.BattleID
	Rule    values.BattleRule
	State   values.BattleState
	Players []values.BattlePlayer
}

func NewBattle(id values.BattleID, rule values.BattleRule, state values.BattleState, players []values.BattlePlayer) (*Battle, error) {
	return &Battle{
		ID:      id,
		Rule:    rule,
		State:   state,
		Players: players,
	}, nil
}

func (b *Battle) ChangeState(event values.BattleEvent) (bool, error) {
	return nil
}
