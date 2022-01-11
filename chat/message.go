// message.go
package main

import (
	"encoding/json"
	"log"
)

const JoinRoomPrivateAction = "join-room-private"
const RoomJoinedAction = "room-joined"
const UserJoinedAction = "user-join"
const UserLeftAction = "user-left"
const SendMessageAction = "send-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"

type Message struct {
	Action  string  `json:"action"`
	Message string  `json:"message"`
	Target  *Room   `json:"target"`
	Sender  *Client `json:"sender"`
}

func (message *Message) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return json
}
