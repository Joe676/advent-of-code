// DONE

package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	s "strings"
)

// Lines are split same as in 2.1, but the map now stores max found count, which corresponds to minimum possible count
func Main2() {
	real_count := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	idx_regex, _ := regexp.Compile(`^Game (\d+): (.*)`)
	count_regex, _ := regexp.Compile(`(\d+) (blue|red|green)`)

	var sum int = 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		result := idx_regex.FindStringSubmatch(line)
		idx, _ := strconv.Atoi(result[1])
		fmt.Println("Index:", idx)

		real_count["red"] = 0
		real_count["green"] = 0
		real_count["blue"] = 0

		showings_str := result[2]
		showings := s.Split(showings_str, "; ")

		for _, set := range showings {
			fmt.Println(set)
			count_result := count_regex.FindAllStringSubmatch(set, -1)
			for _, count_color := range count_result {
				count, _ := strconv.Atoi(count_color[1])
				color := count_color[2]

				if real_count[color] < count {
					real_count[color] = count
				}
			}
		}

		sum += real_count["red"] * real_count["green"] * real_count["blue"]

	}

	fmt.Println("Sum:", sum)
}
