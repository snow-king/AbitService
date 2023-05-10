package main

import (
	"AbitService/app/models"
	"AbitService/app/service"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	models.ConnectDatabase()
	ensure := make(chan string)
	go server(ensure)
	<-ensure
}
func server(result chan string) {
	ch, conn, err := service.DeclareQueue("RequestRating", false, false, false, false, nil)
	if err != nil {
		log.Println(err.Error())
		result <- err.Error()
		return
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Println(err.Error())
			result <- err.Error()
			return
		}
	}(ch)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
			result <- err.Error()
			return
		}
	}(conn)
	args := make(amqp.Table)
	args["x-max-length"] = int32(2)
	cnResponse, connResponse, err := service.DeclareQueue("RatingAbit", true, false, false, false, args)
	if err != nil {
		log.Println(err.Error())
		result <- err.Error()
		return
	}
	defer func(ch *amqp.Channel) {
		err := cnResponse.Close()
		if err != nil {
			log.Println(err.Error())
			result <- err.Error()
			return
		}
	}(ch)
	defer func(conn *amqp.Connection) {
		err := connResponse.Close()
		if err != nil {
			log.Println(err.Error())
			result <- err.Error()
			return
		}
	}(conn)
	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")
	// If we can consume the message, then we know it published successfully.
	msgs, err := ch.Consume(
		"RequestRating",   // queue
		"ReplyToConsumer", // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		log.Printf("consume fail")
	}
	forever := make(chan bool)
	for m := range msgs {
		num, err := strconv.Atoi(string(m.Body))
		fmt.Println(num)
		list := service.GetList(num)
		js, _ := json.Marshal(list)
		err = cnResponse.Publish(
			"",           // exchange
			"RatingAbit", // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType:   "application/json",
				CorrelationId: getCorrelationId(),
				Timestamp:     time.Now(),
				Body:          js,
			})
		if err != nil {
			log.Printf("publish fail")
		}
		err = ch.Ack(m.DeliveryTag, false)
		if err != nil {
			log.Printf("ack fail")
		}
	}
	<-forever
}
func getCorrelationId() string {
	return strconv.Itoa(rand.Intn(9999999999))
}
