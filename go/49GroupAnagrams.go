package main

type Key [26]rune

func generateKey(str string) Key {
	var key Key
	for _, val := range str {
		key[val-'a']++
	}
	return key
}
func groupAnagrams(strs []string) [][]string {
	mapping := make(map[Key][]string)
	for _, val := range strs {
		key := generateKey(val)
		mapping[key] = append(mapping[key], val)
	}

	solution := make([][]string, 0, len(mapping))
	for _, value := range mapping {
		solution = append(solution, value)
	}

	return solution
}

// slow Solution O(n2)
func groupAnagramsSlow(strs []string) [][]string {
	var solution [][]string
	for index := 0; index < len(strs); index++ {
		mapping := make(map[rune]int)
		currentAna := []string{strs[index]}
		for _, char := range strs[index] {
			if val, ok := mapping[char]; ok {
				mapping[char] = val + 1
			} else {
				mapping[char] = 1
			}
		}
		for i := index + 1; i < len(strs); i++ {
			if len(strs[i]) == len(strs[index]) {
				mappingClone := cloneMap(&mapping)
				stop := false
				for _, char := range strs[i] {
					if _, ok := mapping[char]; ok {
						mappingClone[char] -= 1
					} else {
						stop = true
						break
					}
				}
				if !stop && ((mapIsEmpty(mappingClone) && strs[index] != "") || (strs[index] == strs[i])) {
					currentAna = append(currentAna, strs[i])
					strs = removeFromArray(i, strs)
					i--
				}
			}
		}
		solution = append(solution, currentAna)
	}
	return solution
}

func mapIsEmpty(mapping map[rune]int) bool {
	for _, val := range mapping {
		if val != 0 {
			return false
		}
	}
	return true
}

func cloneMap(mapping *map[rune]int) map[rune]int {
	mapClone := make(map[rune]int)
	for key, value := range *mapping {
		mapClone[key] = value
	}
	return mapClone
}

func removeFromArray(index int, array []string) []string {
	arrayClone := make([]string, 0, len(array)-1)
	arrayClone = append(arrayClone, array[:index]...)
	arrayClone = append(arrayClone, array[index+1:]...)
	return arrayClone
}
