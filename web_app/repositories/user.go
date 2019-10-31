/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:24 2019-10-29
 */
package repositories

import (
	"context"
	"erguotou-example/web_app/datamodels"
	"erguotou-example/web_app/datasources/mongodb"
	"erguotou-example/web_app/defs"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/mongo/clog"
	"github.com/dollarkillerx/mongo/mongo-driver/bson"
	"github.com/dollarkillerx/mongo/mongo-driver/bson/primitive"
	"github.com/dollarkillerx/mongo/mongo-driver/mongo"
	"log"
)

type User struct {
}

func (u *User) Register(user *datamodels.User) *defs.Err {
	collection := mongodb.Mongo.Collection("user")
	user.Id = primitive.NewObjectID()

	getUser := datamodels.User{}
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&getUser)
	if err == mongo.ErrNoDocuments {
		// 如果 查询不存在
		user.Salt = easyutils.SuperRand()
		user.Password = easyutils.Sha256Encode(user.Salt + user.Password)
		// 入库
		_, err := collection.InsertOne(context.TODO(), user)
		if err == nil {
			return nil
		} else {
			clog.PrintWa(err)
			return &defs.Err{
				HttpCode: 500,
				Msg:      "服务器异常",
			}
		}
	} else {
		log.Println(err)
		log.Println(bson.ErrNilRegistry)
		return &defs.Err{
			HttpCode: 400,
			Msg:      "用户已经存在",
		}
	}
}

func (u *User) Login(user *datamodels.User) (*datamodels.User, *defs.Err) {
	collection := mongodb.Mongo.Collection("user")

	getUser := datamodels.User{}
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&getUser)
	if err != nil {
		return nil, &defs.Err{
			HttpCode: 400,
			Msg:      "用户名或者密码错误",
			Code:     200,
		}
	}
	pass := easyutils.Sha256Encode(getUser.Salt + user.Password)
	if pass == getUser.Password {
		// 如果密码正确
		return &getUser, nil
	} else {
		return nil, &defs.Err{
			HttpCode: 400,
			Msg:      "用户名或者密码错误",
			Code:     200,
		}
	}
}
