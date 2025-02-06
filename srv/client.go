package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	// pongWait is how long we will await a pong response from client
	game_update_interval = 1000 * time.Millisecond
)

// ClientList is a map used to help manage a map of clients
type ClientList map[*Client]bool

// Client is a websocket client, basically a frontend visitor
type Client struct {
	// the websocket connection
	connection *websocket.Conn

	// manager is the manager used to manage the client
	manager *Manager

	egress chan []byte

	roomId uuid.UUID

	playerId int
}

// NewClient is used to initialize a new Client with all required values initialized
func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan []byte),
		roomId:     uuid.Nil,
		playerId:   -1,
	}
}

func (c *Client) readMessages() {

	defer func() {
		// Graceful Close the Connection once this
		// function is done
		c.manager.removeClient(c)
	}()
	// Loop Forever
	for {
		// ReadMessage is used to read the next message in queue
		// in the connection
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break // Break the loop to close conn & Cleanup
		}
		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling message: %v", err)
			break
		}
		// Route the Event

		if err := c.manager.routeEvent(request, c); err != nil {
			log.Println("Error handeling Message: ", err)
		}
	}
}

func (c *Client) writeMessages() {
	defer func() {
		// Graceful close if this triggers a closing
		c.manager.removeClient(c)
	}()

	//for select will read message from channel as soon as it arrives
	//current implementation is to read data from socket then write it back
	for {
		select {
		case message, ok := <-c.egress:
			// Ok will be false Incase the egress channel is closed
			if !ok {
				// Manager has closed this connection channel, so communicate that to frontend
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					// Log that the connection is closed and the reason
					log.Println("connection closed: ", err)
				}
				// Return to close the goroutine
				return
			}
			// Write a Regular text message to the connection
			new_event, _ := json.Marshal(Event{Type: "new_message", Payload: message})
			//c.manager.routeEvent(json.Unmarshal(new_event), c)
			_ = new_event

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println(err)
			}
			log.Println("sent message")
		}

	}
}

func broadcastGameUpdates(room *Room) {
	ticker := time.NewTicker(game_update_interval)
	defer ticker.Stop()
	//TO-DO Add conditions for terminations go function
	for {
		select {
		case <-ticker.C:
			room.Game.mu.Lock()
			room.Game.UpdatePhysics(physicsTickRate)
			// Create game state snapshot quickly
			ball_position := room.Game.world.Objects[2].Shape.GetCenter()
			p1_position := room.Game.world.Objects[0].Shape.GetCenter()
			p2_position := room.Game.world.Objects[1].Shape.GetCenter()

			room.Game.mu.Unlock()
			fmt.Println(ball_position, p1_position, p2_position)
			GameBroadcast := GameState{
				BallPosition:    Position{ball_position.X, ball_position.Y},
				Player1Position: Position{p1_position.X, p1_position.Y},
				Player2Position: Position{p2_position.X, p2_position.Y},
			}

			data, err := json.Marshal(GameBroadcast)
			if err != nil {
				return
			}

			BroadcastEvent := Event{
				Payload: data,
				Type:    EventGameUpdate,
			}

			marshalled_broadcast_event, _ := json.Marshal(BroadcastEvent)

			for _, client := range room.clients {
				client.egress <- marshalled_broadcast_event
			}
		case status := <-room.status:
			if status == FINSIHED {
				break
			}
		}
	}
}
