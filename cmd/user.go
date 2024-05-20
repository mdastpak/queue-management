package cmd

import (
	"fmt"
	"queue-management/config"
	"queue-management/handlers"
	"queue-management/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage RabbitMQ users",
}

var createUserCmd = &cobra.Command{
	Use:   "create [username] [password] [tags]",
	Short: "Create a new RabbitMQ user",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		password := args[1]
		tags := args[2:]

		cfg, err := config.LoadConfig(viper.ConfigFileUsed())
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to load config")
		}

		userHandler, err := handlers.NewUserHandler(cfg)
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to create user handler")
		}

		err = userHandler.CreateUser(username, password, tags)
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to create user")
		}

		fmt.Printf("User %s created successfully\n", username)
	},
}

var deleteUserCmd = &cobra.Command{
	Use:   "delete [username]",
	Short: "Delete a RabbitMQ user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]

		cfg, err := config.LoadConfig(viper.ConfigFileUsed())
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to load config")
		}

		userHandler, err := handlers.NewUserHandler(cfg)
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to create user handler")
		}

		err = userHandler.DeleteUser(username)
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to delete user")
		}

		fmt.Printf("User %s deleted successfully\n", username)
	},
}

var updateUserPermissionsCmd = &cobra.Command{
	Use:   "update-permissions [username] [configure] [write] [read]",
	Short: "Update permissions for a RabbitMQ user",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		configure := args[1]
		write := args[2]
		read := args[3]

		cfg, err := config.LoadConfig(viper.ConfigFileUsed())
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to load config")
		}

		userHandler, err := handlers.NewUserHandler(cfg)
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to create user handler")
		}

		err = userHandler.UpdateUserPermissions(username, configure, write, read)
		if err != nil {
			log.Logger.WithError(err).Fatal("Failed to update user permissions")
		}

		fmt.Printf("Permissions for user %s updated successfully\n", username)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(createUserCmd)
	userCmd.AddCommand(deleteUserCmd)
	userCmd.AddCommand(updateUserPermissionsCmd)
}
