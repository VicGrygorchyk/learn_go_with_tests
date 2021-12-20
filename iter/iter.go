package iter

func Repeat(input string, times int) string {
	repeat := ""
	for i := 0; i < times; i++ {
		repeat += input
	}
	return repeat
}