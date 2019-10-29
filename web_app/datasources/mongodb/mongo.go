/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:54 2019-10-29
 */
package mongodb

import (
	"erguotou-example/web_app/config"
	"github.com/dollarkillerx/mongo"
	"time"
)

var (
	Mongo *mongo.Database
)

func init() {
	db, e := mongo.Open(config.MyConfig.Mongo.Uri)
	if e != nil {
		panic(e)
	}
	db.SetConnMaxLifetime(config.MyConfig.Mongo.TimeOut * time.Millisecond)
	db.SetMaxOpenConn(config.MyConfig.Mongo.MaxOpen)
	Mongo = db.Database(config.MyConfig.Mongo.Dbname)
}


