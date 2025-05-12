package main

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	for i := 0; i < len(s)/2; i++ {
		mem := s[i]
		s[i] = s[(len(s) - 1 - i)]
		s[(len(s) - 1 - i)] = mem
	}
}
