package main

func singleNumber2(nums []int) int {
	value := int32(0)
	for i := 0; i < 32; i++ {
		oneBit := 0
		zeroBit := 0
		for _, num := range nums {
			if ((num >> i) & 1) == 1 {
				oneBit++
			} else {
				zeroBit++
			}
		}
		if oneBit%3 != 0 {
			value = value | (1 << i)
		}
	}
	return int(value)
}
