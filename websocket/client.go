package websocket

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	conn *websocket.Conn
}

// NewWebSocketClient creates a new WebSocket client and connects to the given URL
func NewWebSocketClient(serverURL string) (*WebSocketClient, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return &WebSocketClient{conn: conn}, nil
}

// Subscribe sends a subscription message to the WebSocket server
func (client *WebSocketClient) Subscribe(message string) error {
	err := client.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Printf("Failed to send subscription message: %v", err)
		return err
	}
	return nil
}

// ReceiveMessages listens for incoming messages from the WebSocket server
func (client *WebSocketClient) ReceiveMessages(handler func(messageType int, message []byte)) {
	for {
		messageType, message, err := client.conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		handler(messageType, message)
	}
}

// Close closes the WebSocket connection
func (client *WebSocketClient) Close() error {
	return client.conn.Close()
}
