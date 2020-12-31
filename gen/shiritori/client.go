// Code generated by goa v3.2.6, DO NOT EDIT.
//
// shiritori client
//
// Command:
// $ goa gen shiritori/design

package shiritori

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "shiritori" service client.
type Client struct {
	AddEndpoint    goa.Endpoint
	WordsEndpoint  goa.Endpoint
	BattleEndpoint goa.Endpoint
}

// NewClient initializes a "shiritori" service client given the endpoints.
func NewClient(add, words, battle goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:    add,
		WordsEndpoint:  words,
		BattleEndpoint: battle,
	}
}

// Add calls the "add" endpoint of the "shiritori" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Words calls the "words" endpoint of the "shiritori" service.
func (c *Client) Words(ctx context.Context, p *WordsPayload) (res *Wordresult, err error) {
	var ires interface{}
	ires, err = c.WordsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Wordresult), nil
}

// Battle calls the "battle" endpoint of the "shiritori" service.
func (c *Client) Battle(ctx context.Context, p *BattlePayload) (res BattleClientStream, err error) {
	var ires interface{}
	ires, err = c.BattleEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(BattleClientStream), nil
}
