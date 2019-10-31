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
	"github.com/dollarkillerx/erguotou/token"
	"github.com/dollarkillerx/mongo/clog"
)

type user struct {
}

func UserController() *user {
	return &user{}
}

func (u *user) UserRegister(ctx *erguotou.Context) {
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

func (u *user) Login(ctx *erguotou.Context) {
	inputUser := &datamodels.User{}
	err := ctx.BindValue(inputUser)
	if err != nil {
		resp.RespMsg(ctx, defs.Err400)
		return
	}
	user := repositories.User{}

	data, er := user.Login(inputUser)
	if er == nil {
		// 用户验证成功返回token

		// 生成token
		jwt := token.NewJwt()
		jwt.Data = data
		jwt.User = data.Name
		jwt.Email = data.Email
		// erguotou的jwt默认单点登录
		s, err := token.Token.GeneraJwtToken(jwt)
		if err != nil {
			clog.PrintWa(err)
			resp.RespMsg(ctx, defs.Err500Token)
			return
		}

		resp.RespMsg(ctx, &defs.Resp{
			HttpCode: 200,
			Code:     200,
			Data: defs.H{
				"token": s,
			},
		})
		return
	} else {
		// 用户验证失败 返回错误信息
		resp.RespMsg(ctx, &defs.Resp{
			HttpCode: er.HttpCode,
			Code:     er.Code,
			Msg:      er.Msg,
		})
	}
}
