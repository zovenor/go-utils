package filter

import "fmt"

func RemoveItemsFromListByCondition[T any](l *[]T, checkRemoving func(item T) bool) error {
	newList := make([]T, 0)
	f := false
	for _, item := range *l {
		if checkRemoving(item) {
			f = true
			continue
		}
		newList = append(newList, item)
	}
	*l = newList
	if !f {
		return fmt.Errorf("items not found")
	} else {
		return nil
	}
}
