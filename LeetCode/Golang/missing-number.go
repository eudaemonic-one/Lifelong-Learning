func missingNumber(nums []int) int {
    res := len(nums)
    for i, num := range nums {
        res = res ^ i ^ num
    }
    return res
}
