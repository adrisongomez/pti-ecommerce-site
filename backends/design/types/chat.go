package types

import (
	. "goa.design/goa/v3/dsl"
)

var ChatSession = Type("ChatSession", func() {
	Attribute("id", Int, fieldID)
	Attribute("messages", ArrayOf(ChatMessage))
	Attribute("createdAt", String, fieldDatetime)
	Attribute("updatedAt", String, fieldDatetime)
	Required("id", "messages", "createdAt")
})

var ChatMessage = Type("ChatMessage", func() {
	Attribute("id", Int, fieldID)
	Attribute("message", String)
	Attribute("source", ChatSource)
	Attribute("createdAt", String, fieldDatetime)
	Required("id", "message", "source", "createdAt")
})

var ChatSource = Type("ChatSource", String, func() {
	Description("Type of source of message")
	Enum("Assistant", "User")
})
