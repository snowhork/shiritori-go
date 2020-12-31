// Code generated by goa v3.2.6, DO NOT EDIT.
//
// shiritori service
//
// Command:
// $ goa gen shiritori/design

package shiritori

import (
	"context"
	shiritoriviews "shiritori/gen/shiritori/views"
)

// The calc service performs operations on numbers
type Service interface {
	// Add implements add.
	Add(context.Context, *AddPayload) (res int, err error)
	// Words implements words.
	Words(context.Context, *WordsPayload) (res *Wordresult, err error)
	// Battle implements battle.
	Battle(context.Context, *BattlePayload, BattleServerStream) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "shiritori"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"add", "words", "battle"}

// BattleServerStream is the interface a "battle" endpoint server stream must
// satisfy.
type BattleServerStream interface {
	// Send streams instances of "Battlestreamingresult".
	Send(*Battlestreamingresult) error
	// Recv reads instances of "Battlestreamingpayload" from the stream.
	Recv() (*Battlestreamingpayload, error)
	// Close closes the stream.
	Close() error
}

// BattleClientStream is the interface a "battle" endpoint client stream must
// satisfy.
type BattleClientStream interface {
	// Send streams instances of "Battlestreamingpayload".
	Send(*Battlestreamingpayload) error
	// Recv reads instances of "Battlestreamingresult" from the stream.
	Recv() (*Battlestreamingresult, error)
	// Close closes the stream.
	Close() error
}

// AddPayload is the payload type of the shiritori service add method.
type AddPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// WordsPayload is the payload type of the shiritori service words method.
type WordsPayload struct {
	Word string
}

// Wordresult is the result type of the shiritori service words method.
type Wordresult struct {
	Word   string
	Exists bool
	Hash   string
}

// BattlePayload is the payload type of the shiritori service battle method.
type BattlePayload struct {
	BattleID string
}

// Battlestreamingpayload is the streaming payload type of the shiritori
// service battle method.
type Battlestreamingpayload struct {
	Type           string
	MessagePayload *MessagePayload
}

// Battlestreamingresult is the result type of the shiritori service battle
// method.
type Battlestreamingresult struct {
	Type           string
	Timestamp      int64
	MessagePayload *MessagePayload
}

type MessagePayload struct {
	Message string
}

// NewWordresult initializes result type Wordresult from viewed result type
// Wordresult.
func NewWordresult(vres *shiritoriviews.Wordresult) *Wordresult {
	return newWordresult(vres.Projected)
}

// NewViewedWordresult initializes viewed result type Wordresult from result
// type Wordresult using the given view.
func NewViewedWordresult(res *Wordresult, view string) *shiritoriviews.Wordresult {
	p := newWordresultView(res)
	return &shiritoriviews.Wordresult{Projected: p, View: "default"}
}

// NewBattlestreamingresult initializes result type Battlestreamingresult from
// viewed result type Battlestreamingresult.
func NewBattlestreamingresult(vres *shiritoriviews.Battlestreamingresult) *Battlestreamingresult {
	return newBattlestreamingresult(vres.Projected)
}

// NewViewedBattlestreamingresult initializes viewed result type
// Battlestreamingresult from result type Battlestreamingresult using the given
// view.
func NewViewedBattlestreamingresult(res *Battlestreamingresult, view string) *shiritoriviews.Battlestreamingresult {
	p := newBattlestreamingresultView(res)
	return &shiritoriviews.Battlestreamingresult{Projected: p, View: "default"}
}

// newWordresult converts projected type Wordresult to service type Wordresult.
func newWordresult(vres *shiritoriviews.WordresultView) *Wordresult {
	res := &Wordresult{}
	if vres.Word != nil {
		res.Word = *vres.Word
	}
	if vres.Exists != nil {
		res.Exists = *vres.Exists
	}
	if vres.Hash != nil {
		res.Hash = *vres.Hash
	}
	return res
}

// newWordresultView projects result type Wordresult to projected type
// WordresultView using the "default" view.
func newWordresultView(res *Wordresult) *shiritoriviews.WordresultView {
	vres := &shiritoriviews.WordresultView{
		Word:   &res.Word,
		Exists: &res.Exists,
		Hash:   &res.Hash,
	}
	return vres
}

// newBattlestreamingresult converts projected type Battlestreamingresult to
// service type Battlestreamingresult.
func newBattlestreamingresult(vres *shiritoriviews.BattlestreamingresultView) *Battlestreamingresult {
	res := &Battlestreamingresult{}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Timestamp != nil {
		res.Timestamp = *vres.Timestamp
	}
	if vres.MessagePayload != nil {
		res.MessagePayload = transformShiritoriviewsMessagePayloadViewToMessagePayload(vres.MessagePayload)
	}
	return res
}

// newBattlestreamingresultView projects result type Battlestreamingresult to
// projected type BattlestreamingresultView using the "default" view.
func newBattlestreamingresultView(res *Battlestreamingresult) *shiritoriviews.BattlestreamingresultView {
	vres := &shiritoriviews.BattlestreamingresultView{
		Type:      &res.Type,
		Timestamp: &res.Timestamp,
	}
	if res.MessagePayload != nil {
		vres.MessagePayload = transformMessagePayloadToShiritoriviewsMessagePayloadView(res.MessagePayload)
	}
	return vres
}

// transformShiritoriviewsMessagePayloadViewToMessagePayload builds a value of
// type *MessagePayload from a value of type *shiritoriviews.MessagePayloadView.
func transformShiritoriviewsMessagePayloadViewToMessagePayload(v *shiritoriviews.MessagePayloadView) *MessagePayload {
	if v == nil {
		return nil
	}
	res := &MessagePayload{
		Message: *v.Message,
	}

	return res
}

// transformMessagePayloadToShiritoriviewsMessagePayloadView builds a value of
// type *shiritoriviews.MessagePayloadView from a value of type *MessagePayload.
func transformMessagePayloadToShiritoriviewsMessagePayloadView(v *MessagePayload) *shiritoriviews.MessagePayloadView {
	if v == nil {
		return nil
	}
	res := &shiritoriviews.MessagePayloadView{
		Message: &v.Message,
	}

	return res
}
