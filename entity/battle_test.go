package entity

import (
	"shiritori/values"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Battle_ChangeStateByPostWord(t *testing.T) {
	battleID := values.BattleID("123456")

	n1 := values.ThemeNumber{Length: 1, Plus: true}
	n2 := values.ThemeNumber{Length: 2, Plus: true}
	n3 := values.ThemeNumber{Length: 3, Plus: true}

	p1 := values.BattlePlayer{UserID: values.UserID("1"), PlayerNumber: values.BattlePlayerNumber(1)}
	p2 := values.BattlePlayer{UserID: values.UserID("2"), PlayerNumber: values.BattlePlayerNumber(2)}

	battle := Battle{
		ID: battleID,
		Rule: values.BattleRule{
			NumberSet: []values.ThemeNumber{n1, n2, n3},
		},
		PlayerList: values.NewBattlePlayerList([]values.BattlePlayer{p1, p2}),
		State: values.NewBattleState(
			n3,
			values.WordChar("り"),
			p1,
			values.NewPlayersNumbersQueueMap(map[values.BattlePlayerNumber]values.ThemeNumbersQueue{
				values.BattlePlayerNumber(1): values.NewThemeNumbersQueue([]values.ThemeNumber{n1, n2, n3}),
				values.BattlePlayerNumber(2): values.NewThemeNumbersQueue([]values.ThemeNumber{n1}),
			}),
		),
	}

	action := values.NewMessageBattleActionPostWord(values.BattleEventTimestamp(1), values.UserID("1"), values.WordBody("りんご"), values.WordBodyHash(""), true)
	err := battle.ChangeStateByPostWord(*action.PostWordPayload)

	assert.NoError(t, err)
	assert.Equal(t, values.WordChar("ご"), battle.State.ThemeChar)
	assert.Equal(t, n1, battle.State.ThemeNumber)
	assert.Equal(t, p2, battle.State.CurrentBattlePlayer)

	q, _ := battle.State.PlayersNumbersQueueMap.Get(p2.PlayerNumber)
	assert.Equal(t, 3, len(q.Numbers)) // Queue
}
