/**
 * @Author: DollarKiller
 * @Description: profiles
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:17 2019-10-30
 */
package repositories

import (
	"context"
	"erguotou-example/web_app/datamodels"
	"erguotou-example/web_app/datasources/mongodb"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/mongo/clog"
	"github.com/dollarkillerx/mongo/mongo-driver/bson"
	"github.com/dollarkillerx/mongo/mongo-driver/bson/primitive"
	"log"
)

type Profiles struct {
}

// 保存
func (p *Profiles) Save(data *datamodels.Profile) error {
	collection := mongodb.Mongo.Collection("profiles")
	data.Id = primitive.NewObjectID()
	data.Data = easyutils.TimeGetNowTime()
	_, e := collection.InsertOne(context.TODO(), data)
	return e
}

// 查询所有
func (p *Profiles) All() []*datamodels.Profile {
	collection := mongodb.Mongo.Collection("profiles")
	cursor, e := collection.Find(context.TODO(), bson.D{})
	if e != nil {
		return nil
	}
	e = cursor.Err()
	if e != nil {
		return nil
	}

	data := make([]*datamodels.Profile, 0)

	e = cursor.All(context.TODO(), &data)
	if e != nil {
		return nil
	}
	defer cursor.Close(context.TODO())
	return data
}

// 修改
func (p *Profiles) Edit(id string, data *datamodels.Profile) error {
	update := bson.M{"$set": data}
	collection := mongodb.Mongo.Collection("profiles")
	_, e := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	return e
}

// 删除
func (p *Profiles) Delete(id string) error {
	log.Println(id)
	collection := mongodb.Mongo.Collection("profiles")
	// 把 str 转换成 objectId
	ids, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		clog.PrintWa(e)
		return e
	}
	_, e = collection.DeleteOne(context.TODO(), bson.M{"_id": ids})
	return e
}
