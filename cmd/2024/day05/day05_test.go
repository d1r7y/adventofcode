/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day05

import (
	"fmt"
	"strings"
	"testing"
)

func TestUpdateValidity(t *testing.T) {
	type testCase struct {
		text string
	}
	testCases := []testCase{
		{
			text: `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`,
		},
	}

	for _, test := range testCases {
		orderingRules := NewOrderingRules()
		handleOrderingRules := true
		for _, line := range strings.Split(test.text, "\n") {
			if line == "" {
				handleOrderingRules = false
				continue
			}

			if handleOrderingRules {
				orderingRules.ParseOrderingRule(line)
			} else {
				update := ParseUpdate(line)
				if update.ValidOrder(orderingRules) {
					middlePage, _ := update.MiddlePage()
					fmt.Printf("Middle page: %d\n", middlePage)
				} else {
					update.FixOrder(orderingRules)
					middlePage, _ := update.MiddlePage()
					fmt.Printf("Fixed middle page: %d\n", middlePage)
				}
			}
		}
	}
}
