package color

import "fmt"

type Red struct {
}

func (r *Red) Fill() {
	fmt.Println("Red Filling...")
}
