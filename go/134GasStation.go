package main

func canCompleteCircuit(gas []int, cost []int) int {
    if (len(gas) == 1) {
        if gas[0] >= cost[0] {
            return 0
        } else {
            return -1
        }
    }
    cache := make([]int, len(gas))
    for i, g := range gas {
        cache[i] = g - cost[i]
    }
    for i := 0; i < len(cache); i++ {
        currentValue := cache[i]
        if (currentValue <= 0) {
            continue;
        }
        for j := 1; j < len(cache); j++ {
            currentValue += cache[(i + j) % len(cache)]
            if (currentValue < 0) {
                break;
            }
            if j == len(cache) - 1 {
                return i
            }
        }
    }
    return -1
}
