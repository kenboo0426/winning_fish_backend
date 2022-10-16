package models

import (
	"github.com/gorilla/websocket"
)

type WsUser struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	RemainedTime *float32 `json:"remained_time"`
	Icon         string   `json:"icon"`
}

type WsResponse struct {
	Action                   string   `json:"action"`
	// JoinedOnlineMatchUserIDs []string `json:"joined_onine_match_user_ids"`
	Users                    []WsUser `json:"users"`
	OnlineMatchID            int      `json:"online_match_id"`
}

type WsConnection struct {
	*websocket.Conn
}

type WsRequest struct {
	Action        string       `json:"action"`
	UserID        string       `json:"user_id"`
	UserName      string       `json:"user_name"`
	RemainedTime  float32      `json:"remained_time"`
	UserIcon      string       `json:"user_icon"`
	OnlineMatchID int          `json:"online_match_id"`
	Conn          WsConnection `json:"-"`
}

type WsClient struct {
	OnlineMatchID int
	WsUser        WsUser
}
