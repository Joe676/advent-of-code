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

func Main1() {
	real_count := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
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

		var is_legal bool = true

		showings_str := result[2]
		showings := s.Split(showings_str, "; ")
		for _, set := range showings {
			fmt.Println(set)
			count_result := count_regex.FindAllStringSubmatch(set, -1)
			for _, count_color := range count_result {
				count, _ := strconv.Atoi(count_color[1])
				color := count_color[2]
				
				fmt.Println("Color:", color, "Count:", count, "Max count:", real_count[color])
				if count > real_count[color] {
					is_legal = false
				}
			}
		}

		if is_legal {
			sum += idx
		}

	}

	fmt.Println("Sum:", sum)
}
