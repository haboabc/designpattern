/* 代理工厂(superfactory 的 mechine) */
package factoryproducer

import "color"

type ColorFactory struct {
	/* 这里用匿名类实现继承 */
	/* 需要吗*/
	//color.Color
}

/* 代理工厂和底层工常交流 */
func (c ColorFactory)StartService(param string) {
	color.Fill(param)
}
