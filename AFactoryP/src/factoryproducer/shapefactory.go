/* 代理工厂(本质是superfactory的mechine) */
package factoryproducer

import "shape"

type ShapeFactory struct {
	/* 是否需要继承? */
	//shape.Shape
}

/* 代理工厂负责跟底层的工厂交流 */
func (s ShapeFactory)StartService(param string) {
	/* 先只提供服务，不提供机器 */
	/* 调用底层工厂提供的 */
	shape.Draw(param)
}
