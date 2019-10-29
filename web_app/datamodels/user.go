/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:43 2019-10-29
 */
package datamodels

import "github.com/dollarkillerx/mongo/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `form:"name"`
	Email    string             `form:"email" json:"email"`
	Password string             `form:"password" json:"password"`
	Salt     string
	Avatar   string
	Data     int
}
