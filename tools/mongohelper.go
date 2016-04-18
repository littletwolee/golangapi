package tools

import(
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
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

func GetOne(collectionname string, objectId string) []byte{
	session := Session()
	db := session.DB(dbname)
	collection := db.C(collectionname)
	data := bson.M{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(objectId)}).One(&data)
	result, err := bson.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return result
}
