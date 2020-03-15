package roby

import (
	"fmt"
	"math"
)

func ExampleGroup_Run() {
	g := NewGroup(200)
	g.Run(10)

	// Output:
	// 10
}

func ExampleNewGroup() {
	fmt.Println(int(math.Ceil(3 / 2)))
	fmt.Println([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}[10-2:])

	// Output:
	// 10
}
