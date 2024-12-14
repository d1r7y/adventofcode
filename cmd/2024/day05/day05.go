/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day05

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day05Cmd represents the day05 command
var Day05Cmd = &cobra.Command{
	Use:   "day05",
	Short: `Print Queue`,
	Run: func(cmd *cobra.Command, args []string) {
		df, err := os.Open(utilities.GetInputPath(cmd))
		if err != nil {
			log.Fatal(err)
		}

		defer df.Close()

		fileContent, err := io.ReadAll(df)
		if err != nil {
			log.Fatal(err)
		}
		err = day(cmd, string(fileContent))
		if err != nil {
			log.Fatal(err)
		}
	},
}

type PageList []int

type OrderingRules struct {
	PrecursorMap map[int]PageList
}

func (r *OrderingRules) ParseOrderingRule(line string) {
	precursorRE := regexp.MustCompile(`([0-9]+)\|([0-9]+)`)

	precursorMatches := precursorRE.FindAllStringSubmatch(line, -1)

	precursor, err := strconv.Atoi(precursorMatches[0][1])
	if err != nil {
		return
	}

	page, err := strconv.Atoi(precursorMatches[0][2])
	if err != nil {
		return
	}

	if precursorPages, ok := r.PrecursorMap[page]; ok {
		newPrecursorPages := make(PageList, len(precursorPages))
		copy(newPrecursorPages, precursorPages)
		newPrecursorPages = append(newPrecursorPages, precursor)
		slices.Sort(newPrecursorPages)
		r.PrecursorMap[page] = newPrecursorPages
	} else {
		newPrecursorPages := PageList{precursor}
		r.PrecursorMap[page] = newPrecursorPages
	}
}

func (r *OrderingRules) GetPagePrecursors(page int) PageList {
	if pages, ok := r.PrecursorMap[page]; ok {
		return pages
	}

	return PageList{}
}

func NewOrderingRules() *OrderingRules {
	rules := &OrderingRules{}
	rules.PrecursorMap = make(map[int]PageList)

	return rules
}

type Update struct {
	Pages   PageList
	PageMap map[int]bool
}

func (u *Update) PageInUpdate(page int) bool {
	_, ok := u.PageMap[page]
	return ok
}

func (u *Update) MiddlePage() (int, bool) {
	pageListLength := len(u.Pages)
	if pageListLength == 0 {
		return 0, false
	}
	if pageListLength%2 == 0 {
		return 0, false
	}

	return u.Pages[pageListLength/2], true
}

func (u *Update) ValidOrder(orderingRules *OrderingRules) bool {
	seenPages := make(map[int]bool)

	for _, page := range u.Pages {
		precursorPages := orderingRules.GetPagePrecursors(page)

		for _, precursorPage := range precursorPages {
			if u.PageInUpdate(precursorPage) {
				if _, ok := seenPages[precursorPage]; !ok {
					return false
				}
			}
		}

		seenPages[page] = true
	}

	return true
}

func (u *Update) FixOrder(orderingRules *OrderingRules) {
	pageList := make(PageList, len(u.Pages))
	copy(pageList, u.Pages)

	for {
		seenPages := make(map[int]bool)

		haveSeenPage := func(p int) bool {
			_, ok := seenPages[p]
			return ok
		}

		newPageList := make(PageList, 0)

		reorderedList := false

		for _, page := range pageList {
			if haveSeenPage(page) {
				continue
			}

			precursorPages := orderingRules.GetPagePrecursors(page)

			for _, precursorPage := range precursorPages {
				if u.PageInUpdate(precursorPage) {
					if !haveSeenPage(precursorPage) {
						newPageList = append(newPageList, precursorPage)
						seenPages[precursorPage] = true
						reorderedList = true
					}
				}
			}

			newPageList = append(newPageList, page)
			seenPages[page] = true
		}
		if !reorderedList {
			copy(u.Pages, pageList)
			return
		} else {
			copy(pageList, newPageList)
		}
	}
}

func ParseUpdate(line string) *Update {
	update := &Update{}
	update.PageMap = make(map[int]bool)

	update.Pages = utilities.ParseIntList(line)
	for _, page := range update.Pages {
		update.PageMap[page] = true
	}

	return update
}

func day(_ *cobra.Command, fileContents string) error {
	// Part 1: The Elf must recognize you, because they waste no time explaining that the
	// new sleigh launch safety manual updates won't print correctly. Failure to update the
	// safety manuals would be dire indeed, so you offer your services.
	//
	// Safety protocols clearly indicate that new pages for the safety manuals must be printed
	// in a very specific order. The notation X|Y means that if both page number X and page
	// number Y are to be produced as part of an update, page number X must be printed at some
	// point before page number Y.
	//
	// The Elf has for you both the page ordering rules and the pages to produce in each update
	// (your puzzle input), but can't figure out whether each update has the pages in the right order.

	orderingRules := NewOrderingRules()

	handleOrderingRules := true
	validMiddlePageTotal := 0
	invalidMiddlePageTotal := 0

	for _, line := range strings.Split(fileContents, "\n") {
		if line == "" {
			handleOrderingRules = false
		}

		if handleOrderingRules {
			orderingRules.ParseOrderingRule(line)
		} else {
			update := ParseUpdate(line)
			if update.ValidOrder(orderingRules) {
				middlePage, _ := update.MiddlePage()
				validMiddlePageTotal += middlePage
			} else {
				update.FixOrder(orderingRules)
				middlePage, _ := update.MiddlePage()
				invalidMiddlePageTotal += middlePage
			}
		}
	}

	fmt.Printf("Valid order middle page number sum: %d\n", validMiddlePageTotal)

	// Part 2: For each of the incorrectly-ordered updates, use the page ordering rules to put the
	// page numbers in the right order.
	//
	// Find the updates which are not in the correct order. What do you get if you add up the middle
	// page numbers after correctly ordering just those updates?

	fmt.Printf("Invalid order middle page number sum: %d\n", invalidMiddlePageTotal)

	return nil
}
