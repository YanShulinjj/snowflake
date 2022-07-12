/* ----------------------------------
*  @author suyame 2022-07-12 11:09:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package snowflake

import "errors"

var (
	WorkerSpaceIdOverFlowErr = errors.New("worker id is overflowed, bigger than 31(5bit)! ")
	NodeIdOverFlowErr        = errors.New("node is id overflowed, bigger than 31(5bit)! ")
)
