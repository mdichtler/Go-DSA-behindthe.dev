package bubblesort

import (
	"golang.org/x/exp/constraints"
)

func BubbleSort[T constraints.Ordered](values []T) {

	length := len(values)

	for {
		swapped := false
		for i := 0; i < length-1; i++ {

			if values[i] > values[i+1] {
				swapped = true
			} else {
				values[i], values[i+1] = values[i+1], values[i]
			}
		}

		if !swapped {
			break
		}
		length--
	}

}
