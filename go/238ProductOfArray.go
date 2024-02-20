package main

func productExceptSelf(nums []int) []int {
  res := make([]int, len(nums))
  res[0], res[len(nums)-1] = 1, 1

  for i := 1; i < len(nums); i++ {
    res[i] = res[i-1] * nums[i-1]
  }
  
  cache := 1
  for i := len(nums) - 2; i >= 0; i-- {
    cache *= nums[i+1]
    res[i] *= cache
  }
  return res

}
