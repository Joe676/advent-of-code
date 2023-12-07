package util

import (
	"log"
	"strconv"
)

func CastAll(in []string) []int {
	var out []int = make([]int, 0)
	for _, str := range in {
		cast, _ := strconv.Atoi(str)
		out = append(out, cast)
	}
	return out
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
