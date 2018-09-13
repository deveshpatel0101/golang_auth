package controllers

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang_workspace/authentication/models"
	"gopkg.in/mgo.v2"
)

var dbUser, dbSession, dbReset *mgo.Collection
var s *mgo.Session

// Connect establishes connection to db
func Connect() {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}
	fmt.Println("Connected")
	dbUser = s.DB("go").C("first")
	dbSession = s.DB("go").C("session")
	dbReset = s.DB("go").C("reset")
}

// CreateUser is a function
func CreateUser(u models.UserDB) error {
	var userDb models.UserDB
	err := dbUser.Find(struct{ Email string }{Email: u.Email}).One(&userDb)
	if err == nil {
		return errors.New("user already exists")
	}
	if u.UserType == "local" {
		hshPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hshPass)
	}
	u.ID = bson.NewObjectId()
	err = dbUser.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

// ValidateUser is a function
func ValidateUser(u models.UserDB) (models.UserDB, error) {
	result := models.UserDB{}
	err := dbUser.Find(struct{ Email string }{Email: u.Email}).One(&result)
	if err != nil {
		return models.UserDB{}, errors.New("email or password is wrong")
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(u.Password))
	if err != nil {
		return models.UserDB{}, errors.New("email or password is wrong")
	}
	return result, nil
}

// GetUserByID is a function returns all information
// about user
func GetUserByID(u string) (models.UserDB, error) {
	ui := models.UserDB{}
	err := dbUser.FindId(bson.ObjectIdHex(u)).One(&ui)
	if err != nil {
		return models.UserDB{}, err
	}
	return ui, nil
}

// GetUserByEmail will return data retrieved from db
func GetUserByEmail(u string) (models.UserDB, error) {
	ui := models.UserDB{}
	err := dbUser.Find(struct{ Email string }{Email: u}).One(&ui)
	if err != nil {
		return models.UserDB{}, err
	}
	return ui, nil
}
