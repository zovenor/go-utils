package filter

import "fmt"

func RemoveItemsFromListByCondition[T any](l *[]T, checkRemoving func(item T) bool) error {
	newList := make([]T, 0)
	f := false
	for _, item := range *l {
		if checkRemoving(item) {
			continue
		}
		newList = append(newList, item)
		f = true
	}
	*l = newList
	if !f {
		return fmt.Errorf("items not found")
	} else {
		return nil
	}
}
