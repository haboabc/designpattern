package builder

/* 假如负载的对象是以下这样
 * type complex struct {
	 PartA()
	 PartB()
	 PartB()
 	}
 */

type BuildA interface {
	/* 定义需要用到零件的部分，
	 * 各部分的零件的实现由component 做
	 */

	/* like 
	 * PartA_component()
	 * PartB_component()
	 * PartC_component()
	 */

	PartA()

}
/* 由于go的原因，这里需要定义多个类似以下接口*/
/*
	type build_part_A interface {
		PartA_component()
	}
*/

type BuildB interface {
	PartB()
}

/* 选择零件 */
func SelectPart(param string) (BuildA ,BuildB) {
	switch (param) {
		case "A1":
			ret := &CompA1{}
			return ret, nil
		case "A2":
			ret := &CompA2{}
			return ret, nil
		case "B1":
			ret := &CompB1{}
			return nil, ret
		case "B2":
			ret := &CompB2{}
			return nil, ret
		default:
			return nil, nil
	}
}

/* 启用零件 */
func PartA(param BuildA) {
	param.PartA()
}

func PartB(param BuildB) {
	param.PartB()
}
