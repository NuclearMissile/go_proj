package RBMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

// amqp://[username]:[password]@[sever_address]:[port]/[virtual_host]
const MQURL = "amqp://nuclear:nuclear@localhost:5672/nuclear_vh"

type RBMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	QueueName  string
	Exchange   string
	Key        string
	MQurl      string
}

func NewRBMQ(queueName, exchange, key string) *RBMQ {
	return &RBMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MQurl:     MQURL,
	}
}

func (r *RBMQ) Close() {
	r.channel.Close()
	r.connection.Close()
}

func (r *RBMQ) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func NewRBMQSimple(queueName string) *RBMQ {
	rabbitMQ := NewRBMQ(queueName, "", "")
	var err error
	rabbitMQ.connection, err = amqp.Dial(rabbitMQ.MQurl)
	rabbitMQ.failOnError(err, "fail to establish RBMQ connection")
	rabbitMQ.channel, err = rabbitMQ.connection.Channel()
	rabbitMQ.failOnError(err, "fail to get RBMQ channel")
	return rabbitMQ
}

func (r *RBMQ) PublishSimple(msg string) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

func (r *RBMQ) ReceiveSimple() {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a msg: %s", d.Body)
		}
	}()
	log.Printf("[*] Wait for messages...")
	<-forever
}

func NewRBMQPubSub(exchangeName string) *RBMQ {
	rabbitMQ := NewRBMQ("", exchangeName, "")
	var err error
	rabbitMQ.connection, err = amqp.Dial(rabbitMQ.MQurl)
	rabbitMQ.failOnError(err, "fail to establish RBMQ connection")
	rabbitMQ.channel, err = rabbitMQ.connection.Channel()
	rabbitMQ.failOnError(err, "fail to get RBMQ channel")
	return rabbitMQ
}

func (r *RBMQ) PublishSub(msg string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnError(err, "failed to declare exchange")
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(msg),
		},
	)
	r.failOnError(err, "failed to publish message")
}

func (r *RBMQ) ReceiveSub() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnError(err, "failed to declare exchange")
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnError(err, "failed to declare queue")
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(<-chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received a msg: %s", msg.Body)
		}
	}()
	log.Printf("[*] Wait for messages...")
	<-forever
}

func NewRBMQRouting(exchangeName, routingKey string) *RBMQ {
	rabbitMQ := NewRBMQ("", exchangeName, routingKey)
	var err error
	rabbitMQ.connection, err = amqp.Dial(rabbitMQ.MQurl)
	rabbitMQ.failOnError(err, "fail to establish RBMQ connection")
	rabbitMQ.channel, err = rabbitMQ.connection.Channel()
	rabbitMQ.failOnError(err, "fail to get RBMQ channel")
	return rabbitMQ
}

func (r *RBMQ) PublishRouting(msg string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnError(err, "fail to declare exchange")
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(msg),
		},
	)
}

func (r *RBMQ) ReceiveRouting()  {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnError(err, "failed to declare exchange")
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnError(err, "failed to declare queue")
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received a msg: %s", msg.Body)
		}
	}()
	log.Printf("[*] Wait for messages...")
	<-forever
}