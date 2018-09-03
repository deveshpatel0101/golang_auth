package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// UserDB to insert data in db
type UserDB struct {
	Fname    string        `json:"fname" bson:"fname"`
	Lname    string        `json:"lname" bson:"lname"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
	ID       bson.ObjectId `json:"_id" bson:"_id"`
}

// UserAlerts to alert user for wrong field value inputs
type UserAlerts struct {
	Fname          string
	Lname          string
	Email          string
	Password       []byte
	SuccessMessage string
	ErrorMessage   string
}

// UserSession to store session in db
type UserSession struct {
	UUID      string        `json:"uuid" bson:"uuid"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UserID    bson.ObjectId `json:"userId" bson:"userId"`
}
