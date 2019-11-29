func minSubArrayLen(s int, nums []int) int {
    const INT_MAX int = int(^uint(0) >> 1)
    res, sum, l, r := INT_MAX, 0, 0, 0
    for ; r < len(nums); r++ {
        sum += nums[r]
        for sum >= s {
            res = min(res, r-l+1)
            sum -= nums[l]
            l++
        }
    }
    if res == INT_MAX {
        return 0
    }
    return res
}

func min (x, y int) int {
    if x < y {
        return x
    }
    return y
}
