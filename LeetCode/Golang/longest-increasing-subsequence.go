// Approach 4 Dynamic Programming with Binary Search (Final tricks to reduce time complexity)

func lengthOfLIS(nums []int) int {
    length := 0
    dp := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        insertionPos := binarySearch(dp, 0, length, nums[i])
        dp[insertionPos] = nums[i]
        if insertionPos == length {
            length++
        }
    }
    return length
}

func binarySearch(nums []int, start, end, key int) int {
    l, r := start, end
    for l < r {
        m := (l+r) / 2
        if nums[m] < key {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}

// Approach 3 Dynamic Programming (Bottom-up dynamic programming)

// func lengthOfLIS(nums []int) int {
//     if len(nums) == 0 {
//         return 0
//     }
//     res := 1
//     dp := make([]int, len(nums))
//     dp[0] = 1
//     for j := 1; j < len(nums); j++ {
//         maxSeqLen := 0
//         for i := 0; i < j; i++ {
//             if nums[i] < nums[j] {
//                 maxSeqLen = max(maxSeqLen, dp[i])
//             }
//         }
//         dp[j] = maxSeqLen + 1
//         res = max(res, dp[j])
//     }
//     return res
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

// Approach 2 Recursion with Memoization (Top-down dynamic programming)

// func lengthOfLIS(nums []int) int {
//     memo := make([][]int, len(nums))
//     for i := 0; i < len(nums); i++ {
//         memo[i] = make([]int, len(nums))
//         for j := 0; j < len(nums); j++ {
//             memo[i][j] = -1
//         }
//     }
//     return dfs(nums, memo, -1, 0)
// }

// func dfs(nums []int, memo [][]int, prev int, curr int) int {
//     if curr == len(nums) {
//         return 0
//     }
//     if memo[prev+1][curr] >= 0 {
//         return memo[prev+1][curr]
//     }
//     taken, notTaken := 0, 0
//     if prev < 0 || nums[curr] > nums[prev] {
//         taken = 1 + dfs(nums, memo, curr, curr+1)
//     }
//     notTaken = dfs(nums, memo, prev, curr+1)
//     memo[prev+1][curr] = max(taken, notTaken)
//     return memo[prev+1][curr]
// }
    
// Approach 1 Brute Force (Recursive backtracking solution)

// func lengthOfLIS(nums []int) int {
//     const INT_MIN int = ^int(^uint(0) >> 1)
//     return dfs(nums, INT_MIN, 0)
// }

// func dfs(nums []int, prev int, idx int) int {
//     if idx == len(nums) {
//         return 0
//     }
//     taken, notTaken := 0, 0
//     if nums[idx] > prev {
//         taken = dfs(nums, nums[idx], idx+1)
//     }
//     notTaken = dfs(nums, prev, idx+1)
//     return max(taken, notTaken)
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }
