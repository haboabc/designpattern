package director

import "builder"

func Build() {

	/* 获取零件 */
	a,_ := builder.SelectPart("A1")
	_,b := builder.SelectPart("B2")

	/* 指定顺序 */
	builder.PartA(a)
	builder.PartB(b)
}
