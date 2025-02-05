package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Event is the Messages sent over the websocket
// Used to differ between different actions
type Event struct {
	// Type is the message type sent
	Type string `json:"type"`
	// Payload is the data Based on the Type
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

const (
	// EventSendMessage is the event name for new chat messages sent
	EventSendMessage      = "send_message"
	EventNewMessage       = "new_message"
	EventCreateRoom       = "create_room"
	EventWaitingForPlayer = "waiting"
	EventJoinRoom         = "join_room"
	EventAboutToStart     = "about_start"
)

// SendMessageEvent is the payload sent in the
// send_message event
type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

// struct for handling data sent to client after a room is created
type SendCreatedRoom struct {
	ID uuid.UUID `json:"id"`
}

type JoinRoomMessage struct {
	ID uuid.UUID `json:"id"`
}

// NewMessageEvent is returned when responding to send_message
type NewMessageEvent struct {
	SendMessageEvent
	Sent time.Time `json:"sent"`
}

func CreateRoomEvent(event Event, c *Client) error {
	//No need to work with any payload here only assign a new room id to the client and add it to the managers list
	c.roomId = uuid.New()
	c.manager.addRoom(c)
	var created_room SendCreatedRoom

	created_room.ID = c.roomId
	//respond with the created room id and with the waiting type
	data, err := json.Marshal(created_room)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}
	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventWaitingForPlayer

	message, _ := json.Marshal(outgoingEvent)
	c.egress <- message

	return nil
}

func JoinRoomEvent(event Event, c *Client) error {
	var joinevent JoinRoomMessage

	if err := json.Unmarshal(event.Payload, &joinevent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}
	//Here we are either
	if err := c.manager.addUserToRoom(c, joinevent.ID); err != nil {
		return err
	}
	var outgoingEvent Event
	outgoingEvent.Payload = nil
	outgoingEvent.Type = EventAboutToStart
	message, _ := json.Marshal(outgoingEvent)

	//We are letting both players in the room know that the game is about to begin
	for _, client := range c.manager.rooms[joinevent.ID].clients {
		client.egress <- message
	}

	return nil

}

// SendMessageHandler will send out a message to all other participants in the chat
func SendMessageHandler(event Event, c *Client) error {
	// Marshal Payload into wanted format
	var chatevent SendMessageEvent
	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	// Prepare an Outgoing Message to others
	var broadMessage NewMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatevent.Message
	broadMessage.From = chatevent.From

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	// Place payload into an Event
	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewMessage
	// Broadcast to all other Clients
	message, _ := json.Marshal(outgoingEvent)
	for client := range c.manager.clients {
		client.egress <- message
	}

	return nil

}
