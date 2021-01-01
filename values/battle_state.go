package values

import (
	"errors"
	"math/rand"
)

type BattleState struct {
	ThemeNumber            ThemeNumber
	ThemeChar              WordChar
	CurrentBattlePlayer    BattlePlayer
	PlayersNumbersQueueMap PlayersNumbersQueueMap
}

func NewBattleState(num ThemeNumber, c WordChar, currentPlayer BattlePlayer, queueMap PlayersNumbersQueueMap) BattleState {
	return BattleState{
		ThemeNumber:            num,
		ThemeChar:              c,
		CurrentBattlePlayer:    currentPlayer,
		PlayersNumbersQueueMap: queueMap,
	}
}

type ThemeNumber struct {
	Length int
	Plus   bool
}

func NewThemeNumber(length int, plus bool) ThemeNumber {
	return ThemeNumber{Length: length, Plus: plus}
}

type ThemeNumbersQueue struct {
	Numbers []ThemeNumber
}

func NewThemeNumbersQueue(nums []ThemeNumber) ThemeNumbersQueue {
	copied := make([]ThemeNumber, len(nums))
	copy(copied, nums)
	return ThemeNumbersQueue{Numbers: copied}
}

func (s BattleState) Validate(w WordBody) bool {
	return s.ThemeNumber.Validate(w) && s.ThemeChar == w.TopChar()
}

func (n ThemeNumber) Validate(w WordBody) bool {
	if n.Plus {
		return w.Length() >= n.Length
	}
	return w.Length() == n.Length
}

var EmptyThemeNumbersQueueError = errors.New("try to sample empty ThemeNumbersQueue")

func (q ThemeNumbersQueue) Sample() (ThemeNumber, []ThemeNumber, error) {
	if len(q.Numbers) == 0 {
		return ThemeNumber{}, nil, EmptyThemeNumbersQueueError
	}

	idx := rand.Intn(len(q.Numbers))

	var res []ThemeNumber
	for i, num := range q.Numbers {
		if i == idx {
			continue
		}
		res = append(res, num)
	}

	return q.Numbers[idx], res, nil
}

type PlayersNumbersQueueMap struct {
	queueMap map[BattlePlayerNumber]ThemeNumbersQueue
}

func NewPlayersNumbersQueueMap(queueMap map[BattlePlayerNumber]ThemeNumbersQueue) PlayersNumbersQueueMap {
	return PlayersNumbersQueueMap{
		queueMap: queueMap,
	}
}

var NotFoundQueueError = errors.New("not found player queue")

func (m PlayersNumbersQueueMap) Get(num BattlePlayerNumber) (ThemeNumbersQueue, error) {
	if q, ok := m.queueMap[num]; ok {
		return q, nil
	}

	return ThemeNumbersQueue{}, NotFoundQueueError
}
