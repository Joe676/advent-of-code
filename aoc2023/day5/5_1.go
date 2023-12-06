package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Main1() {
	file, err := os.Open("day5/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// fields := strings.Fields(line)
	}
}
