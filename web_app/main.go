/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:58 2019-10-29
 */
package main

import (
	"erguotou-example/web_app/api/router"
	"erguotou-example/web_app/config"
	"github.com/dollarkillerx/erguotou"
)

func main() {
	app := erguotou.New()

	router.ApiRouter(app)

	app.Run(erguotou.SetHost(config.MyConfig.App.Host))
}
