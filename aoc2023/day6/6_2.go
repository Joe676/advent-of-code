package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// The values are just stuck together and processed same as 6.1
func Main2() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var time int
	var distance int
	var sum int = 1

	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)
	fields := strings.Fields(line)
	joined := strings.Join(fields[1:], "")
	time, _ = strconv.Atoi(joined)

	scanner.Scan()
	line = scanner.Text()
	fmt.Println(line)
	fields = strings.Fields(line)
	joined = strings.Join(fields[1:], "")
	distance, _ = strconv.Atoi(joined)

	fmt.Println(time)
	fmt.Println(distance)

	var min_idx int
	var max_idx int
	for j := 0; j <= time; j++ {
		cur_s := s(j, time-j)
		if cur_s > distance {
			min_idx = j
			break
		}
	}
	for j := time; j >= 0; j-- {
		cur_s := s(j, time-j)
		if cur_s > distance {
			max_idx = j
			break
		}
	}
	sum *= max_idx - min_idx + 1

	fmt.Println(sum)
}
