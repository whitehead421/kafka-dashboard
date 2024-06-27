package main

import (
	"encoding/json"
	"kafka-dashboard/pkg/models"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *models.ProcessedTicker)

const (
	kafkaBrokers = "kafka:9092"
	kafkaTopic   = "processed-data"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ws", handleWebSocket)

	go handleMessages()
	go startKafkaConsumer()

	r.Run(":8080")
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	defer delete(clients, conn)

	for {
		var msg models.ProcessedTicker
		err := conn.ReadJSON(&msg)
		log.Printf("Received message: %v", msg)
		if err != nil {
			log.Printf("Error reading from WebSocket: %v", err)
			break
		}
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing to WebSocket: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func startKafkaConsumer() {
	consumer, err := sarama.NewConsumer([]string{kafkaBrokers}, nil)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(kafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case message := <-partitionConsumer.Messages():
			var ticker models.ProcessedTicker
			err := json.Unmarshal(message.Value, &ticker)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			broadcast <- &ticker

		case <-signals:
			return
		}
	}
}
