package filter

import "fmt"

func RemoveItemsFromListByCondition[T any](l *[]T, checkRemoving func(item T) bool) ([]T, error) {
	newList := make([]T, 0)
	removedItems := make([]T, 0)
	for _, item := range *l {
		if checkRemoving(item) {
			removedItems = append(removedItems, item)
			continue
		}
		newList = append(newList, item)
	}
	*l = newList
	if len(removedItems) == 0 {
		return make([]T, 0), fmt.Errorf("items not found")
	} else {
		return removedItems, nil
	}
}
