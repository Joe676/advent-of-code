// IN PROGRESS

package day5

import (
	"aoc2023/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Main2() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	seeds_line := scanner.Text()
	seeds := util.CastAll(strings.Fields(seeds_line)[1:])
	sum := 0
	for i, seed := range seeds {
		sum += seed * (i % 2)
	}

	fmt.Println(seeds)
	fmt.Println(sum)
	// seed_paths := make([][]int, 0)
	// for _, seed := range seeds {
	// 	path := make([]int, 0, 1)
	// 	seed_paths = append(seed_paths, append(path, seed))
	// }
	// fmt.Println("Seeds:", seeds)

	// cur_map := 0
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	// fmt.Println(line, "Len:", len(line))
	// 	if len(line) == 0 {
	// 		continue
	// 	}
	// 	cur_range := util.CastAll(strings.Fields(line))
	// 	// fmt.Println("Current Map:", cur_map, "Current Range:", cur_range)
	// 	if len(cur_range) < 3 {

	// 		for i, path := range seed_paths {
	// 			if len(path) == cur_map {
	// 				seed_paths[i] = append(seed_paths[i], path[cur_map-1])
	// 			}
	// 		}
	// 		cur_map += 1
	// 		continue
	// 	}
	// 	for i, path := range seed_paths {
	// 		if len(path) > cur_map {
	// 			continue
	// 		}
	// 		if path[cur_map-1] >= cur_range[1] && path[cur_map-1] < cur_range[1]+cur_range[2] {
	// 			offset := path[cur_map-1] - cur_range[1]
	// 			seed_paths[i] = append(seed_paths[i], cur_range[0]+offset)
	// 		}
	// 	}
	// }

	// min_last := math.MaxInt

	// for _, path := range seed_paths {
	// 	if path[len(path)-1] < min_last {
	// 		min_last = path[len(path)-1]
	// 	}
	// 	fmt.Println(path)
	// }
	// fmt.Println("Min last:", min_last)
}
