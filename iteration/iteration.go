package iteration

func Repeat(character string, count int) string {
	if count == 0 {
		count = 1
	}
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
