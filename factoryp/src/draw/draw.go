package main

import "shape"

func main() {
	/* 使用工厂的功能 */
	shape.Draw("circle")
	shape.Draw("square")
	shape.Draw("rectangle")
	shape.Draw("line")

	/* 使用工厂的机器 */
	mechine := shape.GetMechine("circle")
	mechine.Draw()
}
