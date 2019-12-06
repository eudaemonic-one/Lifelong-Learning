func validTree(n int, edges [][]int) bool {
    // a valid tree means n == e - 1 and acyclic
    if n-1 != len(edges) {
        return false
    }
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = -1
    }
    for i := 0; i < n-1; i++ {
        x := find(nums, edges[i][0])
        y := find(nums, edges[i][1])
        if x == y {
            return false
        }
        nums[x] = y
    }
    return true
}

func find(nums []int, i int) int {
    if nums[i] == -1 {
        return i
    }
    return find(nums, nums[i])
}
