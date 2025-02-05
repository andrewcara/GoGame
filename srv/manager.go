package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// In development, you might want to allow all origins
		// For production, you should strictly define allowed origins
		return true // Or use a more restrictive check like below:
		// origin := r.Header.Get("Origin")
		// return origin == "http://127.0.0.1:5501" || origin == "http://localhost:5501"
	},
}

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

type Room struct {
	ID      uuid.UUID
	clients []*Client
}

type Manager struct {
	clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
	rooms    map[uuid.UUID]Room
}

func NewManager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
		rooms:    make(map[uuid.UUID]Room),
	}
	m.setupEventHandlers()
	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventCreateRoom] = CreateRoomEvent
	m.handlers[EventJoinRoom] = JoinRoomEvent
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	// Check if Handler is present in Map
	if handler, ok := m.handlers[event.Type]; ok {
		// Execute the handler and return any err
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {

	log.Println("New connection")
	// Begin by upgrading the HTTP request
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Create New Client
	client := NewClient(conn, m)
	// Add the newly created client to the manager
	m.addClient(client)
	// Start the read / write processes
	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *Manager) addRoom(client *Client) {
	m.Lock()
	defer m.Unlock()
	room_clients := make([]*Client, 0)
	room_clients = append(room_clients, client)
	new_room := Room{ID: client.roomId, clients: room_clients}
	m.rooms[client.roomId] = new_room

}

func (m *Manager) removeRoom(ID uuid.UUID) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.rooms[ID]; ok {

		delete(m.rooms, ID)
	}
}

func (m *Manager) addUserToRoom(client *Client, id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()

	room, ok := m.rooms[id]
	if !ok {
		return fmt.Errorf("room not found")
	}

	if len(room.clients) == 1 && client.roomId != id {
		room.clients = append(room.clients, client)
		client.roomId = id
		m.rooms[id] = room // Update the room in the map
		return nil
	}

	return fmt.Errorf("could not join room")
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()

		delete(m.clients, client)
	}
}
