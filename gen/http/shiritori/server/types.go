// Code generated by goa v3.2.6, DO NOT EDIT.
//
// shiritori HTTP server types
//
// Command:
// $ goa gen shiritori/design

package server

import (
	shiritori "shiritori/gen/shiritori"
	shiritoriviews "shiritori/gen/shiritori/views"

	goa "goa.design/goa/v3/pkg"
)

// BattleStreamingBody is the type of the "shiritori" service "battle" endpoint
// HTTP request body.
type BattleStreamingBody BattlemessageStreamingBody

// BattleResponseBody is the type of the "shiritori" service "battle" endpoint
// HTTP response body.
type BattleResponseBody struct {
	BattleID *string `form:"battleId,omitempty" json:"battleId,omitempty" xml:"battleId,omitempty"`
	Name     *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// BattleResponseBodyOther is the type of the "shiritori" service "battle"
// endpoint HTTP response body.
type BattleResponseBodyOther struct {
	BattleID *string `form:"battleId,omitempty" json:"battleId,omitempty" xml:"battleId,omitempty"`
	Name     *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Param    *string `form:"param,omitempty" json:"param,omitempty" xml:"param,omitempty"`
}

// BattlemessageStreamingBody is used to define fields on request body types.
type BattlemessageStreamingBody struct {
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	Msg  *string `form:"msg,omitempty" json:"msg,omitempty" xml:"msg,omitempty"`
	Data *string `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// NewBattleResponseBody builds the HTTP response body from the result of the
// "battle" endpoint of the "shiritori" service.
func NewBattleResponseBody(res *shiritoriviews.BattleeventView) *BattleResponseBody {
	body := &BattleResponseBody{
		BattleID: res.BattleID,
		Name:     res.Name,
	}
	return body
}

// NewBattleResponseBodyOther builds the HTTP response body from the result of
// the "battle" endpoint of the "shiritori" service.
func NewBattleResponseBodyOther(res *shiritoriviews.BattleeventView) *BattleResponseBodyOther {
	body := &BattleResponseBodyOther{
		BattleID: res.BattleID,
		Name:     res.Name,
		Param:    res.Param,
	}
	return body
}

// NewAddPayload builds a shiritori service add endpoint payload.
func NewAddPayload(a int, b int) *shiritori.AddPayload {
	v := &shiritori.AddPayload{}
	v.A = a
	v.B = b

	return v
}

// NewBattlePayload builds a shiritori service battle endpoint payload.
func NewBattlePayload(battleID string) *shiritori.BattlePayload {
	v := &shiritori.BattlePayload{}
	v.BattleID = battleID

	return v
}

// NewBattleStreamingBody builds a shiritori service battle endpoint payload.
func NewBattleStreamingBody(body *BattleStreamingBody) *shiritori.Battlemessage {
	v := &shiritori.Battlemessage{
		Msg:  body.Msg,
		Data: body.Data,
	}
	if body.Type != nil {
		v.Type = *body.Type
	}

	return v
}

// ValidateBattleStreamingBody runs the validations defined on
// BattleStreamingBody
func ValidateBattleStreamingBody(body *BattleStreamingBody) (err error) {
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	return
}

// ValidateBattlemessageStreamingBody runs the validations defined on
// BattlemessageStreamingBody
func ValidateBattlemessageStreamingBody(body *BattlemessageStreamingBody) (err error) {
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	return
}
