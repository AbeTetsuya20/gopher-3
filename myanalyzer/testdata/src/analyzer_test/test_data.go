package analyzer_test

// 何も無い関数でもエラーにならないことをチェック
func f0() {}

// 引数が多く、複雑な関数
func f1(a, b, c, d, e, f int) int {
	return a + b + c + d + e + f
}

// 引数が少なく、簡単な関数
func f2(a int) int {
	return a
}

// 関数内で if 文を大量に使っていて複雑な関数
func f3(a int) int {
	if a > 0 {
		return a
	} else if a < 0 {
		return -a
	} else {
		return 0
	}
}

// 関数内で if 文が少ない、簡単な関数
func f4(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	f1(1, 2, 3, 4, 5, 6)
	f2(1)
	f3(1)
	f4(1)
}
