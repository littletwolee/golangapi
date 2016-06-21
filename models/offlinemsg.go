package models

import (
	"time"
)

type OfflineMsg struct {
//	Id           string
	Fromuser string ;
	Touser string ;
	Msgdate time.Time ;
	Msgtype int ;
	Msgbody []byte ;
}
