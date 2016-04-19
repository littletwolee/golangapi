package tools

import(
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
	"log"
//	"mongoapi/models"
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

func GetOneById(collectionname string, objectId string) (result []byte, err error){
	session := Session()
	db := session.DB(dbname)
	collection := db.C(collectionname)
	data := bson.M{}
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(objectId)}).One(&data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err = bson.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func GetOneByFilter(collectionname string, filters map[string]string) (result []byte, err error){
	session := Session()
	db := session.DB(dbname)
	collection := db.C(collectionname)
	filter, err := bson.Marshal(filters)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	data := bson.M{}
	err = bson.Unmarshal(filter, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = collection.Find(data).One(&data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err = bson.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func GetAll(collectionname string) (result [][]byte, err error){
	session := Session()
	db := session.DB(dbname)
	collection := db.C(collectionname)
	data := []bson.M{}
	err = collection.Find(nil).All(&data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, item := range data {
		itembyte, err := bson.Marshal(item)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, itembyte)
	}
	return result, nil
}

//func GetIter
