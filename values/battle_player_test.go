package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BattlePlayerList_NextPlayer(t *testing.T) {
	p1 := BattlePlayer{UserID: UserID("1"), PlayerNumber: BattlePlayerNumber(1)}
	p2 := BattlePlayer{UserID: UserID("2"), PlayerNumber: BattlePlayerNumber(2)}
	list := NewBattlePlayerList([]BattlePlayer{p1, p2})

	p1Next, _ := list.NextPlayer(p1.PlayerNumber)
	assert.Equal(t, p2, p1Next)

	p2Next, _ := list.NextPlayer(p2.PlayerNumber)
	assert.Equal(t, p1, p2Next)
}
