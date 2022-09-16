package models

import (
	"github.com/gorilla/websocket"
)

type WsUser struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	RemainedTime float32 `json:"remained_time"`
}

type WsJsonResponse struct {
	Action                   string      `json:"action"`
	JoinedOnlineMatchUserIDs []string    `json:"joined_onine_match_user_ids"`
	Users                    []WsUser    `json:"users"`
	OnlineMatch              OnlineMatch `json:"online_match"`
}

type WebSocketConnection struct {
	*websocket.Conn
}

type WsPayload struct {
	Action       string              `json:"action"`
	UserID       string              `json:"user_id"`
	UserName     string              `json:"user_name"`
	RemainedTime float32             `json:"remained_time"`
	Conn         WebSocketConnection `json:"-"`
}
