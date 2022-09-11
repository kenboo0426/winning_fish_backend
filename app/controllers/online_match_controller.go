package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"winning_fish_backend/app/models"

	"github.com/gorilla/websocket"
	// "golang.org/x/net/websocket"
)

func HandleOnlineMatchRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))

	switch r.Method {
	case http.MethodPost:
		createOnlineMatch(w, r)
	case http.MethodOptions:
		// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

func HandleOnlineMatchUpdateRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))

	switch r.Method {
	case http.MethodGet:
		showOnlineMatch(w, r)
	case http.MethodPut:
		updateOnlineMatch(w, r)
	case http.MethodOptions:
		// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var (
	wsChan  = make(chan models.WsPayload)
	clients = make(map[models.WebSocketConnection]string)
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("ok client connectiong")

	var response models.WsJsonResponse
	// response.Message = `<li>Connected to server</li>`

	conn := models.WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go ListenFowWs(&conn)
}

func ListenFowWs(conn *models.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload models.WsPayload

	for {
		err := conn.ReadJSON(&payload)

		if err != nil {
			log.Println(err)
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func broadcastToAll(response models.WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println(err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}

func ListenToWsChannel() {
	var response models.WsJsonResponse

	for {
		e := <-wsChan

		switch e.Action {
		case "join_online_match":
			clients[e.Conn] = e.UserID
			user_ids := getUserList()
			response.Action = "list_users"
			response.JoinedOnlineMatchUserIDs = user_ids
		case "left":
			delete(clients, e.Conn)
			user_ids := getUserList()
			response.Action = "list_users"
			response.JoinedOnlineMatchUserIDs = user_ids
		}
		fmt.Println(response, "resposne")

		broadcastToAll(response)
	}
}

func getUserList() []string {
	var clientList []string
	for _, client := range clients {
		if client != "" {
			clientList = append(clientList, client)
		}
	}

	sort.Strings(clientList)
	return clientList
}

// net/webscokets
// func handleWebSocket(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("sss")
// 	websocket.Handler(func(ws *websocket.Conn) {
// 		defer ws.Close()

// 		err := websocket.Message.Send(ws, "Server: Hello!")
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		for {
// 			msg := ""
// 			err = websocket.Message.Receive(ws, &msg)
// 			if err != nil {
// 				log.Fatalln(err)
// 			}

// 			err := websocket.Message.Send(ws, fmt.Sprintf("server: \"%s\" received", msg))

// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 		}
// 	}).ServeHTTP(w, r)
// 	// return nil
// }

func indexOnlineMatch(w http.ResponseWriter, r *http.Request) {

}

func createOnlineMatch(w http.ResponseWriter, r *http.Request) {

}

func showOnlineMatch(w http.ResponseWriter, r *http.Request) {

}

func updateOnlineMatch(w http.ResponseWriter, r *http.Request) {

}

func startOnlineMatch(w http.ResponseWriter, r *http.Request) {

}
