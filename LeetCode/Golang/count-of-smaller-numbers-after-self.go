func countSmaller(nums []int) []int {
    sorted := make([]int, 0)
    res := make([]int, len(nums))
    for i := len(nums)-1; i >= 0; i-- {
        idx := binarySearch(sorted, nums[i])
        res[i] = idx
        if idx >= len(sorted) {
            sorted = append(sorted, nums[i])
        } else {
            tmp := sorted[idx:]
            sorted = append([]int{}, sorted[:idx]...)
            sorted = append(sorted, nums[i])
            sorted = append(sorted, tmp...)
        }
    }
    return res
}

func binarySearch(nums []int, key int) int {
    if len(nums) == 0 {
        return 0
    }
    if key < nums[0] {
        return 0
    } else if key > nums[len(nums)-1] {
        return len(nums)
    }
    l, r := 0, len(nums)
    for l < r {
        m := (l+r) / 2
        if key > nums[m] {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}
