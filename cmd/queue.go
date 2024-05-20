package cmd

import (
	"fmt"
	"queue-management/config"
	"queue-management/handlers"
	"queue-management/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// queueCmd represents the queue command
var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Manage RabbitMQ queues",
}

var createQueueCmd = &cobra.Command{
	Use:   "create [queueName]",
	Short: "Create a new RabbitMQ queue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		queueName := args[0]

		cfg, err := config.LoadConfig(viper.ConfigFileUsed())
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to load config")
		}

		queueHandler, err := handlers.NewQueueHandler(cfg)
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to create queue handler")
		}

		err = queueHandler.CreateQueue(queueName)
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to create queue")
		}

		fmt.Printf("Queue %s created successfully\n", queueName)
	},
}

func init() {
	rootCmd.AddCommand(queueCmd)
	queueCmd.AddCommand(createQueueCmd)
}
