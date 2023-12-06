package day4

import (
	"aoc2023/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Main1() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		fmt.Println(fields)
		winning := util.CastAll(fields[2:12])
		fmt.Println(winning)
		guesses := util.CastAll(fields[13:])
		fmt.Println(guesses)

		correct_count := 0

		for _, guess := range guesses {
			if slices.Contains(winning, guess) {
				correct_count++
			}
		}
		if correct_count > 0 {
			winnings := 1 << (correct_count - 1)
			sum += winnings
		}
	}

	fmt.Println("Sum:", sum)
}
