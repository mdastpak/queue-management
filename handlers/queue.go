package handlers

import (
	"queue-management/config"
	"queue-management/log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

// QueueHandler handles queue operations in RabbitMQ.
type QueueHandler struct {
	conn *amqp.Connection
}

// NewQueueHandler creates a new QueueHandler with the given configuration.
func NewQueueHandler(cfg *config.Config) (*QueueHandler, error) {
	conn, err := amqp.Dial(cfg.RabbitMQ.URL)
	if err != nil {
		log.Logger.WithError(err).Error("Failed to connect to RabbitMQ")
		return nil, err
	}
	return &QueueHandler{conn: conn}, nil
}

// CreateQueue creates a new queue with the specified name.
func (h *QueueHandler) CreateQueue(queueName string) error {
	ch, err := h.conn.Channel()
	if err != nil {
		log.Logger.WithError(err).Error("Failed to open a channel")
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"queueName": queueName,
			"error":     err,
		}).Error("Failed to declare queue")
		return err
	}
	log.Logger.WithField("queueName", queueName).Info("Queue created successfully")
	return nil
}
