package account

import (
	"context"
	"log"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

// Business logic for our service
// Implement our service interface on a struct
type service struct {
	// Repository so we can interface with our DB and logger so we can see whats happening within our microservice
	repository Repository
	logger     log.Logger
}

// Function that will allow us to create this service
// Takes in a Repository and log.Logger and returns a Service interface (in the service file which itself implements 2 methods)
func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// Implement these methods on the service struct
// Here we are defining the CreateUser method that was called upon in the service file
func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	// Take the logger inside of our service and modify it so that we know we're executing this method
	logger := log.With(s.logger, "method", "CreateUser")

	// Generate an unique user ID
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	// Create a user struct
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	// s calls on service struct, calls on repository interface, calls on method Create user
	// Pass in context and user (user created above)
	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return "Success", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repository.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id)
	return email, nil

}
