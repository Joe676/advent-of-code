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

const (
	unknown_hand = iota
	high_card
	one_pair
	two_pair
	three_of_kind
	full_house
	four_of_kind
	five_of_kind
)

type hand struct {
	cards     string
	bid       int
	hand_kind int
}

var card_value map[byte]int = map[byte]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func setHandKind(h *hand) {
	var card_count map[rune]int = make(map[rune]int)

	for _, card := range h.cards {
		card_count[card] += 1
	}

	vals := make([]int, 0, len(card_count))
	for _, val := range card_count {
		vals = append(vals, val)
	}
	slices.Sort(vals)

	last := vals[len(vals)-1]
	if last == 5 {
		h.hand_kind = five_of_kind
		return
	}
	if last == 4 {
		h.hand_kind = four_of_kind
		return
	}
	if last == 3 {
		if vals[len(vals)-2] == 2 {
			h.hand_kind = full_house
			return
		}
		h.hand_kind = three_of_kind
		return
	}
	if last == 2 {
		if vals[len(vals)-2] == 2 {
			h.hand_kind = two_pair
			return
		}
		h.hand_kind = one_pair
		return
	}

	h.hand_kind = high_card
}

func Main1() {
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
		setHandKind(&hands[i])
		// fmt.Println(hands[i])
	}

	handsCmp := func(a, b hand) int {
		kind_cmp := cmp.Compare(a.hand_kind, b.hand_kind)
		if kind_cmp == 0 {
			for i := 0; i < len(a.cards); i++ {
				card_cmp := cmp.Compare(card_value[a.cards[i]], card_value[b.cards[i]])
				if card_cmp != 0 {
					return card_cmp
				}
			}
			// return cmp.Compare(a.cards, b.cards)
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
