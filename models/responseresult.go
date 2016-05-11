package models

type ResponseResult struct {
	Data         interface{}
	Msg          string
	Code         int
	Status       bool
}

