// DONE

package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	s "strings"
)

// Each line is scanned from left and right in search of one of keys, that are then mapped to values
func Main2() {
	code := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"0":     0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var first int = -1
		var first_idx int = 100
		var last int
		var last_idx int = -1

		for k, v := range code {
			if !s.Contains(line, k) {
				continue
			}

			f_idx := s.Index(line, k)
			l_idx := s.LastIndex(line, k)

			if f_idx < first_idx {
				first = v
				first_idx = f_idx
			}

			if l_idx > last_idx {
				last = v
				last_idx = l_idx
			}
		}

		number := first*10 + last
		fmt.Println("Line:", line, "Number encoded:", number)
		sum += number
	}
	fmt.Printf("The sum: %d\n", sum)

}
