/* 实现提供Draw()功能的circle机器 */
package shape

import "fmt"

type Circle struct {
}

/* 如果是 (c *Circle) 则认为Circle没有实现Draw() */
func (c Circle)Draw() {
	fmt.Println("Circle Mechine is working")
}
