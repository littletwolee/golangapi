package models

import (
	"time"
)

type User struct {
//	Id           string
	Name         string
	Pwd          string
	CreateDate   time.Time
}

// func (u *User)setObjectId(objectId ){
	
// }
