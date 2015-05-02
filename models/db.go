package models

import (
	"errors"
	"gopkg.in/mgo.v2"
)

func getDBConnection() {
	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err == nil {
		return &session, err
	} else {
		return nil, errors.New("Could not connect to database.")
	}
}
