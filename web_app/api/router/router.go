/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:14 2019-10-29
 */
package router

import "github.com/dollarkillerx/erguotou"

func ApiRouter(app *erguotou.Engine) {
	app.Get("/", func(ctx *erguotou.Context) {
		ctx.String(200,"Hello World")
	})
}

