package main

import (
	"encoding/json"
	"kafka-dashboard/pkg/models"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/gorilla/websocket"
)

const (
	kafkaBrokers = "kafka:9092"
	kafkaTopic   = "raw-data"
	processTopic = "processed-data"
	tickers      = "ethusdt,btcusdt,ltcusdt,xrpusdt,bchusdt,adausdt,linkusdt,dotusdt,uniusdt,filusdt@ticker"
)

func main() {
	// Create Kafka producer
	producer, err := sarama.NewSyncProducer(strings.Split(kafkaBrokers, ","), nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Create WebSocket URL with the specified tickers
	webSocketURL := "wss://stream.binance.com:9443/ws/" + strings.ReplaceAll(tickers, ",", "@ticker/")

	// Connect to Binance WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to Binance WebSocket: %v", err)
	}
	defer conn.Close()

	// Channel to handle OS signals for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		<-sigs
		done <- true
	}()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error reading message from Binance WebSocket: %v", err)
				return
			}

			// Parse the received message
			var ticker models.Ticker
			err = json.Unmarshal(message, &ticker)

			if err != nil {
				log.Printf("Error unmarshalling Binance message: %v", err)
				continue
			}

			// Process the ticker data
			processedTicker, err := processData(message)
			if err != nil {
				log.Printf("Failed to process ticker: %v", err)
				continue
			}

			log.Printf("Price for %v \t %v", ticker.Symbol, ticker.LastPrice)

			tickerJSON, err := json.Marshal(ticker)
			if err != nil {
				log.Printf("Failed to marshal ticker: %v", err)
				continue
			}

			msg := &sarama.ProducerMessage{
				Topic: kafkaTopic,
				Value: sarama.StringEncoder(tickerJSON),
			}
			_, _, err = producer.SendMessage(msg)
			if err != nil {
				log.Printf("Failed to send message to Kafka: %v", err)
			}

			processedTickerJSON, err := json.Marshal(processedTicker)
			if err != nil {
				log.Printf("Failed to marshal processed ticker: %v", err)
				continue
			}

			msg = &sarama.ProducerMessage{
				Topic: processTopic,
				Value: sarama.StringEncoder(processedTickerJSON),
			}
			_, _, err = producer.SendMessage(msg)
			if err != nil {
				log.Printf("Failed to send message to Kafka: %v", err)
			}
		}
	}()

	<-done
	log.Println("Shutting down gracefully...")
	conn.Close()
}

func processData(message []byte) (models.ProcessedTicker, error) {
	var ticker models.Ticker
	err := json.Unmarshal(message, &ticker)
	if err != nil {
		log.Printf("Error unmarshalling Binance message: %v", err)
		return models.ProcessedTicker{}, err
	}

	lastPrice, err := strconv.ParseFloat(ticker.LastPrice, 64)
	if err != nil {
		log.Printf("Failed to parse last price: %v", err)
		return models.ProcessedTicker{}, err
	}

	change, err := strconv.ParseFloat(ticker.Change, 64)
	if err != nil {
		log.Printf("Failed to parse change: %v", err)
		return models.ProcessedTicker{}, err
	}

	changePct, err := strconv.ParseFloat(ticker.ChangePct, 64)
	if err != nil {
		log.Printf("Failed to parse change percentage: %v", err)
		return models.ProcessedTicker{}, err
	}

	volume, err := strconv.ParseFloat(ticker.Volume, 64)
	if err != nil {
		log.Printf("Failed to parse volume: %v", err)
		return models.ProcessedTicker{}, err
	}

	processedTicker := models.ProcessedTicker{
		Symbol:    strings.Replace(ticker.Symbol, "USDT", "", 1),
		Name:      ticker.Symbol,
		Price:     lastPrice,
		Change:    change,
		ChangePct: changePct,
		Volume:    volume,
	}

	return processedTicker, nil
}
