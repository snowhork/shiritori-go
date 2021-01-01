package entity

import "shiritori/values"

type Battle struct {
	ID         values.BattleID
	Rule       values.BattleRule
	State      values.BattleState
	PlayerList values.BattlePlayerList
}

func NewBattle(id values.BattleID, rule values.BattleRule, state values.BattleState, playerList values.BattlePlayerList) (*Battle, error) {
	return &Battle{
		ID:         id,
		Rule:       rule,
		State:      state,
		PlayerList: playerList,
	}, nil
}

func (b *Battle) ChangeStateByPostWord(p values.BattleActionPostWordPayload) error {
	nextChar := p.WordBody.LastChar()
	nextPlayer, err := b.PlayerList.NextPlayer(b.State.CurrentPlayerNumber)
	nextQueueMap := map[values.BattlePlayerNumber]values.ThemeNumbersQueue{}

	if err != nil {
		return err
	}

	q, err := b.State.PlayersNumbersQueueMap.Get(nextPlayer.PlayerNumber)
	if err != nil {
		return err
	}

	nextNumber, rest, err := q.Sample()
	if err != nil {
		return err
	}

	for _, p := range b.PlayerList.Players {
		if p.PlayerNumber == nextPlayer.PlayerNumber {
			if len(rest) == 0 {
				newList := make([]values.ThemeNumber, len(b.Rule.NumberSet))
				copy(newList, b.Rule.NumberSet)
				nextQueueMap[p.PlayerNumber] = values.NewThemeNumbersQueue(newList)
			} else {
				nextQueueMap[p.PlayerNumber] = values.NewThemeNumbersQueue(rest)
			}

		} else {
			q, err := b.State.PlayersNumbersQueueMap.Get(p.PlayerNumber)
			if err != nil {
				return err
			}
			nextQueueMap[p.PlayerNumber] = q
		}
	}

	b.State = values.NewBattleState(nextNumber, nextChar, nextPlayer.PlayerNumber, values.NewPlayersNumbersQueueMap(nextQueueMap))
	return nil
}
