/**
 * @Author: DollarKiller
 * @Description: 通用返回
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:14 2019-10-29
 */
package resp

import (
	"erguotou-example/web_app/defs"
	"github.com/dollarkillerx/erguotou"
)

func RespMsg(ctx *erguotou.Context, msg *defs.Resp) {
	ctx.Json(msg.HttpCode, msg)
}
