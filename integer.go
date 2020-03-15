package roby

// Exp 指数运算
func Exp(num, n int) int {
	result := 1
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= num
		}
		num *= num
	}
	return result
}
