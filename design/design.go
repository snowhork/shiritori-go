package design

import . "goa.design/goa/v3/dsl"

// API describes the global properties of the API server.
var _ = API("shiritori", func() {
	Title("Shiritori Service")
	Description("HTTP service for adding numbers, a goa teaser")
	Server("shiritori", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

// Service describes a service
var _ = Service("shiritori", func() {
	Description("The calc service performs operations on numbers")
	// Method describes a service method (endpoint)
	Method("add", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Attribute("a", Int, "Left operand")
			Attribute("b", Int, "Right operand")
			// Both attributes must be provided when invoking "add"
			Required("a", "b")
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(Int)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/add/{a}/{b}")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})

	Method("battle", func() {
		Payload(func() {
			Attribute("battleId", String)
			Required("battleId")
		})

		StreamingPayload(BattleMessage)
		StreamingResult(BattleEvent)

		HTTP(func() {
			GET("/streams/battles/{battleId}")
			Response(StatusOK)
		})
	})

	Files("/", "./frontend/index.html", func() {
	})
})

var BattleMessage = ResultType("BattleMessage", func() {
	Attributes(func() {
		Attribute("type", String)
		Attribute("msg", String)
		Attribute("data", String)

		Required("type")
	})
})

var BattleEvent = ResultType("BattleEvent", func() {
	Attributes(func() {
		Attribute("battleId", String)
		Attribute("name", String)
		Attribute("param", String)
	})

	View("default", func() {
		Attribute("battleId")
		Attribute("name")
	})

	View("other", func() {
		Attribute("battleId")
		Attribute("name")
		Attribute("param")
	})

})
