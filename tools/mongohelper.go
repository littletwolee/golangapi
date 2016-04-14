package tools

import(
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
	"mongoapi/models"
//	"fmt"
	"log"
)

var (
	session      *mgo.Session
	host         string = beego.AppConfig.String("mongohost")
	port         string = beego.AppConfig.String("mongoport")
	dbname       string = beego.AppConfig.String("mongodbname")
	url          string = host + ":" + port
	result       string
)

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

func GetOne(collectionname string, objectId string) interface{}{
	session := Session()
	db := session.DB(dbname)
	collection := db.C(collectionname)
	result := models.User{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(objectId)}).One(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}
func GetAll(collectionname string) interface{}{
	session := Session()
	db := session.DB(dbname)
	collection := db.C(collectionname)
	result := models.User{}
	err := collection.Find(nil).All(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}
