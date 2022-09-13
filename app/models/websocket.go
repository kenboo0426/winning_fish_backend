package models

import (
	"github.com/gorilla/websocket"
)

type WsJsonResponse struct {
	Action                   string   `json:"action"`
	Message                  string   `json:"message"`
	JoinedOnlineMatchUserIDs []string `json:"joined_onine_match_user_ids"`
}

type WebSocketConnection struct {
	*websocket.Conn
}

type WsPayload struct {
	Action  string              `json:"action"`
	Message string              `json:"message"`
	UserID  string              `json:"user_id"`
	Conn    WebSocketConnection `json:"-"`
}
