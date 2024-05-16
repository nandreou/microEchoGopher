package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoUri                  = "mongodb://192.168.1.17:27017"
	brokerRequestsCollection  = "BrokerRequests"
	brokerResponsesCollection = "BrokerResponses"
	authRequestsCollection    = "AuthRequests"
	authResponsesCollection   = "AuthResponses"
)

type Mongo struct {
	MongoDB                   *mongo.Client
	MongoUri                  string
	BrokerRequestsCollection  string
	BrokerResponsesCollection string
	AuthRequestsCollection    string
	AuthResponsesCollection   string
}

type DB struct {
	Mongo *Mongo
}

var db = &DB{&Mongo{
	BrokerRequestsCollection:  brokerRequestsCollection,
	BrokerResponsesCollection: brokerResponsesCollection,
	AuthRequestsCollection:    authRequestsCollection,
	AuthResponsesCollection:   authResponsesCollection,
}}

func ConnectToDB() (*DB, error) {
	var err error

	db.Mongo.MongoDB, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri).SetAuth(options.Credential{
		Username: "admin",
		Password: "pass123",
	}))

	if err != nil {
		return nil, err
	}

	var pingBsonResult bson.M

	if err := db.Mongo.MongoDB.Database("logger").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&pingBsonResult); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) InsertToDB(log any, collection string) error {

	coll := db.Mongo.MongoDB.Database("logs").Collection(collection)
	_, err := coll.InsertOne(context.TODO(), log)

	if err != nil {
		return err
	}

	return nil
}
