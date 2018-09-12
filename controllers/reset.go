package controllers

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang_workspace/authentication/models"
	"gopkg.in/mgo.v2/bson"
)

// CreateReset function will create reset link and will store data in db
func CreateReset(email string) (models.UserReset, error) {
	id := bson.NewObjectId()
	ur := models.UserReset{
		Email:     email,
		ID:        id,
		CreatedAt: time.Now(),
	}

	err := dbReset.Insert(ur)
	if err != nil {
		return models.UserReset{}, err
	}
	return ur, nil
}

// GetReset function will fetch reset data from db using
// id and will check if token will still be valid or not
// returns error if token invalid or either an error occured
// during fetch else returns user info
func GetReset(id string) (models.UserReset, error) {
	ur := models.UserReset{}
	err := dbReset.FindId(bson.ObjectIdHex(id)).One(&ur)
	if err != nil {
		return models.UserReset{}, err
	}

	if time.Now().Sub(ur.CreatedAt).Seconds() > 3600 {
		err = dbReset.Remove(struct{ ID bson.ObjectId }{ID: bson.ObjectIdHex(id)})
		if err != nil {
			return models.UserReset{}, err
		}

		return models.UserReset{}, errors.New("token timed out")
	}
	return ur, nil
}

// UpdatePassword function will update password
// checks if new password matches old password
// else it will save new password in db
func UpdatePassword(email, p string) error {
	ui := models.UserDB{}

	// Generate hahs password
	hsh, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Find user info from user db
	err = dbUser.Find(struct{ Email string }{Email: email}).One(&ui)
	if err != nil {
		return err
	}

	// Compare new password and old password
	if ui.Password != "" {
		err = bcrypt.CompareHashAndPassword([]byte(ui.Password), []byte(p))
		if err == nil {
			return errors.New("You new password should not match old password")
		}
	}

	// Remove reset info from reset db
	dbReset.Remove(struct{ Email string }{Email: email})

	p = string(hsh)

	// Update password in user db
	err = dbUser.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"password": p}})
	if err != nil {
		return err
	}

	return nil
}
