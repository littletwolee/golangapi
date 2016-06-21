package models

import "time"

type Userinfo struct {
	Id               interface{}  `json:"id" bson:"_id"`
	Name             string
	Qrcode           string
	Address          string
	Gender           string
	Area             string
	Signature        string
	Pic              string
	LastModifyDate   time.Time
	CreateDate       time.Time
}

