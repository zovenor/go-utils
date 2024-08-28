package tests

import (
	"testing"

	"github.com/zovenor/go-utils/filter"
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

func TestFilterDeleteByIndex(t *testing.T) {
	testData := []struct {
		ListItems             []string
		ListItemsUponDeleting []string
		Index                 int
		DeletedItem           string
	}{
		{
			ListItems:             []string{"string", "a", "wef", "qwerty"},
			ListItemsUponDeleting: []string{"string", "wef", "qwerty"},
			Index:                 1,
			DeletedItem:           "a",
		},
	}
	for _, td := range testData {
		copiedList := append([]string(nil), td.ListItems...)
		deletedItem, err := filter.RemoveItemByIndex(&td.ListItems, td.Index)
		if err != nil {
			t.Fatalf("filter delete by index error: %v", err)
		}
		if deletedItem != td.DeletedItem {
			t.Fatalf("deleted item %v not equal %v", deletedItem, td.DeletedItem)
		}
		if !equalLists(td.ListItems, td.ListItemsUponDeleting) {
			t.Fatalf("%v not equal %v", td.ListItems, copiedList)
		}
	}
}
