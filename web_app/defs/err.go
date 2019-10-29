/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:41 2019-10-29
 */
package defs

type Err struct {
	HttpCode int    `json:"http_code"`
	Msg      string `json:"msg"`
	Code     int    `json:"code"`
}
