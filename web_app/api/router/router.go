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
			// 用户注册
			user.Post("/register", controller.UserController().UserRegister)
			// 用户登录
			user.Post("/login", controller.UserController().Login)
		}

		profiles := apiV1.Group("/profiles")
		{
			// 添加
			profiles.Post("/add", controller.Profiles().Add)
			// 获取所有
			profiles.Get("/", controller.Profiles().All)
			// 修改
			profiles.Put("/edit", controller.Profiles().Edit)
			// 删除
			profiles.Delete("/delete/:id", controller.Profiles().Delete)
		}
	}
}
