package values

import "errors"

type BattlePlayer struct {
	UserID       UserID
	PlayerNumber BattlePlayerNumber
}

type BattlePlayerNumber int

type BattlePlayerList struct {
	Players       []BattlePlayer
	nextPlayerMap map[BattlePlayerNumber]BattlePlayer
}

func NewBattlePlayerList(players []BattlePlayer) BattlePlayerList {
	nextPlayerMap := map[BattlePlayerNumber]BattlePlayer{}

	for i := range players {
		nextPlayerMap[players[i].PlayerNumber] = players[(i+1)%len(players)]
	}

	return BattlePlayerList{
		Players:       players,
		nextPlayerMap: nextPlayerMap,
	}
}

var NextPlayerNotFoundError = errors.New("next player not found(this is fatal bug)")

func (list *BattlePlayerList) NextPlayer(num BattlePlayerNumber) (BattlePlayer, error) {
	if p, ok := list.nextPlayerMap[num]; ok {
		return p, nil
	}

	return BattlePlayer{}, NextPlayerNotFoundError
}
