package main

func isHappy(n int) bool {
	mapping := make(map[int]bool)
	for {
		square := 0
		for n > 0 {
			digit := n % 10
			square += digit * digit
			n = n / 10
		}
		if square == 1 {
			return true
		}
		if _, ok := mapping[square]; ok {
			return false
		} else {
			mapping[square] = true
		}
		n = square
	}
}
