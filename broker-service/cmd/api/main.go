package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "8080"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {

	// try to connect to rabbitmq

	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	app := Config{
		Rabbit: rabbitConn,
	}

	log.Printf("Starting broker server on port %s", webPort)
	fmt.Println("HELOOOOOOOOOOOOOOOOOOOO")

	srv := *&http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connect() (*amqp.Connection, error) {

	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	// dont continue until rabbit is ready

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbimq")

		if err != nil {
			fmt.Println("Rabbitmq is not ready")
			counts++
		} else {
			log.Println("Connected to Rabbitmq")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backking of...")

		time.Sleep(backoff)
	}

	return connection, nil
}
