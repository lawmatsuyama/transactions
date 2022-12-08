package domain

import "github.com/sirupsen/logrus"

type User struct {
	ID     string
	Name   string
	Active bool
}

func (user User) IsValid() error {
	l := logrus.WithField("userID", user.ID)
	if !user.Active {
		l.WithError(ErrDisabledUser).Error("User is disabled")
		return ErrDisabledUser
	}

	return nil
}
