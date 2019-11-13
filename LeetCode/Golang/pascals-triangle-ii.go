func getRow(rowIndex int) []int {
    ans := make([]int, rowIndex+1)
    ans[0] = 1
    for i := 1; i <= rowIndex; i++ {
        for j := i; j >= 1; j-- {
            ans[j] += ans[j-1]
        }
    }
    return ans
}