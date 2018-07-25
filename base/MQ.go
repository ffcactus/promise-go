package base

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"strings"
)

const (
	// PromiseExchange is the exchange name.
	PromiseExchange = "promise_exchange"
	// CreateOperation is create operation.
	CreateOperation = "Create"
	// UpdateOperation is update operation.
	UpdateOperation = "Update"
	// DeleteOperation is delete operation.
	DeleteOperation = "Delete"
	// DeleteCollectionOperation is delete collection operation.
	DeleteCollectionOperation = "DeleteCollection"
)

var conn *amqp.Connection
var ch *amqp.Channel

// InitMQService will create the rabbitmq exchange.
func InitMQService() {
	var err error
	if conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/"); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Init MQ service failed, dail failed.")
		return
	}

	if ch, err = conn.Channel(); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Init MQ service failed, create channel failed.")
	}

	if err := ch.ExchangeDeclare(
		PromiseExchange, // name
		"topic",         // type
		true,            // duarable
		false,           // auto-deleted
		false,           // internal,
		false,           // no-wait,
		nil,             // args
	); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Init MQ service failed, create exchange failed.")
	}
	log.Info("MQ service initialized.")
}

// StopMQService stop event service, do all the clean jobs.
func StopMQService() {
	if ch != nil {
		ch.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

// Publish a message.
func Publish(category string, op string, body []byte) error {
	if ch == nil {
		log.Warn("Publish event failed, no channel, forgot to init event service?")
		return fmt.Errorf("no channel")
	}
	topic := strings.Join([]string{category, op}, ".")
	if err := ch.Publish(PromiseExchange, topic, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}); err != nil {
		log.WithFields(log.Fields{"category": category, "operation": op, "err": err}).Error("Publish event failed.")
		return err
	}
	log.WithFields(log.Fields{"category": category, "operation": op}).Info("Publish message.")
	return nil
}

type message struct {
	Type     string
	Category string
	Data     json.RawMessage
}

// PublishResourceMessage publishes a message according to the operation.
func PublishResourceMessage(op string, dto GetResponseInterface) error {
	return Publish(dto.GetCategory(), op, []byte(StructToString(message{
		Type:     op,
		Category: dto.GetCategory(),
		Data:     []byte(StructToString(dto)),
	})))
}

// PublishCreateMessage publishes a resource create message.
func PublishCreateMessage(dto GetResponseInterface) error {
	return Publish(dto.GetCategory(), CreateOperation, []byte(StructToString(message{
		Type:     CreateOperation,
		Category: dto.GetCategory(),
		Data:     []byte(StructToString(dto)),
	})))
}

// PublishUpdateMessage publishes a resource update message.
func PublishUpdateMessage(dto GetResponseInterface) error {
	return Publish(dto.GetCategory(), UpdateOperation, []byte(StructToString(message{
		Type:     UpdateOperation,
		Category: dto.GetCategory(),
		Data:     []byte(StructToString(dto)),
	})))
}

// PublishDeleteMessage publishes a resource delete message.
func PublishDeleteMessage(dto GetResponseInterface) error {
	return Publish(dto.GetCategory(), DeleteOperation, []byte(StructToString(message{
		Type:     DeleteOperation,
		Category: dto.GetCategory(),
		Data:     []byte(StructToString(dto)),
	})))
}

// PublishDeleteCollectionMessage publishes a resource collection delete message.
func PublishDeleteCollectionMessage(category string) error {
	return Publish(category, DeleteCollectionOperation, []byte(StructToString(message{
		Type:     DeleteCollectionOperation,
		Category: category,
		Data:     []byte(""),
	})))
}

// Subscribe the topics.
// The handler will process each of the delivery.
// You can call this method mutiple times to use other handlers to process other topics.
func Subscribe(topic []string, handler func(d *amqp.Delivery)) error {
	if ch == nil {
		log.Warn("Subscribe event failed, no channel, forgot to init event service?")
		return fmt.Errorf("no channel")
	}
	q, err := ch.QueueDeclare("", true, false, false, false, nil)
	if err != nil {
		log.WithFields(log.Fields{"topic": topic, "error": err}).Warn("Subscribe event failed, declare queue failed.")
		return err
	}
	log.WithFields(log.Fields{"Name": q.Name}).Info("Event queue created.")
	for _, v := range topic {
		if err := ch.QueueBind(q.Name, v, PromiseExchange, false, nil); err != nil {
			log.WithFields(log.Fields{"topic": v, "error": err}).Warn("Subscribe event failed, bind queue failed.")
			return err
		}
		log.WithFields(log.Fields{"topic": v}).Info("Event queue bind.")
	}
	delivery, err := ch.Consume(q.Name, "my-consume", false, false, false, false, nil)
	if err != nil {
		log.WithFields(log.Fields{"topic": topic, "error": err}).Warn("Subscribe event failed, consume failed.")
	}

	for each := range delivery {
		log.WithFields(log.Fields{"route": each.RoutingKey}).Info("Handle event.")
		handler(&each)
		each.Ack(false)
	}
	log.WithFields(log.Fields{"topic": topic}).Warn("Subscribe event exit.")
	return nil
}
