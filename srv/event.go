package main

import (
	"encoding/json"
	"fmt"

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
	EventCreateRoom       = "create_room"
	EventWaitingForPlayer = "waiting"
	EventJoinRoom         = "join_room"
	EventAboutToStart     = "about_start"
	EventMove             = "move_Received"
	EventGameUpdate       = "game_update"
)

// struct for handling data sent to client after a room is created
type SendCreatedRoom struct {
	ID uuid.UUID `json:"id"`
}
type Move struct {
	Move   string    `json:"move"`
	GameID uuid.UUID `json:"game_id"`
}

type JoinRoomMessage struct {
	ID uuid.UUID `json:"id"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type GameState struct {
	BallPosition    Position `json:"ball_position"`
	Player1Position Position `json:"p1_position"`
	Player2Position Position `json:"p2_position"`
}

// NewMessageEvent is returned when responding to send_message

func MoveEvent(event Event, c *Client) error {
	// m.Lock()
	// defer m.Unlock()
	var moveevent Move

	if err := json.Unmarshal(event.Payload, &moveevent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}
	//TO-Do investigate why movement.roomId is null
	fmt.Println(moveevent.Move, c.roomId)

	room, ok := c.manager.rooms[c.roomId]
	if !ok {
		return fmt.Errorf("room not found")
	}
	if room.Game == nil {
		fmt.Println(room.clients, room.ID, room.Game)
		return fmt.Errorf("game not initialized in room")
	}
	room.Game.mu.Lock()
	defer room.Game.mu.Unlock()

	room.Game.HandleUserInput(c.playerId, moveevent.Move)
	return nil
}

func CreateRoomEvent(event Event, c *Client) error {
	//No need to work with any payload here only assign a new room id to the client and add it to the managers list
	c.roomId = uuid.New()
	c.playerId = 0
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
	c.playerId = 1
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
