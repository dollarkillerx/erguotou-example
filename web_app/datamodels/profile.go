/**
 * @Author: DollarKiller
 * @Description: profile
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 09:51 2019-10-30
 */
package datamodels

import "github.com/dollarkillerx/mongo/mongo-driver/bson/primitive"

type Profile struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Describe string             `json:"describe" form:"describe" validate:"required"`
	Type     string             `json:"type" form:"type" validate:"required"`
	Incode   string             `json:"incode" form:"incode" validate:"required"`
	Expend   string             `json:"expend" form:"expend" validate:"required"`
	Cash     string             `json:"cash" form:"cash" validate:"required"`
	Remark   string             `json:"remark" form:"remark" validate:"required"`
	Data     int                `json:"data"`
}
