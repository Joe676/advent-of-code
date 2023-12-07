package day7

import (
	"aoc2023/util"
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var new_card_value map[rune]int = map[rune]int{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func newSetHandKind(cards string) int {
	var card_count map[rune]int = make(map[rune]int)

	for _, card := range cards {
		card_count[card] += 1
	}

	vals := make([]int, 0, len(card_count))
	for _, val := range card_count {
		vals = append(vals, val)
	}
	slices.Sort(vals)

	last := vals[len(vals)-1]
	if last == 5 {
		return five_of_kind
	}
	if last == 4 {
		return four_of_kind
	}
	if last == 3 {
		if vals[len(vals)-2] == 2 {
			return full_house
		}
		return three_of_kind
	}
	if last == 2 {
		if vals[len(vals)-2] == 2 {
			return two_pair
		}
		return one_pair
	}

	return high_card
}

func setBestHandKind(h *hand) {
	best_kind := 0
	for k := range new_card_value { //change Js to each card to find the best one
		cards := []rune(h.cards)
		for i := 0; i < len(cards); i++ {
			if cards[i] == 'J' {
				cards[i] = k
			}
		}
		cur_kind := newSetHandKind(string(cards))
		if cur_kind > best_kind {
			best_kind = cur_kind
		}
	}
	h.hand_kind = best_kind
}

// To determine hand's kind, Js are swapped into each card face. The best kind of these changes are then used.
func Main2() {
	file, err := os.Open("day7/input.txt")
	util.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	util.Check(scanner.Err())

	var sum int

	hands := make([]hand, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fields := strings.Fields(line)
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, hand{cards: fields[0], bid: bid})
	}

	for i := 0; i < len(hands); i++ {
		setBestHandKind(&hands[i])
	}

	handsCmp := func(a, b hand) int {
		kind_cmp := cmp.Compare(a.hand_kind, b.hand_kind)
		if kind_cmp == 0 {
			for i := 0; i < len(a.cards); i++ {
				card_cmp := cmp.Compare(new_card_value[rune(a.cards[i])], new_card_value[rune(b.cards[i])])
				if card_cmp != 0 {
					return card_cmp
				}
			}
		}
		return kind_cmp
	}

	slices.SortFunc(hands, handsCmp)

	for i, h := range hands {
		fmt.Println("Rank:", i, "Hand:", h)

		sum += h.bid * (i + 1)
	}
	fmt.Println("Sum:", sum)
}
