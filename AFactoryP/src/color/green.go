package color

import "fmt"

type Green struct {}

func (g *Green) Fill() {
	fmt.Println("Green id Filling...")
}
