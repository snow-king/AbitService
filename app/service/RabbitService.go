package service

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func DeclareQueue(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (*amqp.Channel, *amqp.Connection, error) {
	url := viper.GetString("RABBIT_URL")
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil || ch == nil {
		return nil, nil, errors.New("error getting channel")
	}
	_, err = ch.QueueDeclare(
		name,       // name
		durable,    // durable
		autoDelete, // delete when unused
		exclusive,  // exclusive
		noWait,     // noWait
		args,       // arguments
	)
	return ch, conn, err
}
