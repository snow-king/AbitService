package BrokerRepository

import (
	"AbitService/app/service"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func RequestRating() error {
	ch, conn, err := service.DeclareQueue("RequestRating", false, false, false, false, nil)
	if err != nil {
		return err
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("1"),
	}
	if err := ch.Publish(
		"",              // exchange
		"RequestRating", // queue name
		false,           // mandatory
		false,           // immediate
		message,         // message to publish
	); err != nil {
		return err
	}
	return nil
}
func RatingList() ([]service.RatingList, error) {
	fresh, err := checkFreshness()
	if !fresh {
		fmt.Println("ПОГРУЖАЮСЬ")
		err := RequestRating()
		if err != nil {
			return nil, err
		}
	}
	args := make(amqp.Table)
	args["x-max-length"] = int32(2)
	cnResponse, connResponse, err := service.DeclareQueue("RatingAbit", true, false, false, false, args)
	if err != nil {
		return nil, err
	}
	messages, err := cnResponse.Consume(
		"RatingAbit", // queue name
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // arguments
	)
	defer func(connResponse *amqp.Connection) {
		err := connResponse.Close()
		if err != nil {

		}
	}(connResponse)
	requestChan := make(chan []byte)
	go func() {
		for message := range messages {
			requestChan <- message.Body
		}
	}()
	var body []service.RatingList
	err = json.Unmarshal(<-requestChan, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func checkFreshness() (bool, error) {
	args := make(amqp.Table)
	args["x-max-length"] = int32(2)
	cnResponse, connResponse, err := service.DeclareQueue("RatingAbit", true, false, false, false, args)
	if err != nil {
		return false, err
	}
	inspect, err := cnResponse.QueueInspect("RatingAbit")
	if inspect.Messages != 2 {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	messages, err := cnResponse.Consume(
		"RatingAbit", // queue name
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // arguments
	)
	if err != nil {
		return false, err
	}
	defer func(connResponse *amqp.Connection) {
		err := connResponse.Close()
		if err != nil {

		}
	}(connResponse)

	then := time.Now().Add(time.Duration(-5) * time.Minute)
	requestChan := make(chan bool)
	go func() {
		i := 0
		for message := range messages {
			i++
			if message.Timestamp.Before(then) && i == 1 {
				requestChan <- false
			}
			if i == 2 {
				requestChan <- true
			}
		}
	}()
	fresh := <-requestChan
	return fresh, nil
}
