# Models

## Array

### Trapping Rain Water

```python
    def trap(self, height: List[int]) -> int:
        l, r = 0, len(height)-1
        lmax, rmax = 0, 0
        ans = 0
        while l < r:
            if height[l] < height[r]:
                if height[l] > lmax:
                    lmax = height[l]
                else:
                    ans += lmax - height[l]
                    l += 1
            else:
                if height[r] > rmax:
                    rmax = height[r]
                else:
                    ans += rmax - height[r]
                    r -= 1
        return ans
```

### Rotate Matrix

* clockwise rotate
* first reverse up to down, then swap the symmetry

### Array Index Hashing

use the array index as the hash to restore the frequency of each number within the range [1,...,l+1]

```python
nums[nums[i]%n]+=n
```

### Find If On Diagonals

```python
diagonal_lt_rb[x+y]
diagonal_lb_rt[x-y]
```

## Binary Search

### lower_bound

```c++
int lower_bound(int *array, int size, int key)
{
    int first = 0, middle;
    int half, len;
    len = size;

    while(len > 0) {
        half = len >> 1;
        middle = first + half;
        if(array[middle] < key) {
            first = middle + 1;
            len = len-half-1;
        } else
            len = half;
    }
    return first;
}
```

### upper_bound

```c++
int upper_bound(int *array, int size, int key)
{
    int first = 0, len = size-1;
    int half, middle;

    while(len > 0) {
        half = len >> 1;
        middle = first + half;
        if(array[middle] > key)
            len = half;
        else{
            first = middle + 1;
            len = len - half - 1;
        }
    }
    return first;
}
```

## Math

### N Sum (Two Pointers)

```python
def NSum(nums, target, N, temp_set, final_set):
        if len(nums) < N or N < 2 or target < nums[0]*N or target > nums[-1]*N:
            return
        if N == 2:
            l, r = 0, len(nums)-1
            while l < r:
                two_sum = nums[l] + nums[r]
                if two_sum == target:
                    final_set.append(temp_set + [nums[l], nums[r]])
                    l += 1
                    while l < r and nums[l] == nums[l-1]:
                        l += 1
                elif two_sum < target:
                    l += 1
                else:
                    r -= 1
        else:
            for i in range(len(nums)-N+1):
                if i == 0 or nums[i-1] != nums[i]:
                    NSum(nums[i+1:], target-nums[i], N-1, temp_set+[nums[i]], final_set)
```

### Next Permutation

```python
    def nextPermutation(self, nums: List[int]) -> None:
        i = len(nums) - 2
        while i >= 0 and nums[i] >= nums[i+1]:
            i -= 1
        if i >= 0:
            j = len(nums) - 1
            while j >= 0 and nums[j] <= nums[i]:
                j -= 1
            nums[i], nums[j] = nums[j], nums[i]
        start, end = i+1, len(nums)-1
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1
```
