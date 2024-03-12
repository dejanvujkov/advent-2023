package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	sum := 0
	var line_number string

	for _, v := range strings.Split(input, "\n") {
		for _, character := range v {
			if unicode.IsDigit(character) {
				line_number = line_number + string(character)
			}
		}

		if len(line_number) == 0 {
			fmt.Println("No numbers")
		} else {
			twodigitstring := string(line_number[0]) + string(line_number[len(line_number)-1])
			parsed, err := strconv.Atoi(twodigitstring)
			if err != nil {
				panic(err.Error())
			}
			sum += parsed
			line_number = ""
		}
	}
	fmt.Printf("%v\n", sum)
}
