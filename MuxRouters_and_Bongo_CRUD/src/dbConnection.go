package main

import "github.com/go-bongo/bongo"

func dbConnection() (connection *bongo.Connection) {

	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "pelada",
	}

	connection, err := bongo.Connect(config)

	errHandler(err)

	return connection

}

