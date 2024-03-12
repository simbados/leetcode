package main

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, val := range letters {
		if val > target {
			return val
		}
	}
	return letters[0]
}

