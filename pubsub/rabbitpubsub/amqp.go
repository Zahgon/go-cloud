package rabbitpubsub

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	wait = false

	mandatory = true

	immediate = false
)

type amqpConnection interface {
	Channel() (amqpChannel, error)
	Close() error
}

type amqpChannel interface {
	Publish(exchange, routingKey string, msg amqp.Publishing) error
	PublishWithContext(ctx context.Context, exchange, routingKey string, msg amqp.Publishing) error
	Consume(queue, consumer string) (<-chan amqp.Delivery, error)
	Ack(tag uint64) error
	Nack(tag uint64) error
	Cancel(consumer string) error
	Close() error
	NotifyPublish(chan amqp.Confirmation) chan amqp.Confirmation
	NotifyReturn(chan amqp.Return) chan amqp.Return
	NotifyClose(chan *amqp.Error) chan *amqp.Error
	ExchangeDeclare(string) error
	QueueDeclareAndBind(qname, ename string) error
	ExchangeDelete(string) error
	QueueDelete(qname string) error
	Qos(prefetchCount, prefetchSize int, global bool) error
}

type connection struct {
	conn *amqp.Connection
}

func (c *connection) Channel() (amqpChannel, error) {
	_ = "STUB: not implemented"
	return *new(amqpChannel), nil
}

func (c *connection) Close() error { _ = "STUB: not implemented"; return nil }

type channel struct {
	ch *amqp.Channel
}

func (ch *channel) Publish(exchange, routingKey string, msg amqp.Publishing) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *channel) PublishWithContext(ctx context.Context, exchange, routingKey string, msg amqp.Publishing) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *channel) Consume(queue, consumer string) (<-chan amqp.Delivery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ch *channel) Ack(tag uint64) error { _ = "STUB: not implemented"; return nil }

func (ch *channel) Nack(tag uint64) error { _ = "STUB: not implemented"; return nil }

func (ch *channel) Confirm() error { _ = "STUB: not implemented"; return nil }

func (ch *channel) Cancel(consumer string) error { _ = "STUB: not implemented"; return nil }

func (ch *channel) Close() error { _ = "STUB: not implemented"; return nil }

func (ch *channel) NotifyPublish(c chan amqp.Confirmation) chan amqp.Confirmation {
	_ = "STUB: not implemented"
	return nil
}

func (ch *channel) NotifyReturn(c chan amqp.Return) chan amqp.Return {
	_ = "STUB: not implemented"
	return nil
}

func (ch *channel) NotifyClose(c chan *amqp.Error) chan *amqp.Error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *channel) ExchangeDeclare(name string) error { _ = "STUB: not implemented"; return nil }

func (ch *channel) QueueDeclareAndBind(queueName, exchangeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *channel) ExchangeDelete(name string) error { _ = "STUB: not implemented"; return nil }

func (ch *channel) QueueDelete(qname string) error { _ = "STUB: not implemented"; return nil }

func (ch *channel) Qos(prefetchCount, prefetchSize int, global bool) error {
	_ = "STUB: not implemented"
	return nil
}
