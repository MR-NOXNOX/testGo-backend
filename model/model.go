package model

import (
	
	"github.com/globalsign/mgo/bson"
)
type UserInfo struct {
	User []User  `json:"User"`
}

type User struct{
	UserID    bson.ObjectId `json:"user_id" bson:"_id,omitempty"`
	UserName 	string      `json:"user_name" bson:"user_name"`
	TId        	string      `json:"t_id" bson:"t_id"`
	TEmail     	string		`json:"t_email" bson:"t_email"`
	TWorkPlace 	string  	`json:"t_workplace" bson:"t_workplace"`
	TPassword  	string 		`json:"t_password" bson:"t_password"`
	

}