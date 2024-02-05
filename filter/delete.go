package filter

func RemoveItemsFromListByCondition[T any](l *[]T, checkRemoving func(i int) bool) {
	newList := make([]T, 0)
	for i, item := range *l {
		if checkRemoving(i) {
			continue
		}
		newList = append(newList, item)
	}
	*l = newList
}
