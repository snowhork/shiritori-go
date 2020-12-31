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

	Method("words", func() {
		Payload(func() {
			Attribute("word", String)
			Required("word")
		})

		Result(WordResult)

		HTTP(func() {
			GET("/words/{word}")
			Response(StatusOK)
		})
	})

	Method("battle", func() {
		Payload(func() {
			Attribute("battleId", String)
			Required("battleId")
		})

		StreamingPayload(BattleStreamingPayload)
		StreamingResult(BattleStreamingResult)

		HTTP(func() {
			GET("/streams/battles/{battleId}")
			Response(StatusOK)
		})
	})

	Files("/", "./frontend/index.html", func() {
	})
})

var BattleStreamingPayload = ResultType("BattleStreamingPayload", func() {
	Attributes(func() {
		Attribute("type", String, func() {
			Enum("message", "close")
		})

		Attribute("message_payload", MessagePayload)
		Required("type")
	})
})

var BattleStreamingResult = ResultType("BattleStreamingResult", func() {
	Attributes(func() {
		Attribute("type", String)
		Attribute("timestamp", Int64)

		Attribute("message_payload", MessagePayload)
		Required("type", "timestamp")
	})
})

var MessagePayload = Type("MessagePayload", func() {
	Attribute("message", String)
	Required("message")
})

var WordResult = ResultType("WordResult", func() {
	Attributes(func() {
		Attribute("word", String)
		Attribute("exists", Boolean)
		Attribute("hash", String)

		Required("word", "hash", "exists")
	})
})
