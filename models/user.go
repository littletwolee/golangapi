package models

import "time"

type User struct {
//	ObjectId string
	Name         string
	Pwd          string
	CreateDate   time.Time
}

