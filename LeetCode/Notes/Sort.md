# Sort

## Comparision Sort

## Distribution Sort

### Radix Sort

```go
func radixSort(nums []int) {
    exp := 1
    radix := 10
    aux := make([]int, len(nums))
    // LSD Radix Sort
    for max/exp > 0{
        // Counting Sort
        count := make([]int, radix)
        for i := 0; i < len(nums); i++ {
            count[(nums[i]/exp) % radix] += 1
        }
        for i := 1; i < radix; i++ {
            count[i] += count[i-1]
        }
        for i := len(nums)-1; i >= 0; i-- {
            count[(nums[i]/exp) % radix] -= 1
            aux[count[(nums[i]/exp) % radix]] = nums[i]
        }
        for i := 0; i < len(nums); i++ {
            nums[i] = aux[i]
        }
        exp *= 10
    }
}
```

e.g. [164. Maximum Gap](https://leetcode.com/problems/maximum-gap/solution/)
