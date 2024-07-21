package main

func canPlaceFlowers(flowerbed []int, n int) bool {
    maxFLowers := 0
    previous, next := 0, 0
    for index, value := range flowerbed {
        if index == len(flowerbed) - 1 {
            next = 0
        } else {
            next = flowerbed[index + 1]
        }
        if previous == 0 && next == 0 && value == 0 {
            maxFLowers += 1
            value = 1
        }
        previous = value
    }
    return n <= maxFLowers
}
