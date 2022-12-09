package domain

import "github.com/sirupsen/logrus"

type User struct {
	ID     string `json:"_id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Active bool   `json:"active" bson:"active"`
}

func (user User) IsValid() error {
	l := logrus.WithField("userID", user.ID)
	if !user.Active {
		l.WithError(ErrDisabledUser).Error("User is disabled")
		return ErrDisabledUser
	}

	return nil
}
