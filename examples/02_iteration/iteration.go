package iteration

func Repeat(character string, times int) string {
	var response string
	for i := 0; i < times; i++ {
		response += character
	}
	return response
}
