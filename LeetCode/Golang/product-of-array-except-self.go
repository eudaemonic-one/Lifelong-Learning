func productExceptSelf(nums []int) []int {
    output := make([]int, len(nums))
    output[0] = 1
    for i := 1; i < len(nums); i++ {
        output[i] = output[i-1] * nums[i-1]
    }
    rightProduct := nums[len(nums)-1]
    for i := len(nums)-2; i >= 0; i-- {
        output[i] = output[i] * rightProduct
        rightProduct *= nums[i]
    }
    return output
}
