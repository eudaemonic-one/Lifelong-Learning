func findMin(nums []int) int {
    l, r := 0, len(nums)-1
    // if there is no rotation
    if nums[r] > nums[0] {
        return nums[0]
    }
    // search for the turning point
    for l < r {
        m := (l+r) / 2
        if nums[m] > nums[r] {
            l = m + 1
        } else if nums[m] < nums[l] {
            r = m
            l += 1
        } else {
            r -= 1
        }
    }
    return nums[l]
}
