package iteration

func Repeat(a string, b int) string {
	var tmp string
	for i := 0; i < b; i++ {
		tmp += a
	}
	return tmp
}
