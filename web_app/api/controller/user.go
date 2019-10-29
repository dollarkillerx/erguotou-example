/**
 * @Author: DollarKiller
 * @Description: user controller
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:11 2019-10-29
 */
package controller

import (
	"erguotou-example/web_app/datamodels"
	"erguotou-example/web_app/defs"
	"erguotou-example/web_app/repositories"
	"erguotou-example/web_app/resp"
	"github.com/dollarkillerx/erguotou"
)

func UserRegister(ctx *erguotou.Context) {
	inputUser := &datamodels.User{}
	err := ctx.BindValue(inputUser)
	if err != nil {
		resp.RespMsg(ctx, defs.Err400)
		return
	}
	user := repositories.User{}
	errc := user.Register(inputUser)
	if errc != nil {
		// 如果插入错误
		resp.RespMsg(ctx, &defs.Resp{
			HttpCode: errc.HttpCode,
			Msg:      errc.Msg,
		})
		return
	}

	resp.RespMsg(ctx, defs.Success200)
}

func Login(ctx *erguotou.Context) {
	inputUser := &datamodels.User{}
	err := ctx.BindValue(inputUser)
	if err != nil {
		resp.RespMsg(ctx, defs.Err400)
		return
	}

}
