package controllers

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/golang_workspace/authentication/flash"
	"github.com/golang_workspace/authentication/models"
	"github.com/satori/go.uuid"
)

// CreateSession is a function
func CreateSession(u models.UserDB) (models.UserSession, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return models.UserSession{}, err
	}
	us := models.UserSession{
		UUID:      uid.String(),
		CreatedAt: time.Now(),
		UserID:    u.ID,
	}
	err = dbSession.Insert(us)
	fmt.Println(err)
	if err != nil {
		return models.UserSession{}, err
	}
	return us, nil
}

// GetSession function gives user session from cookie id value
func GetSession(s string) (bson.ObjectId, error) {
	ds, err := flash.Decode(s)
	if err != nil {
		return "", err
	}
	var us models.UserSession
	err = dbSession.Find(struct{ UUID string }{UUID: string(ds)}).One(&us)
	if err != nil {
		return "", err
	}
	if time.Now().Sub(us.CreatedAt).Seconds() >= 3600 {
		err = dbSession.Remove(struct{ UUID string }{UUID: string(ds)})
		if err != nil {
			return "", errors.New("error removing session from db")
		}
		return "", errors.New("session timed out")
	}
	return us.UserID, nil
}
