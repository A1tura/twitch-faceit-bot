package chat

import (
	"fmt"
	"log"
	"os"
	"time"
	"twitch/types"

	"github.com/gorilla/websocket"
)

type Chat struct {
	Config      types.ChatConfig
	Con         *websocket.Conn
	CanRespondAt time.Time
}

func NewChat(config types.ChatConfig) *Chat {
	conn, err := connectToWebSocket(config)
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
		os.Exit(1)
	}

	return &Chat{
		Config: config,
		Con:    conn,
	}
}

func connectToWebSocket(config types.ChatConfig) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(config.IrcServer, nil)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to WebSocket: %v", err)
	}

	// Send authentication details (username and OAuth token)
	err = conn.WriteMessage(websocket.TextMessage, []byte("PASS "+config.OauthToken+"\r\n"))
	if err != nil {
		return nil, fmt.Errorf("Error sending PASS command: %v", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte("NICK "+config.BotUsername+"\r\n"))
	if err != nil {
		return nil, fmt.Errorf("Error sending NICK command: %v", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte("JOIN "+config.Channel+"\r\n"))
	if err != nil {
		return nil, fmt.Errorf("Error sending JOIN command: %v", err)
	}

	return conn, nil
}

func (chat *Chat) SendMessage(message string) {
	// Check if enough time has passed since the last message
	if time.Since(chat.CanRespondAt) < chat.Config.MessageCooldown {
		return
	}

	// Create the PRIVMSG command
	msg := fmt.Sprintf("PRIVMSG %s :%s\r\n", chat.Config.Channel, message)

	// Send the message to the WebSocket server
	err := chat.Con.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Printf("Error sending message: %v", err)
		chat.Reconnect() // Attempt to reconnect if the connection fails
	}
	// Update the time when the bot can respond again
	chat.CanRespondAt = time.Now()
}

func (chat *Chat) Reconnect() {
	// Attempt to reconnect and handle failure
	log.Println("Attempting to reconnect...")
	conn, err := connectToWebSocket(chat.Config)
	if err != nil {
		log.Printf("Error reconnecting: %v", err)
		return
	}
	chat.Con = conn
	log.Println("Reconnected successfully")
}

