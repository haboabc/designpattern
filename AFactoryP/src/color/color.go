package color

import "fmt"

type Color interface {
	Fill()
}

func Fill(color string) {
	var mechine Color

	switch(color) {
		case "red":
			mechine = &Red{}
		case "green":
			mechine = &Green{}
		case "blue":
			mechine = &Blue{}
		default:
			mechine = nil
	}

	if mechine != nil {
		mechine.Fill()
	} else {
		fmt.Println("unknown color")
	}
}
/* 不能返回接口指针? 
func GetMechine(color string)  *Color{
	switch (color) {
		case "red":
			return &Red{}
		case "green":
			return &Green{}
		case "blue":
			return &Blue{}
		default:
			return nil
	}
}
*/
