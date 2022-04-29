package iteration

// Repeat takes a character and repeats it 5 times.
func Repeat(character string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return
}