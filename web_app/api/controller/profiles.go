/**
 * @Author: DollarKiller
 * @Description: profiles
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:01 2019-10-30
 */
package controller

import (
	"erguotou-example/web_app/datamodels"
	"erguotou-example/web_app/defs"
	"erguotou-example/web_app/repositories"
	"erguotou-example/web_app/resp"
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/mongo/clog"
	"log"
)

type profiles struct {
}

func Profiles() *profiles {
	return &profiles{}
}

func (p *profiles) Add(ctx *erguotou.Context) {
	inputData := datamodels.Profile{}
	err := ctx.BindValue(&inputData)
	if err != nil {
		resp.RespMsg(ctx, defs.Err400)
		return
	}
	// 存储
	profile := repositories.Profiles{}

	err = profile.Save(&inputData)
	if err != nil {
		// 入库失败
		log.Println(err)
		resp.RespMsg(ctx, defs.Err400)
		return
	}
	resp.RespMsg(ctx, defs.Success200)
}

func (p *profiles) All(ctx *erguotou.Context) {
	profile := repositories.Profiles{}
	all := profile.All()
	ctx.Json(200, all)
}

func (p *profiles) Edit(ctx *erguotou.Context) {
	id := ctx.PostVal("id")

	inputData := datamodels.Profile{}
	err := ctx.BindValue(&inputData)
	if err != nil {
		resp.RespMsg(ctx, defs.Err400)
		return
	}

	profile := repositories.Profiles{}
	err = profile.Edit(string(id), &inputData)
	if err != nil {
		// 入库失败
		clog.PrintWa(err)
		resp.RespMsg(ctx, defs.Err400)
		return
	}
	resp.RespMsg(ctx, defs.Success200)
}

func (p *profiles) Delete(ctx *erguotou.Context) {
	s, b := ctx.PathValueString("id")
	if !b {
		resp.RespMsg(ctx, defs.Err400)
		return
	}
	profile := repositories.Profiles{}
	e := profile.Delete(s)
	if e != nil {
		log.Println(e)
		resp.RespMsg(ctx, defs.Err400)
		return
	}

	resp.RespMsg(ctx, defs.Success200)
}
