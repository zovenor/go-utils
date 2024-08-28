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

func RemoveItemByIndex[T any](l *[]T, index int) (T, error) {
	if len(*l) <= index {
		return *new(T), fmt.Errorf("index out of range")
	}
	newList := make([]T, 0, len(*l)-1)
	newList = append(newList, (*l)[:index]...)
	if len(*l) > index+1 {
		newList = append(newList, (*l)[index+1:]...)
	}
	removedItem := (*l)[index]
	*l = newList
	return removedItem, nil
}
