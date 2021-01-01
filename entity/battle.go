package entity

import "shiritori/values"

type Battle struct {
	ID      string
	Rule    values.BattleRule
	State   values.BattleState
	Players []values.BattlePlayer
}

func NewBattle(id string, rule values.BattleRule, state values.BattleState, players []values.BattlePlayer) (*Battle, error) {
	return &Battle{
		ID:      id,
		Rule:    rule,
		State:   state,
		Players: players,
	}, nil
}

func (b *Battle) TransitState(event values.BattleEvent) error {
	return nil
}
