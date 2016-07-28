/* shape 接口，shape提供draw()功能 */
package shape

import "fmt"

type Shape interface {
	Draw()
}

/* 工厂的功能Draw， 内置机器选择器 */
func Draw(shapetype string) {
	var mechine Shape
	switch (shapetype) {
	case "circle":
		mechine = Circle{}
	case "square":
		mechine = Square{}
	case "rectangle":
		mechine = Rectangle{}
	default:
		mechine = nil
	}

	if mechine != nil {
		mechine.Draw()
	} else {
		fmt.Println("unknown shape")
	}
}
