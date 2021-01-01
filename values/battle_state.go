package values

type BattleState struct {
	ThemeNumber         ThemeNumber
	ThemeChar           WordChar
	CurrentPlayerNumber BattlePlayerNumber
}

type ThemeNumber struct {
	Length int
	Plus   bool
}

func (s BattleState) Validate(w WordBody) bool {
	return s.ThemeNumber.Validate(w) && s.ThemeChar == w.TopChar()
}

func NewThemeNumber(length int, plus bool) ThemeNumber {
	return ThemeNumber{Length: length, Plus: plus}
}

func (n ThemeNumber) Validate(w WordBody) bool {
	if n.Plus {
		return w.Length() >= n.Length
	}
	return w.Length() == n.Length
}
