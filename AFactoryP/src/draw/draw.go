package main

import "factoryproducer"

func main() {
	/* 选择工厂 */
	shapeFactory := factoryproducer.GetFactory("shape")
	shapeFactory.StartService("circle")

	shapeFactory = factoryproducer.GetFactory("color")
	shapeFactory.StartService("red")

	/* non-exist factory */
	shapeFactory = factoryproducer.GetFactory("size")

	/* non-exist shape in a factory */
	shapeFactory = factoryproducer.GetFactory("shape")
	shapeFactory.StartService("line")

	/* start a sevice */
	factoryproducer.StartService("shape", "square")
}
