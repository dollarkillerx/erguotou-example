/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:14 2019-10-29
 */
package defs

type Resp struct {
	HttpCode int         `json:"-"`
	Code     int         `json:"code,omitempty"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data,omitempty"`
}

var (
	Err401      = &Resp{HttpCode: 401, Code: 401, Msg: "401 Unauthorized"}
	Err400      = &Resp{HttpCode: 400, Code: 400, Msg: "400 Parameter error"}
	Err500Token = &Resp{HttpCode: 500, Code: 500, Msg: "500 Token generation failed"}

	Success200 = &Resp{HttpCode: 200, Code: 200, Msg: "200 Success"}
)
