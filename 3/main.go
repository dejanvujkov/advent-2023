package main

/*

--- Day 3: Gear Ratios ---

You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/

import (
	"strconv"
	"strings"
)

func main() {
	input := `467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..`

	sum := 0
	potential_signs := []string{"*", "#", "+", "$"}
	delimiter := "."
	rows := strings.Split(input, "\n\t")

	for row_index, row_string := range rows {
		for _, row_char := range strings.Split(row_string, delimiter) {

			if row_char == "" {
				continue
			}

			// number next to sign
			if stringContainsAnySign(row_char, potential_signs...) && len(row_char) > 1 {
				row_char = removeCharFromString(row_char, potential_signs...)
				number, err := strconv.Atoi(row_char)
				if err != nil {
					continue
				}

				sum += number
				continue
			}

			// number isolated
			number_in_row, err := strconv.Atoi(row_char)
			if err != nil {
				continue
			}

			row_char_index := strings.Index(row_string, row_char)

			bottom_left_diagonal := false
			bottom_right_diagonal := false
			top_left_diagonal := false
			top_right_diagonal := false
			bottom := false
			top := false

			// first row, no above check, just below
			if row_index == 0 {
				if row_char_index-1 >= 0 {
					bottom_left_diagonal = stringContainsAnySign(string(rows[row_index+1][row_char_index-1]), potential_signs...)
				}
				if row_char_index+len(row_char)+1 <= len(rows[row_index+1]) {
					bottom_right_diagonal = stringContainsAnySign(string(rows[row_index+1][row_char_index+len(row_char)]), potential_signs...)
				}

				bottom = signInRow(rows[row_index+1], row_char_index, row_char_index+len(row_char), potential_signs...)

				// last row, no bottom check
			} else if row_index == len(rows)-1 {
				if row_char_index-1 >= 0 {
					top_left_diagonal = stringContainsAnySign(string(rows[row_index-1][row_char_index-1]), potential_signs...)
				}
				if row_char_index+len(row_char)+1 <= len(rows[row_index-1]) {
					top_right_diagonal = stringContainsAnySign(string(rows[row_index-1][row_char_index+len(row_char)]), potential_signs...)
				}

				top = signInRow(rows[row_index-1], row_char_index, row_char_index+len(row_char), potential_signs...)

				// somewhere middle, need check both sides
			} else {

				if row_char_index-1 >= 0 {
					bottom_left_diagonal = stringContainsAnySign(string(rows[row_index+1][row_char_index-1]), potential_signs...)
				}
				if row_char_index+len(row_char)+1 <= len(rows[row_index+1]) {
					bottom_right_diagonal = stringContainsAnySign(string(rows[row_index+1][row_char_index+len(row_char)]), potential_signs...)
				}

				bottom = signInRow(rows[row_index+1], row_char_index, row_char_index+len(row_char), potential_signs...)

				if row_char_index-1 >= 0 {
					top_left_diagonal = stringContainsAnySign(string(rows[row_index-1][row_char_index-1]), potential_signs...)
				}
				if row_char_index+len(row_char)+1 <= len(rows[row_index-1]) {
					top_right_diagonal = stringContainsAnySign(string(rows[row_index-1][row_char_index+len(row_char)]), potential_signs...)
				}
				top = signInRow(rows[row_index-1], row_char_index, row_char_index+len(row_char), potential_signs...)
			}

			if bottom_left_diagonal || bottom_right_diagonal || top_left_diagonal || top_right_diagonal || top || bottom {
				sum += number_in_row
			}

		}
	}

	println("Final sum: ", sum)
}

func signInRow(s string, start_index, end_index int, potential_signs ...string) bool {
	substring := strings.TrimSpace(s[start_index:end_index])

	for _, str := range substring {
		if stringContainsAnySign(string(str), potential_signs...) {
			return true
		}
	}
	return false
}

func stringContainsAnySign(row_char string, potential_signs ...string) bool {
	for _, sign := range potential_signs {
		if strings.Contains(row_char, sign) {
			return true
		}
	}
	return false
}

func removeCharFromString(row_char string, potential_signs ...string) string {
	for _, sign := range potential_signs {
		if strings.Contains(row_char, sign) {
			row_char = strings.Trim(row_char, sign)
		}
	}
	return row_char
}
