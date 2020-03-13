package roby

func ExampleGroup_Run() {
	g := NewGroup(5)
	g.Run(3)

	// Output:
	// 10
}
