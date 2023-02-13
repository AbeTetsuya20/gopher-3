package a

// 最初からある関数
func f() {
	// The pattern can be written in regular expression.
	var gopher int
	print(gopher)
}

// 引数がない関数
func f0() {}

// 引数が多く、複雑な関数
func f1(a, b, c, d, e, f int) int { // want "too many arguments: func name: f1"
	return a + b + c + d + e + f
}

// 引数が少なく、簡単な関数
func f2(a int) int {
	return a
}
