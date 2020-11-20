
// 1. stack

package main
 
import (
	"fmt"
)
 
type item struct {
	value interface{} 
	next  *item
}
 
type Stack struct {
	top  *item
	size int
}
 
func (stack *Stack) Len() int {
	return stack.size
}
 
func (stack *Stack) Push(value interface{}) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}
 
func (stack *Stack) Pop() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
 
	return nil
}
 
func main() {
	stack := new(Stack)

	stack.Push(1)
	stack.Push("prasanthi")
	stack.Push(4.0)

	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
}

$go run main.go
4
prasanthi
1


