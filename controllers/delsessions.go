package controllers

import (
	"fmt"
	"time"

	"../models"
	"gopkg.in/mgo.v2/bson"
)

// RemoveSessions watches for expired sessions and drops them
func RemoveSessions() {
	us := models.UserSession{}
	c := dbSession.Find(bson.M{}).Iter()
	for c.Next(&us) {
		if time.Now().Sub(us.CreatedAt).Seconds() >= 3600 {
			dbSession.Remove(struct{ UUID string }{UUID: string(us.UUID)})
			fmt.Println("dropping sessions...")
		}
	}
	if err := c.Close(); err != nil {
		fmt.Println("Iterations completed")
		return
	}
	fmt.Println("Closed successfully")
}
