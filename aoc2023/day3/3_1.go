// DONE

package day3

import (
	"aoc2023/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Main1() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	number_regex, _ := regexp.Compile(`(\d+)`)
	symbol_regex, _ := regexp.Compile(`([^\d.\n])`)
	symbols_map := make(map[int][]int)
	numbers := make([]int, 0, 10)
	numbers_line := make([]int, 0, 10)
	numbers_idx := make([][]int, 0, 10)
	line_number := 0
	line_width := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_width = len(line)
		fmt.Println(line)
		number_indeces := number_regex.FindAllStringIndex(line, -1)
		number_result := number_regex.FindAllString(line, -1)
		symbol_result := symbol_regex.FindAllStringIndex(line, -1)

		numbers = append(numbers, util.CastAll(number_result)...)
		numbers_idx = append(numbers_idx, number_indeces...)
		for range number_result {
			numbers_line = append(numbers_line, line_number)
		}

		symbol_indices := make([]int, 0, 6)
		for _, sym := range symbol_result {
			symbol_indices = append(symbol_indices, sym[0])
		}

		symbols_map[line_number] = append(symbols_map[line_number], symbol_indices...)

		line_number++
	}

	fmt.Println("symbols_map:", symbols_map)
	fmt.Println("numbers:", numbers)
	fmt.Println("numbers_idx:", numbers_idx)
	fmt.Println("numbers_line:", numbers_line)
	var sum int = 0
	// Find numbers next to symbols
	for i, number := range numbers {
		line_number = numbers_line[i]
		index := numbers_idx[i]
		var is_part bool = false
		for y_offset := -1; y_offset <= 1; y_offset++ {
			y := line_number + y_offset
			if y < 0 {
				y = 0
			}
			low_x := index[0] - 1
			if low_x < 0 {
				low_x = 0
			}
			high_x := index[1] //CAUTION! the top indexes from regexp.FindAllStringIndex are one higher than the real index!
			if high_x >= line_width {
				high_x = line_width - 1
			}

			for _, symbol_x := range symbols_map[y] {
				fmt.Println("Checking", symbol_x, "in line", y, "against", low_x, high_x)
				if symbol_x >= low_x && symbol_x <= high_x {
					is_part = true
					fmt.Println("TRUE")
					break
				}
			}
			if is_part {
				break
			}
		}
		if is_part {
			fmt.Println(number)
			sum += number
		}
	}

	fmt.Println("Sum:", sum)

}
