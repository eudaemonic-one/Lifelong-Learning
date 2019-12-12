type NumArray struct {
    tree []int
    n int
}


func Constructor(nums []int) NumArray {
    numArray := NumArray{}
    if len(nums) > 0 {
        numArray.n = len(nums)
        numArray.tree = make([]int, len(nums)*2) // segment tree
        buildTree(nums, numArray.n, numArray.tree)
    }
    return numArray
}


func buildTree(nums []int, n int, tree []int) {
    for i := 0; i < n; i++ {
        tree[i+n] = nums[i]
    }
    for i := n-1; i >= 0; i-- {
        tree[i] = tree[i*2] + tree[i*2+1]
    }
}


func (this *NumArray) Update(i int, val int)  {
    pos := i + this.n
    this.tree[pos] = val
    for pos > 0 {
        left, right := pos, pos
        if pos % 2 == 0 {
            right++
        } else {
            left--
        }
        this.tree[pos/2] = this.tree[left] + this.tree[right]
        pos /= 2
    }
}


func (this *NumArray) SumRange(i int, j int) int {
    l := i + this.n
    r := j + this.n
    sum := 0
    for l <= r {
        if l % 2 == 1 {
            sum += this.tree[l]
            l++
        }
        if r % 2 == 0 {
            sum += this.tree[r]
            r--
        }
        l /= 2
        r /= 2
    }
    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
