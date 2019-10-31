/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:43 2019-10-29
 */
package datamodels

import "github.com/dollarkillerx/mongo/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id" form:"id"`
	Name     string             `form:"name"`
	Email    string             `form:"email" json:"email" validate:"required"`
	Password string             `form:"password" json:"password" validate:"required"`
	Salt     string
	Avatar   string
	Data     int
}
