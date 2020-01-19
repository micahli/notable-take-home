package mongodb

import (
	"context"

	// "github.com/mongodb/mongo-go-driver/bson"
	// "github.com/mongodb/mongo-go-driver/mongo"
	// "github.com/mongodb/mongo-go-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config represents MongoDB configuration
type Config struct {
	ConnectionURI string `yaml:"connection_uri"`
	DatabaseName  string `yaml:"database_name"`
}

// DB represents the structure of the database
type DB struct {
	config      *Config
	client      *mongo.Client
	collections *Collections
}

// Collections represents all needed db collections
type Collections struct {
	users *mongo.Collection
}

// NewConnection creates a new database connection
func NewConnection(config *Config) (*DB, error) {
	//client, err := mongo.Connect(context.Background(), config.ConnectionURI)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.ConnectionURI))
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	userIndexOptions := options.Index()
	userIndexOptions.SetUnique(true)

	users := client.Database(config.DatabaseName).Collection("users")
	users.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: userIndexOptions,
	})

	// // TODO the index should modify later
	// trackInfoIndexOptions := options.Index()
	// trackInfoIndexOptions.SetUnique(true)
	// trackInfo := client.Database(config.DatabaseName).Collection("cpntrackinfo")
	// trackInfo.Indexes().CreateOne(context.Background(), mongo.IndexModel{
	// 	Keys: bson.M{
	// 		"idfa": 1,
	// 	},
	// 	Options: trackInfoIndexOptions,
	// })

	// awardInfoOptions := options.Index()
	// awardInfoOptions.SetUnique(true)

	// awardInfo := client.Database(config.DatabaseName).Collection("awardinfo")
	// awardInfo.Indexes().CreateOne(context.Background(), mongo.IndexModel{
	// 	Keys: bson.M{
	// 		"awardid": 1,
	// 	},
	// 	Options: awardInfoOptions,
	// })

	collections := &Collections{
		users: users,
		//cpntrackinfo: trackInfo,
		//awardinfo:    awardInfo,
	}

	return &DB{
		config:      config,
		client:      client,
		collections: collections,
	}, nil
}

// CloseConnection closes the database connection
func (db *DB) CloseConnection() error {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		return err
	}

	return nil
}
