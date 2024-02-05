package finder

import (
	"fmt"
)

func FindItemBy[T any](l *[]T, by func(item T) bool) (T, error) {
	for _, item := range *l {
		if by(item) {
			return item, nil
		}
	}
	return *new(T), fmt.Errorf("item not found")
}
