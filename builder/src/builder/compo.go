package builder

import "fmt"

type CompA1 struct {
}

type CompA2 struct {
}

type CompB1 struct {}

type CompB2 struct {}

func (c *CompA1) PartA() {
	fmt.Println("Building A part in CompA1...")
}

func (c *CompA2) PartA() {
	fmt.Println("Building A part in CompA2...")
}

func (c *CompB1) PartB() {
	fmt.Println("Building B part in CompB1...")
}

func (c *CompB2) PartB() {
	fmt.Println("Building B part in CompB2...")
}
