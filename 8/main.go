package main

/*
--- Day 8: Haunted Wasteland ---

You're still riding a camel across Desert Island when you spot a sandstorm quickly approaching. When you turn to warn the Elf, she disappears before your eyes! To be fair, she had just finished warning you about ghosts a few minutes ago.

One of the camel's pouches is labeled "maps" - sure enough, it's full of documents (your puzzle input) about how to navigate the desert. At least, you're pretty sure that's what they are; one of the documents contains a list of left/right instructions, and the rest of the documents seem to describe some kind of network of labeled nodes.

It seems like you're meant to use the left/right instructions to navigate the network. Perhaps if you have the camel follow the same instructions, you can escape the haunted wasteland!

After examining the maps for a bit, two nodes stick out: AAA and ZZZ. You feel like AAA is where you are now, and you have to follow the left/right instructions until you reach ZZZ.

This format defines each node of the network individually. For example:

RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)

Starting with AAA, you need to look up the next element based on the next left/right instruction in your input. In this example, start with AAA and go right (R) by choosing the right element of AAA, CCC. Then, L means to choose the left element of CCC, ZZZ. By following the left/right instructions, you reach ZZZ in 2 steps.

Of course, you might not find ZZZ right away. If you run out of left/right instructions, repeat the whole sequence of instructions as necessary: RL really means RLRLRLRLRLRLRLRL... and so on. For example, here is a situation that takes 6 steps to reach ZZZ:

LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)

Starting at AAA, follow the left/right instructions. How many steps are required to reach ZZZ?
*/
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	score := findSolution(string(content))
	fmt.Printf("To reach ZZZ you need %v steps\n", score)
}

func findSolution(input string) int {

	m := map[string][]string{}
	first_line := ""

	for index, line := range strings.Split(input, "\n") {
		if index == 0 {
			first_line = line
			continue
		}

		if line == "" {
			continue
		}

		splitted := strings.Split(line, " = ")
		node_name := splitted[0]
		right_side_split := strings.Replace(splitted[1], "(", "", 1)
		right_side_split = strings.Replace(right_side_split, ")", "", 1)

		left_side := strings.Split(right_side_split, ", ")[0]
		right_side := strings.Split(right_side_split, ", ")[1]

		m[node_name] = []string{left_side, right_side}
	}

	current := "AAA"
	score := 0

	for current != "ZZZ" {
		for _, direction := range first_line {
			score++
			if direction == 'L' {
				current = m[current][0]
			} else {
				current = m[current][1]
			}

			if current == "ZZZ" {
				break
			}
		}
	}
	return score
}
