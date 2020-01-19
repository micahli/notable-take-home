package db

import (
	"github.com/micahli/notable-take-home/db/mongodb"
	//"github.com/micahli/notable-take-home/model"
)

// Config represents the configuration of the database interface
type Config struct {
	MongoDB *mongodb.Config
}

// DB is the interface which must be implemented by all db drivers
type DB interface {
	CloseConnection() error

	// CreateUser(u *model.User) error
	// GetUser(id string) (*model.User, error)
	// GetUserByEmail(email string) (*model.User, error)
	// SaveUser(u *model.User) error
	// DeleteUser(id string) error

	// GetUnusedAward() (*model.AwardInfo, error)
	// UpdateAward(awardID string, userID string, cpnID string) error
}

// NewConnection creates a new database connection
func NewConnection(config *Config) (DB, error) {
	// Use MongoDB
	db, err := mongodb.NewConnection(config.MongoDB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
