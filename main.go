package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"twitch/chat"
	"twitch/faceit"
	"twitch/stream"
	"twitch/types"
	"twitch/utils"

	"github.com/joho/godotenv"

	"github.com/gorilla/websocket"
)

var dayStats types.Stats
var mounthStats types.Stats
var streamOnline bool
var streamStartedAt time.Time

func startUpdate(stream stream.Stream) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	streamOnline, streamStartedAt = stream.GetStreamStrartingTime()
	fmt.Println(streamStartedAt)

	// Initial stats fetch
	dayStats = faceit.Get_day_stats(streamStartedAt)
	mounthStats = faceit.Get_mounth_stats()

	fmt.Println("Initial day: ", dayStats)
	fmt.Println("Initial mounth: ", mounthStats)

	for {
		select {
		case <-ticker.C:
			streamOnline, streamStartedAt = stream.GetStreamStrartingTime()

			if streamOnline {
				dayStats = faceit.Get_day_stats(streamStartedAt)
			}
			mounthStats = faceit.Get_mounth_stats()

			fmt.Println("Updated day: ", dayStats)
			fmt.Println("Updated mounth: ", mounthStats)
		}
	}
}

func main() {
    godotenv.Load()
	cooldown, err := strconv.Atoi(os.Getenv("COOLDOWN"))
	if err != nil {
		panic(err)
	}

	stream := stream.Stream{
		Streamer_username: os.Getenv("STREAMER_USERNAME"),
		ClientId:          os.Getenv("CLIENT_ID"),
		Auth:              os.Getenv("AUTH"),
	}

	go startUpdate(stream)

	chatConfig := types.ChatConfig{
		IrcServer:       "wss://irc-ws.chat.twitch.tv:443",
		Channel:         "#" + os.Getenv("STREAMER_USERNAME"),
		OauthToken:      "oauth:" + os.Getenv("AUTH"),
		BotUsername:     os.Getenv("BOT_USERNAME"),
		MessageCooldown: time.Duration(cooldown),
	}

	chat := chat.NewChat(chatConfig)

	// Start the WebSocket message reading loop
	for {
		// Read a message from the WebSocket
		_, msg, err := chat.Con.ReadMessage()
		if err != nil {
			log.Println("Connection lost. Attempting to reconnect...")
			chat.Reconnect()
			continue
		}

		// Print the raw message to see the format
		fmt.Println("Raw Message:", string(msg))

		// Handle PING/PONG for connection keep-alive
		if string(msg)[:4] == "PING" {
			// Respond to PING to keep the connection alive
			err := chat.Con.WriteMessage(websocket.TextMessage, []byte("PONG "+string(msg[5:])+"\r\n"))
			if err != nil {
				log.Printf("Error responding to PING: %v", err)
			}
		}

		// Parse the message and extract only the text
		if len(msg) > 0 && msg[0] == ':' {
			// Example message: ":username!username@username.tmi.twitch.tv PRIVMSG #channel :Hello, world!"
			// Split the message into parts based on the spaces
			parts := strings.SplitN(string(msg), " :", 2)
			if len(parts) > 1 {
				// The actual message is in parts[1]
				textMessage := strings.TrimSpace(parts[1]) // Trim spaces from the message
				fmt.Println("Received Message:", textMessage)

				// If the received message is "!stats", send a response
				if textMessage == "!stats Day" {
					streamOnline, _ := stream.GetStreamStrartingTime()
					if streamOnline {
						chat.SendMessage(utils.FormatStatsMessage(utils.ParseUsername(string(msg)), dayStats))
					} else {
						chat.SendMessage(fmt.Sprintf("%s :( Stream is offline", utils.ParseUsername(string(msg))))
					}
				}
				if textMessage == "!stats Month" {
					chat.SendMessage(utils.FormatStatsMessage(utils.ParseUsername(string(msg)), mounthStats))
				}
			}
		}
	}
}
