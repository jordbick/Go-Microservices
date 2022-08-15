package account

// repo logic that interfaces with the database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/go-kit/kit/log"
)

// Normalise errors by creating a custom error
var RepoErr = errors.New("Unable to handle Repo Request")

// Use the repo struct to implement the Repository interface (in user.go) via function below
type repo struct {
	db     *sql.DB
	logger log.Logger
}

// Generate a new Repository interface that we implement on the repo struct
func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

// Need to create the 2 methods CreateUser and GetUser as the Repository interface implements both of these methods

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	// string of SQL to put user into our SQL DB
	sql := `
	INSERT INTO users (id, email, password)
	VALUES	($1, $2, $3)`

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	// Executes our SQL string with user ID, email and password to replace values
	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	// Query our DB and scan that result into our email variable
	err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)

	if err != nil {
		return "", RepoErr
	}

	return email, nil
}
