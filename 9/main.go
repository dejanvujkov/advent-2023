package main

/*
--- Day 9: Mirage Maintenance ---

You ride the camel through the sandstorm and stop where the ghost's maps told you to stop. The sandstorm subsequently subsides, somehow seeing you standing at an oasis!

The camel goes to get some water and you stretch your neck. As you look up, you discover what must be yet another giant floating island, this one made of metal! That must be where the parts to fix the sand machines come from.

There's even a hang glider partially buried in the sand here; once the sun rises and heats up the sand, you might be able to use the glider and the hot air to get all the way up to the metal island!

While you wait for the sun to rise, you admire the oasis hidden here in the middle of Desert Island. It must have a delicate ecosystem; you might as well take some ecological readings while you wait. Maybe you can report any environmental instabilities you find to someone so the oasis can be around for the next sandstorm-worn traveler.

You pull out your handy Oasis And Sand Instability Sensor and analyze your surroundings. The OASIS produces a report of many values and how they are changing over time (your puzzle input). Each line in the report contains the history of a single value. For example:

0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45

To best protect the oasis, your environmental report should include a prediction of the next value in each history. To do this, start by making a new sequence from the difference at each step of your history. If that sequence is not all zeroes, repeat this process, using the sequence you just generated as the input sequence. Once all of the values in your latest sequence are zeroes, you can extrapolate what the next value of the original history should be.

In the above dataset, the first history is 0 3 6 9 12 15. Because the values increase by 3 each step, the first sequence of differences that you generate will be 3 3 3 3 3. Note that this sequence has one fewer value than the input sequence because at each step it considers two numbers from the input. Since these values aren't all zero, repeat the process: the values differ by 0 at each step, so the next sequence is 0 0 0 0. This means you have enough information to extrapolate the history! Visually, these sequences can be arranged like this:

0   3   6   9  12  15
  3   3   3   3   3
    0   0   0   0

To extrapolate, start by adding a new zero to the end of your list of zeroes; because the zeroes represent differences between the two values above them, this also means there is now a placeholder in every sequence above it:

0   3   6   9  12  15   B
  3   3   3   3   3   A
    0   0   0   0   0

You can then start filling in placeholders from the bottom up. A needs to be the result of increasing 3 (the value to its left) by 0 (the value below it); this means A must be 3:

0   3   6   9  12  15   B
  3   3   3   3   3   3
    0   0   0   0   0

Finally, you can fill in B, which needs to be the result of increasing 15 (the value to its left) by 3 (the value below it), or 18:

0   3   6   9  12  15  18
  3   3   3   3   3   3
    0   0   0   0   0

So, the next value of the first history is 18.

Finding all-zero differences for the second history requires an additional sequence:

1   3   6  10  15  21
  2   3   4   5   6
    1   1   1   1
      0   0   0

Then, following the same process as before, work out the next value in each sequence from the bottom up:

1   3   6  10  15  21  28
  2   3   4   5   6   7
    1   1   1   1   1
      0   0   0   0

So, the next value of the second history is 28.

The third history requires even more sequences, but its next value can be found the same way:

10  13  16  21  30  45  68
   3   3   5   9  15  23
     0   2   4   6   8
       2   2   2   2
         0   0   0

So, the next value of the third history is 68.

If you find the next value for each history in this example and add them together, you get 114.

Analyze your OASIS report and extrapolate the next value for each history. What is the sum of these extrapolated values?
*/
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sos := [][]string{}
	sum := 0

	first_element := true
	for _, input_line := range strings.Split(string(content), "\n") {
		sos = [][]string{}
		if input_line == "" {
			continue
		}

		subslice := []string{}

		if first_element {
			subslice = strings.Split(input_line, " ")
			sos = append(sos, subslice)
			first_element = false
		}

		for {
			if !allMembersAreZero(subslice) {
				subslice = createSlice(subslice)
				sos = append(sos, subslice)
			} else {
				break
			}
		}

		sum += extrapolate(&sos)
		first_element = true
	}

	fmt.Printf("Sum of extrapolated values: %d\n", sum)
}

func extrapolate(sos *[][]string) int {
	last_num := 0
	for i := len(*sos) - 1; i >= 0; i-- {
		if i == len(*sos)-1 {
			(*sos)[i] = append((*sos)[i], strconv.Itoa(last_num))
		} else {
			last_number_slice, _ := strconv.Atoi((*sos)[i][len((*sos)[i])-1])
			last_num = last_num + last_number_slice
			(*sos)[i] = append((*sos)[i], strconv.Itoa(last_num))
		}
	}
	return last_num
}

func allMembersAreZero(subslice []string) bool {
	for _, element := range subslice {
		if element != "0" {
			return false
		}
	}
	return true
}

func createSlice(input_slice []string) []string {
	returnSlice := []string{}
	for index, input_number := range input_slice {
		if input_number == "" {
			continue
		}
		if index == 0 {
			continue
		}

		first_number, err := strconv.Atoi(input_number)
		if err != nil {
			panic("NaN!")
		}
		second_number, err := strconv.Atoi(input_slice[index-1])
		if err != nil {
			panic("NaN!")
		}

		returnSlice = append(returnSlice, strconv.Itoa(first_number-second_number))
	}

	return returnSlice
}
