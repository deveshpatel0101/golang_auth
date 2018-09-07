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
	Picture  string        `json:"picture" bson:"picture"`
	UserType string        `json:"userType" bson:"userType"`
	GoogleID string        `json:"googleID" bson:"googleID"`
}

// UserAlerts to alert user for wrong field value inputs
type UserAlerts struct {
	SuccessMessage string
	ErrorMessage   string
	LoggedIn       bool
}

// UserSession to store session in db
type UserSession struct {
	UUID      string        `json:"uuid" bson:"uuid"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	UserID    bson.ObjectId `json:"userId" bson:"userId"`
}

// GoogleUser to store google user info
type GoogleUser struct {
	Fname    string `json:"given_name" bson:"given_name"`
	Lname    string `json:"family_name" bson:"family_name"`
	Email    string `json:"email" bson:"email"`
	Picture  string `json:"picture" bson:"picture"`
	GoogleID string `json:"id" bson:"id"`
}
