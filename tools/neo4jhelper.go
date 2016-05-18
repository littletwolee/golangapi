package tools

import (
	"github.com/jmcvetta/neoism"
	"github.com/astaxie/beego"
	"log"
)
//your_user:your_password@neo4j.yourdomain.com/db/data/
var (
	neo4jhost = beego.AppConfig.String("redishost")
	neo4jport = beego.AppConfig.String("neo4jport")
	neo4juser = beego.AppConfig.String("neo4juser")
	neo4jpwd = beego.AppConfig.String("neo4jpwd")
)

type Neo4jHelper struct{}
//*neoism.Database
func newNeo4jDB() neoism.Database {
	if neo4jport != "" {
		neo4jport = ":"+ neo4jport
	}
	db, err := neoism.Connect("http://" + neo4juser + ":" + neo4jpwd + "@" + neo4jhost + neo4jport + "/db/data")
	if err != nil {
		log.Println(err.Error())
	}
	return *db
}


func (r *Neo4jHelper) CreateNode(node map[string]interface{}) error {
	db := newNeo4jDB()
	if _, err := db.CreateNode(node); err != nil {
		return err
	}
	return nil
}
