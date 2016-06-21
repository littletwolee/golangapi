package models

import (
	"time"
)

type User struct {
	Id           interface{}  `json:"id" bson:"_id"`
	Name         string
	Pwd          string
	NodeId       int
	CreateDate   time.Time
}

// func (u *User)setObjectId(objectId ){
	
// }
