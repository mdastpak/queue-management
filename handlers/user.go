package handlers

import (
	"queue-management/config"
	"queue-management/log"

	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/sirupsen/logrus"
)

// UserHandler handles user operations in RabbitMQ.
type UserHandler struct {
	client *rabbithole.Client
}

// NewUserHandler creates a new UserHandler with the given configuration.
func NewUserHandler(cfg *config.Config) (*UserHandler, error) {
	client, err := rabbithole.NewClient(cfg.RabbitMQ.ManagementURL, cfg.RabbitMQ.Username, cfg.RabbitMQ.Password)
	if err != nil {
		log.Logger.WithError(err).Error("Failed to create RabbitMQ client")
		return nil, err
	}
	return &UserHandler{client: client}, nil
}

// CreateUser creates a new user with the specified username, password, and tags.
func (h *UserHandler) CreateUser(username, password string, tags []string) error {
	userInfo := rabbithole.UserSettings{
		Password: password,
		Tags:     rabbithole.UserTags(tags),
	}
	_, err := h.client.PutUser(username, userInfo)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("Failed to create user")
		return err
	}
	log.Logger.WithField("username", username).Info("User created successfully")
	return nil
}

// DeleteUser deletes the user with the specified username.
func (h *UserHandler) DeleteUser(username string) error {
	_, err := h.client.DeleteUser(username)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("Failed to delete user")
		return err
	}
	log.Logger.WithField("username", username).Info("User deleted successfully")
	return nil
}

// UpdateUserPermissions updates the permissions of the specified user.
func (h *UserHandler) UpdateUserPermissions(username, configure, write, read string) error {
	permissions := rabbithole.Permissions{
		Configure: configure,
		Write:     write,
		Read:      read,
	}
	_, err := h.client.UpdatePermissionsIn("/", username, permissions)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"username": username,
			"error":    err,
		}).Error("Failed to update user permissions")
		return err
	}
	log.Logger.WithField("username", username).Info("Permissions for user updated successfully")
	return nil
}
