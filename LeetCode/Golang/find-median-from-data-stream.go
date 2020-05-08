type MedianFinder struct {
    nums []int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{}
}


func (this *MedianFinder) AddNum(num int)  {
    if len(this.nums) == 0 {
        this.nums = append(this.nums, num)
    } else {
        idx := lowerBound(this.nums, len(this.nums), num)
        rear := append([]int{}, this.nums[idx:]...)
        this.nums = append(this.nums[:idx], num)
        this.nums = append(this.nums, rear...)
    }
}


func lowerBound(nums []int, size, key int) int {
    first := 0
    for size > 0 {
        half := size >> 1
        middle := first + half
        if nums[middle] < key {
            first = middle + 1
            size = size - half - 1
        } else {
            size = half
        }
    }
    return first
}


func (this *MedianFinder) FindMedian() float64 {
    L := len(this.nums)
    if L % 2 == 1 {
        return float64(this.nums[L/2])
    }
    return (float64(this.nums[L/2-1]) + float64(this.nums[L/2])) / 2
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
