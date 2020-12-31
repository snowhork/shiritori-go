// Code generated by goa v3.2.6, DO NOT EDIT.
//
// shiritori views
//
// Command:
// $ goa gen shiritori/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Battleevent is the viewed result type that is projected based on a view.
type Battleevent struct {
	// Type to project
	Projected *BattleeventView
	// View to render
	View string
}

// BattleeventView is a type that runs validations on a projected type.
type BattleeventView struct {
	BattleID *string
	Name     *string
	Param    *string
}

var (
	// BattleeventMap is a map of attribute names in result type Battleevent
	// indexed by view name.
	BattleeventMap = map[string][]string{
		"default": []string{
			"battleId",
			"name",
		},
		"other": []string{
			"battleId",
			"name",
			"param",
		},
	}
)

// ValidateBattleevent runs the validations defined on the viewed result type
// Battleevent.
func ValidateBattleevent(result *Battleevent) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateBattleeventView(result.Projected)
	case "other":
		err = ValidateBattleeventViewOther(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "other"})
	}
	return
}

// ValidateBattleeventView runs the validations defined on BattleeventView
// using the "default" view.
func ValidateBattleeventView(result *BattleeventView) (err error) {

	return
}

// ValidateBattleeventViewOther runs the validations defined on BattleeventView
// using the "other" view.
func ValidateBattleeventViewOther(result *BattleeventView) (err error) {

	return
}
