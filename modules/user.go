package modules

import(
	"golangapi/models"
	"golangapi/tools"
	"log"
	"gopkg.in/mgo.v2/bson"
//	"github.com/jmcvetta/neoism"
)

type User struct{}

const usercname = "user"
const offlinemsgcname = "offlinemsg"

func (u *User) GetOneUserById (ObjectId string) (user models.User, err error){
	result := models.User{}
	data, err := (&tools.MongoHelper{}).GetOneById(usercname, ObjectId)
	err = bson.Unmarshal(data, &result)
	return result, err
}

func (u *User) GetOneUserByFilter (filters map[string]string) (user models.User, err error){
	result := models.User{}
	data, err := (&tools.MongoHelper{}).GetOneByFilter(usercname, filters)
	err = bson.Unmarshal(data, &result)
	return result, err
}

func (u *User) GetAllUsers () (users []models.User,err error){
	result := []models.User{}
	data, err := (&tools.MongoHelper{}).GetAll(usercname)
	for _, item := range data {
		resultitem := models.User{}
		err = bson.Unmarshal(item, &resultitem)
		if err != nil {
			log.Panicln(err)
			return result, err
		}
		result = append(result, resultitem)
	} 
	return result, err
}

func (u *User) UpdateUser(objectId string, user map[string]interface{}) error {
	return (&tools.MongoHelper{}).UpdateById(usercname, objectId, user)
}

func (u *User) CreateUser(user models.User) (objectId string ,err error) {
	objectId, err = (&tools.MongoHelper{}).Create(usercname, user)
	if err != nil {
		return "", err
	}
	nodeid, err := u.CreateUserNode(objectId, user.Name)
	if err == nil {
		u.UpdateUser(objectId ,map[string]interface{}{"nodeid" : nodeid})
	}
	return objectId, err
}

func (u *User) DeleteUser (objectId string) error {
	return (&tools.MongoHelper{}).DeleteDoc(usercname, objectId)
}

func (u *User) CreateUserNode(objectId string, name string) (int, error) {
	user := map[string]interface{}{"objectId" : objectId, "name" : name}
	return (&tools.Neo4jHelper{}).CreateNode(user, usercname)
}

func (u *User) AddFriend(relationship models.Relationship) bool {
	// cypherquery := neoism.CypherQuery{
	// 	Statement: `MATCH (f:user), (t:user) 
        //                     WHERE f.objectId = {fromObjectId} AND t.objectId = {toObjectId}
        //                     CREATE (f)-[:friend]->(t)`,
	// 	Parameters: neoism.Props{"fromObjectId": relationship.FromObjectId,"toObjectId": relationship.ToObjectId},
	// }
	
	//objectids := []
	filters :=  bson.M{"_id":bson.M{"$in": []bson.ObjectId{
		bson.ObjectIdHex(relationship.FromObjectId) ,
		bson.ObjectIdHex(relationship.ToObjectId)}}}
	data, err := (&tools.MongoHelper{}).GetSomeByFilter(usercname, filters)
	if err != nil {
		return false
	}
	users := make(map[string]interface{})
	for _, item := range data {
		resultitem := models.User{}
		err = bson.Unmarshal(item, &resultitem)
		if err != nil {
			log.Panicln(err)
			return false
		}
		log.Println(resultitem.Id.(bson.ObjectId).String())
		users[resultitem.Id.(bson.ObjectId).String()] = resultitem
	}
	log.Println(users)
	_ ,err = (&tools.Neo4jHelper{}).CreateRelationship(
		users[relationship.FromObjectId].(models.User).NodeId,
		users[relationship.ToObjectId].(models.User).NodeId, "friend")
	if err != nil {
		return false
	}
	//offlinemsg := models.OfflineMsg{}
	//objectId, err = (&tools.MongoHelper{}).Create(offlinemsgcname, user)
	return true
}
