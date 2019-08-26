package repository

import (
	"/model"

	"github.com/globalsign/mgo"
)

type UserRepository interface {
	GetAllUser() ([]model.User, error)
}
type UserRepositoryMongo struct {
	ConnectionDB *mgo.Session
}

const (
	DBName     = "CheckName"
	collection = "Teacher"
)

func (userMongo UserRepositoryMongo) GetAllUser() ([]model.User, error) {
	var users []model.Product
	err := userMongo.ConnectionDB.DB(DBName).C(collection).Find(nil).All(&users)
	return users, err
}
