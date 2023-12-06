package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Main1() {
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
		fmt.Println(line)

		var first int = -1
		var last int

		for i, l := range line {
			if l >= '0' && l <= '9' {
				fmt.Printf("Found %c at %d\n", l, i)
				if first < 0 {
					first = int(l - '0')
				}
				last = int(l - '0')
			}
		}

		number := first*10 + last
		fmt.Println("Line:", line, "Number encoded:", number)
		sum += number
	}
	fmt.Printf("The sum: %d\n", sum)
}
