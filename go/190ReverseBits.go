package main

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i <= 31; i++ {
		result = result | (num&1)<<(31-i)
		num = num >> 1
	}
	return result
}
