package color

import "fmt"

type Blue struct {}

func (b *Blue) Fill() {
	fmt.Println("Blue is Filling...")
}
