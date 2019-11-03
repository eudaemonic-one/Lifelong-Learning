func twoSum(nums []int, target int) []int {
    var dict = make(map[int]int)
    var complement int
    for i := range nums {
        complement = target - nums[i]
        if num, ok := dict[complement]; ok {
            return []int{num, i}
        }
        dict[nums[i]] = i
    }
    return nil
}