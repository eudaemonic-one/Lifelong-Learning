func singleNumber(nums []int) int {
    var ans = 0
    for i := range nums {
        ans ^= nums[i]
    }
    return ans
}
