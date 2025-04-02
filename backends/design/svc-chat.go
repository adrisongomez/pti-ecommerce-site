package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/securities"
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	. "goa.design/goa/v3/dsl"
)

var _ = Service("chat", func() {
	Description("Chat service")
	HTTP(func() {
		Path("/chats")
	})
	Method("createChatSession", func() {
		Security(securities.JWTAuth, func() {
			Scope(auth.ChatSessionWrite)
		})
		Payload(func() {
			Token("token")
			Required("token")
		})
		Result(Int)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})
	Method("GetSessionById", func() {
		Security(securities.JWTAuth, func() {
			Scope(auth.ChatSessionRead)
		})
		Payload(func() {
			Attribute("sessionId", Int)
			Token("token")
			Required("token")
		})
		Result(types.ChatSession)
		HTTP(func() {
			POST("/{sessionId}")
			Param("sessionId")
			Response(StatusOK)
		})
	})
	Method("submitMessageToSession", func() {
		Security(securities.JWTAuth, func() {
			Scope(auth.ChatSessionWrite)
		})
		Payload(func() {
			Attribute("sessionId", Int)
			Attribute("message", String)
			Token("token")
			Required("token")
		})
		Result(String)
		HTTP(func() {
			POST("/{sessionId}/message")
			Param("sessionId")
			Response(StatusOK)
		})
	})
})
