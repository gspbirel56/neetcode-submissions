// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) (resultPairs [][]Pair) {
	resultPairs = make([][]Pair, len(pairs))
	sortedList := []Pair{}
	for i, currentPair := range pairs {
		inserted := false
		for sorted_i, sortedCurrentPair := range sortedList {
			if currentPair.Key < sortedCurrentPair.Key {
				// do the insertion
				newSortedList := make([]Pair, len(sortedList)+1)
				for new_i := 0; new_i < sorted_i; new_i++ {
					newSortedList[new_i] = sortedList[new_i]
				}
				newSortedList[sorted_i] = currentPair
				for new_i := sorted_i; new_i < len(sortedList); new_i++ {
					newSortedList[new_i+1] = sortedList[new_i]
				}
				//fmt.Println(newSortedList)
				sortedList = newSortedList
				inserted = true
				break
			}
		}
		if !inserted {
			sortedList = append(sortedList, currentPair)
		}
		// debug: include the extra empty pair as a separator
		// resultPairs[i] = append(append(sortedList, Pair{0, "\"|\""}), pairs[i+1:]...)
		// non debug: actual resultPairs[i] formation
		resultPairs[i] = append(sortedList, pairs[i+1:]...)
		fmt.Println(resultPairs[i])
	}
	return
}
