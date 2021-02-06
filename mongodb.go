package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DBConfig struct {
	Host string
	Password string
	Username string
	Port string
}

func MongoConn(dbConfig *DBConfig) (client *mongo.Client) {
	credential := options.Credential{
		Username: dbConfig.Username,
		Password: dbConfig.Password,
	}

	connectionUri := fmt.Sprintf( "mongodb://%s:%s", dbConfig.Host, dbConfig.Port)
	clientOptions := options.Client().ApplyURI(connectionUri).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connection Made")
	return client
}

