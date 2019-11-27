func rob(nums []int) int {
    m := len(nums)
    if m == 0 {
        return 0
    } else if m == 1 {
        return nums[0]
    }
    dp := make([]int, m)
    dp[0], dp[1] = nums[0], nums[1]
    for i := 2; i < m; i++ {
        dp[i] = dp[i-2] + nums[i]
        if i-3 >= 0 && dp[i-3] + nums[i] > dp[i] {
            dp[i] = dp[i-3] + nums[i]
        }
    }
    return max(dp[m-1], dp[m-2])
}

func max (x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}
