// DONE

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

// A list of multipliers is stored to remember how many of each scratchcards we have. These are applied to appropriate winnings.
func Main2() {
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
	var total_count int
	var multipliers []int = make([]int, 0)
	var cur_multiplier int = 0

	for scanner.Scan() {
		total_count += 1 + cur_multiplier
		line := scanner.Text()
		fields := strings.Fields(line)
		fmt.Println(fields)
		// winning := util.CastAll(fields[2:7])
		winning := util.CastAll(fields[2:12])
		// fmt.Println(winning)
		// guesses := util.CastAll(fields[8:])
		guesses := util.CastAll(fields[13:])
		// fmt.Println(guesses)

		correct_count := 0

		for _, guess := range guesses {
			if slices.Contains(winning, guess) {
				correct_count++
			}
		}
		if correct_count > 0 {
			for i := 0; i < correct_count; i++ {
				if len(multipliers) <= i {
					multipliers = append(multipliers, 1+cur_multiplier)
				} else {
					multipliers[i] += 1 + cur_multiplier
				}
			}
			points := 1 << (correct_count - 1)
			// fmt.Println("Count:", correct_count, "Points:", points, "Multiplier:", cur_multiplier)
			sum += points * cur_multiplier
		}
		if len(multipliers) > 0 {
			cur_multiplier = multipliers[0]
			multipliers = multipliers[1:]
		} else {
			cur_multiplier = 0
		}
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Total:", total_count)
}
