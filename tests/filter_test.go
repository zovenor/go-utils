package tests

import (
	"github.com/zovenor/go-utils/filter"
	"testing"
)

type filterDeleteTestData struct {
	ListItems    []string
	DeletedItems []string
	Correct      bool
}

func equalLists[T comparable](l0, l1 []T) bool {
	if len(l0) != len(l1) {
		return false
	} else if len(l0) == 0 && len(l1) == 0 {
		return true
	}
	for _, el0 := range l0 {
		exists := false
		for _, el1 := range l1 {
			if el1 == el0 {
				exists = true
			}
		}
		if !exists {
			return false
		}
	}
	return true
}

func TestFilterDelete(t *testing.T) {
	testData := []filterDeleteTestData{
		{
			ListItems:    []string{"string", "a", "wef", "qwerty"},
			DeletedItems: []string{"a"},
			Correct:      true,
		},
		{
			ListItems:    []string{"string", "a", "wef", "qwerty"},
			DeletedItems: []string{"a", "wef"},
			Correct:      true,
		},
		{
			ListItems:    []string{"string", "a", "wef", "qwerty", "a"},
			DeletedItems: []string{"a", "a"},
			Correct:      true,
		},
		{
			ListItems:    []string{"a", "a", "a", "a", "a"},
			DeletedItems: []string{"a", "a", "a", "a"},
			Correct:      false,
		},
	}
	for _, td := range testData {
		deletedItems, err := filter.RemoveItemsFromListByCondition(&td.ListItems, func(item string) bool {
			for _, el := range td.DeletedItems {
				if el == item {
					return true
				}
			}
			return false
		})
		if err != nil {
			t.Fatalf("filter delete error: %v", err)
		}
		if equalLists(deletedItems, td.DeletedItems) != td.Correct {
			if td.Correct {
				t.Fatalf("%v not equal %v", deletedItems, td.DeletedItems)
			} else {
				t.Fatalf("%v equal %v", deletedItems, td.DeletedItems)
			}
		}
	}
}
