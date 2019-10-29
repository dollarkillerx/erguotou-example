/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:14 2019-10-29
 */
package router

import (
	"erguotou-example/web_app/api/controller"
	"github.com/dollarkillerx/erguotou"
)

func ApiRouter(app *erguotou.Engine) {
	apiV1 := app.Group("/api/v1")
	{
		user := apiV1.Group("/user")
		{
			user.Post("/register", controller.UserRegister)
			user.Post("/login", controller.Login)
		}
	}
}
