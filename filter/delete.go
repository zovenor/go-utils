package filter

func RemoveItemsFromListByCondition[T any](l *[]T, checkRemoving func(item T) bool) {
	newList := make([]T, 0)
	for _, item := range *l {
		if checkRemoving(item) {
			continue
		}
		newList = append(newList, item)
	}
	*l = newList
}
