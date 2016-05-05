package tools

import(
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var (
	session      *mgo.Session
	host         string = beego.AppConfig.String("mongohost")
	port         string = beego.AppConfig.String("mongoport")
	url          string = host + ":" + port
	dbname string = beego.AppConfig.String("mongodbname")
)
type MongoConnectionPool struct{}

func Session() *mgo.Session {
    if session == nil {
        var err error
        session, err = mgo.Dial(url)
        if err != nil {
            panic(err) 
        }
    }
    return session.Clone()
}
