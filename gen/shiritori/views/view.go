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

// Wordresult is the viewed result type that is projected based on a view.
type Wordresult struct {
	// Type to project
	Projected *WordresultView
	// View to render
	View string
}

// Battlestreamingresult is the viewed result type that is projected based on a
// view.
type Battlestreamingresult struct {
	// Type to project
	Projected *BattlestreamingresultView
	// View to render
	View string
}

// WordresultView is a type that runs validations on a projected type.
type WordresultView struct {
	Word   *string
	Exists *bool
	Hash   *string
}

// BattlestreamingresultView is a type that runs validations on a projected
// type.
type BattlestreamingresultView struct {
	Type           *string
	Timestamp      *int64
	MessagePayload *MessagePayloadView
}

// MessagePayloadView is a type that runs validations on a projected type.
type MessagePayloadView struct {
	Message *string
}

var (
	// WordresultMap is a map of attribute names in result type Wordresult indexed
	// by view name.
	WordresultMap = map[string][]string{
		"default": []string{
			"word",
			"exists",
			"hash",
		},
	}
	// BattlestreamingresultMap is a map of attribute names in result type
	// Battlestreamingresult indexed by view name.
	BattlestreamingresultMap = map[string][]string{
		"default": []string{
			"type",
			"timestamp",
			"message_payload",
		},
	}
)

// ValidateWordresult runs the validations defined on the viewed result type
// Wordresult.
func ValidateWordresult(result *Wordresult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateWordresultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateBattlestreamingresult runs the validations defined on the viewed
// result type Battlestreamingresult.
func ValidateBattlestreamingresult(result *Battlestreamingresult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateBattlestreamingresultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateWordresultView runs the validations defined on WordresultView using
// the "default" view.
func ValidateWordresultView(result *WordresultView) (err error) {
	if result.Word == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("word", "result"))
	}
	if result.Hash == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("hash", "result"))
	}
	if result.Exists == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("exists", "result"))
	}
	return
}

// ValidateBattlestreamingresultView runs the validations defined on
// BattlestreamingresultView using the "default" view.
func ValidateBattlestreamingresultView(result *BattlestreamingresultView) (err error) {
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "result"))
	}
	if result.MessagePayload != nil {
		if err2 := ValidateMessagePayloadView(result.MessagePayload); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateMessagePayloadView runs the validations defined on
// MessagePayloadView.
func ValidateMessagePayloadView(result *MessagePayloadView) (err error) {
	if result.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "result"))
	}
	return
}
