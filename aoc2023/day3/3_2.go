package day3

import (
	"aoc2023/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func Main2() {
	file, err := os.Open("day3/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	number_regex, _ := regexp.Compile(`(\d+)`)
	symbol_regex, _ := regexp.Compile(`\*`)
	symbols_map := make(map[int][]int)
	numbers := make([]int, 0, 10)
	numbers_line := make([]int, 0, 10)
	numbers_idx := make([][]int, 0, 10)
	line_number := 0
	// line_width := 0

	for scanner.Scan() {
		line := scanner.Text()
		// line_width = len(line)
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

	for y, xs := range symbols_map {
		for _, x := range xs {
			fmt.Println(x, y)
		}
	}

	fmt.Println("Sum:", sum)
}
