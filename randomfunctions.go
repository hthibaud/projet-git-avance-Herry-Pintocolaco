package projet_git_avance_Herry_Pintocolaco

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}

	return Fibonacci(index-1) + Fibonacci(index-2)
}
func FirstRune(s string) rune {
	return []rune(s)[0]
}
