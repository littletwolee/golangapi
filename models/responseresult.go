package models

type ResponseResult struct {
//	ObjectId string
	Data         interface{}
	Msg          string
	Code         int
	Status       bool
}

