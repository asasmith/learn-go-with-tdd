package main

func Repeat(char string, amount int) string {
	var value string

	for i := 0; i < amount; i++ {
		value += char
	}

	return value
}
