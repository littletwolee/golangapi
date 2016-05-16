package tools

import(
//	"github.com/astaxie/beego"
//	"gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
//	"log"
	"errors"
	"time"
	"golangapi/models"
)

const gridfscname = "fs"

type MongoGridFSHelper struct{}

func (m *MongoGridFSHelper) GetFileById (collectionname string, objectId string) (rangemode interface{},err error){
	session := Session()
	db := session.DB(dbname)
	if ! bson.IsObjectIdHex(objectId) {
		err = errors.New("It is not a objectId")
		return nil, err
	}
	file, err := db.GridFS(gridfscname).OpenId(bson.ObjectIdHex(objectId))
	if err != nil {
		return nil, err
	}
	data := make([]byte, file.Size())
	_ ,err = file.Read(data)
	if err != nil {
		return nil, err
	}
	rangemode = &models.Rangemodel{
		file.Name(),
		data,
		file.ContentType(),
		0,file.Size()}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	return rangemode, nil
}


func (m *MongoGridFSHelper) UploadFile (filemode models.Filemodel) (objectId string, err error) {
	session := Session()
	db := session.DB(dbname)
	newid := bson.NewObjectId()
	filename := newid.Hex() + "." + filemode.Filetype
	file, err := db.GridFS(gridfscname).Create(filename)
	if err != nil {
		return "", err
	}
	file.SetId(newid)
	file.SetName(filename)
	file.SetUploadDate(time.Now())
	file.SetContentType(filemode.Contenttype)
	_, err = file.Write(filemode.Filedata)
	if err != nil {
		return "", err
	}
	err = file.Close()
	if err != nil {
		return "", err
	}
	return newid.Hex(), nil
}


func (m *MongoGridFSHelper) DeleteFileById (objectId string) error {	
	session := Session()
	db := session.DB(dbname)
	err := db.GridFS(gridfscname).RemoveId(bson.ObjectIdHex(objectId))
	if err != nil {
		return err
	}
	return nil
}
