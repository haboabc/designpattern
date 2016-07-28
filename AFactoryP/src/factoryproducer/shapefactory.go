/* 代理工厂(本质是superfactory的mechine) */
package factoryproducer

import "shape"

type ShapeFactory struct {
	/* 是否需要继承? */
	//shape.Shape
}

/* 根据底层工厂shape实现的功能，代理工常可以根据这些功能生成相应的功能 */

/* 代理工厂负责跟底层的工厂交流 */
func (s ShapeFactory)StartService(param string) {
	/* 先只提供服务，不提供机器 */
	/* 调用底层工厂提供的 */
	shape.Draw(param)
}

/* 还可以实现获取机器, 本例不讨论 */
/*
func (s ShapeFactory)GetShape(param string) shape.Shape{
	shape.GetMechine(param)
}
*/
