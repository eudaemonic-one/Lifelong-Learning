func numSquares(n int) int {
    const INT_MAX int = int(^uint(0) >> 1)
    NSqrt := int(math.Sqrt(float64(n)))
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = INT_MAX
        for sq := 1; sq <= NSqrt; sq++ {
            if i < sq * sq {
                break
            }
            if dp[i-sq*sq] + 1 < dp[i] {
                dp[i] = dp[i-sq*sq] + 1
            }
        }
    }
    return dp[n]
}
