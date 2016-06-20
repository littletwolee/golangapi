package tools

import (
	"github.com/jmcvetta/neoism"
	"github.com/astaxie/beego"
	"log"
	"strings"
)
//your_user:your_password@neo4j.yourdomain.com/db/data/
var (
	neo4jhost = beego.AppConfig.String("neo4jhost")
	neo4jport = beego.AppConfig.String("neo4jport")
	neo4juser = beego.AppConfig.String("neo4juser")
	neo4jpwd = beego.AppConfig.String("neo4jpwd")
)

type Neo4jHelper struct{}
//*neoism.Database
func newNeo4jDB() neoism.Database {
	if neo4jport != "" && !strings.Contains(neo4jport,":") {
		neo4jport = ":"+ neo4jport
	}
	db, err := neoism.Connect("http://" + neo4juser + ":" + neo4jpwd + "@" + neo4jhost + neo4jport + "/db/data")
	if err != nil {
		log.Println(err.Error())
	}
	return *db
}

func (r *Neo4jHelper) CommitNodeByQuery(query neoism.CypherQuery) (error) {
	db := newNeo4jDB()
	err := db.Cypher(&query)
        if err != nil{
                log.Println(err.Error())
                return err
        }
        return nil
}

func (r *Neo4jHelper) GetNode(nodeid int) (*neoism.Node, error) {
	db := newNeo4jDB()
	node, err := db.Node(nodeid)
	if err != nil {
		return nil, err
	} 
	return node, nil
}

func (r *Neo4jHelper) CreateNode(properties map[string]interface{}, label string) (int, error) {
	db := newNeo4jDB()
	node, err := db.CreateNode(properties)
	if err != nil {
		return -1, err
	}
	err = node.AddLabel(label)
	if err != nil {
		return -1, err
	}
	return node.Id(), nil
}

func (r *Neo4jHelper) DeleteNode(nodeid int) error {
	db := newNeo4jDB()
	node, err := db.Node(nodeid)
	if err != nil {
		return err
	}
	err = node.Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *Neo4jHelper) UpdateNode(nodeid int, properties map[string]interface{}) error {
	db := newNeo4jDB()
	node, err := db.Node(nodeid)
	if err != nil {
		return err
	}
	err = node.SetProperties(properties)
	if err != nil {
		return err
	}
	return nil
}

func (r *Neo4jHelper) CreateRelationship(startid int, endid int, relationship string) (int, error) {
	db := newNeo4jDB()
	nstart, _ := db.Node(startid)
	rid, err := nstart.Relate(relationship, endid, nil)
	if err != nil {
		return -1, err
	}
	return rid.Id(), nil
}

func (r *Neo4jHelper) DeleteRelationship(relationshipid int) error {
	db := newNeo4jDB()
	relationship, err := db.Relationship(relationshipid)
	if err != nil {
		return err
	}
	err = relationship.Delete()
	if err != nil {
		return err
	}
	return nil
}

// func (r *Neo4jHelper) DeleteRelationship(startid int, endid int, relationship string) error {
// 	db := newNeo4jDB()
// 	nstart, _ := db.Node(startid)
// 	_, err := nstart.Relate(relationship, endid, nil)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
