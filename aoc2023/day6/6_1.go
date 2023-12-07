package day6

import (
	"aoc2023/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func s(v int, t int) int {
	return v * t
}

// For each race a minimum throttle time and maximum throttle time are found (looking from top and from bottom of the range)
// These are just subtracted from each other.  
func Main1() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var times []int
	var distances []int
	var sum int = 1

	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)
	fields := strings.Fields(line)
	times = util.CastAll(fields[1:])

	scanner.Scan()
	line = scanner.Text()
	fmt.Println(line)
	fields = strings.Fields(line)
	distances = util.CastAll(fields[1:])

	fmt.Println(times)
	fmt.Println(distances)

	for i := 0; i < len(times); i++ {
		t := times[i]
		d := distances[i]
		var min_idx, max_idx int
		for j := 0; j <= t; j++ {
			cur_s := s(j, t-j)
			if cur_s > d {
				min_idx = j
				break
			}
		}
		for j := t; j >= 0; j-- {
			cur_s := s(j, t-j)
			if cur_s > d {
				max_idx = j
				break
			}
		}
		sum *= max_idx - min_idx + 1
	}

	fmt.Println(sum)
}
