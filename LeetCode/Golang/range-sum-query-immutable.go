type NumArray struct {
    sums []int
}


func Constructor(nums []int) NumArray {
    if len(nums) == 0 {
        return NumArray{}
    }
    numArray := NumArray{make([]int, len(nums))}
    numArray.sums[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        numArray.sums[i] = numArray.sums[i-1] + nums[i]
    }
    return numArray
}


func (this *NumArray) SumRange(i int, j int) int {
    iSum, jSum := 0, this.sums[j]
    if i > 0 {
        iSum = this.sums[i-1]
    }
    return jSum - iSum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
