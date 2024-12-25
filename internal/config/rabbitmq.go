package config

import (
	"book-service/pkg/logger"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func NewRabbitMQConn() (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		Viper().GetString("RABBITMQ_USER"),
		Viper().GetString("RABBITMQ_PASS"),
		Viper().GetString("RABBITMQ_HOST"),
		Viper().GetString("RABBITMQ_PORT"),
	)

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 5 * time.Second // Maximum time to retry
	maxRetries := 3                     // Number of retries (including the initial attempt)

	var conn *amqp.Connection
	var err error

	err = backoff.Retry(func() error {
		conn, err = amqp.Dial(connAddr)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to connect to RabbitMQ: %v. Connection information: %s", err, connAddr))
			return err
		}

		return nil
	}, backoff.WithMaxRetries(bo, uint64(maxRetries-1)))

	return conn, err
}
