package domain

import "github.com/sirupsen/logrus"

// User represents user
type User struct {
	ID     string `json:"_id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Active bool   `json:"active" bson:"active"`
}

// IsValid check if user is valid
func (user User) IsValid() error {
	l := logrus.WithField("userID", user.ID)
	if !user.Active {
		l.WithError(ErrDisabledUser).Error("User is disabled")
		return ErrDisabledUser
	}

	return nil
}
