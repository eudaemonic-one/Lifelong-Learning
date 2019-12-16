func maxCoins(nums []int) int {
    N := len(nums)
    if N == 0 {
        return 0
    }
    if N == 1 {
        return nums[0]
    }
    // reframe array nums
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)
    N += 2
    // initialize array dp
    dp := make([][]int, N)
    for i := 0; i < N; i++ {
        dp[i] = make([]int, N)
    }
    // iterate over dp
    for left := N-2; left >= 0; left-- {
        for right := left+2; right < N; right++ {
            coins := 0
            for i := left+1; i < right; i++ {
               coins = max(coins, nums[left]*nums[i]*nums[right]+dp[left][i]+dp[i][right])
            }
            dp[left][right] = coins
        }
    }
    return dp[0][N-1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
