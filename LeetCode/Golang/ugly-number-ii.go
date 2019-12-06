func nthUglyNumber(n int) int {
    twos, threes, fives := 0, 0, 0
    nums := []int{1}
    for i := 1; i < n; i++ {
        ugly := min(nums[twos] * 2, min(nums[threes] * 3, nums[fives] * 5))
        nums = append(nums, ugly)
        if ugly == nums[twos] * 2 {
            twos++
        }
        if ugly == nums[threes] * 3 {
            threes++
        }
        if ugly == nums[fives] * 5 {
            fives++
        }
    }
    return nums[n-1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
