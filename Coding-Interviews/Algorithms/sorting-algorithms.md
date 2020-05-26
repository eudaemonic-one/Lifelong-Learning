# Sorting Algorithms

## Bubble Sort

```go
func BubbleSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
```

### Optimized Bubble Sort

```go
func BubbleSort2(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}
```

### Cocktail Sort

```go
func CocktailSort(arr []int, n int) {
	swapped := true
	start, end := 0, n - 1

	for swapped {
		swapped = false

		for i := start; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false

		end--

		for i := end - 1; i >= start; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		start++
	}
}
```

* **Note that** in practice, cocktail sort performs better than bubble sort although their asymptotic time complexities are the same

## Insertion Sort

```go
func InsertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}
```

### Shell Sort

```go
func ShellSort(arr []int, n int) {
	for gap := n/2; gap > 0; gap /= 2 {
		for i := gap; i < n; i += 1 {
			tmp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > tmp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = tmp
		}
	}
}
```

## Heap Sort

```go
func heapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func HeapSort(arr []int, n int) {
	// Build heap (rearrange array)
	for i := n/2-1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// One by one extract an element from heap
	for i := n-1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}
```

* **Note that** the above algorithm performs heap sort using a max-heap

## Merge Sort

```go
func MergeSort(arr []int, l, r int) {
	if l < r {
		m := l + (r - l) / 2
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m
	L := make([]int, n1)
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	R := make([]int, n2)
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}
	idx, i, j := 0, 0, 0
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[idx] = L[i]
			i++
		} else {
			arr[idx] = R[j]
			j++
		}
		idx++
	}

	for i < n1 {
		arr[idx] = L[i]
		i++
		idx++
	}

	for j < n2 {
		arr[idx] = R[j]
		j++
		idx++
	}
}
```

## Quick Sort

```go
func QuickSort(arr []int, low, high int) {
	if low < high {
    // pivot is partitioning index, arr[pivot] is now at right place
		pivot := partition(arr, low, high)
		QuickSort(arr, low, pivot-1) // before pivot
		QuickSort(arr, pivot+1, high) // after pivot
	}
}

func partition(arr []int, low, high int) int {
  pivot := arr[high] // pivot : Element to be placed at right position
	i := low - 1 // Index of smaller element
	for j := low; j <= high-1; j++ {
    // If current element is smaller than the pivot
		if arr[j] < pivot {
			i++ // increment index of smaller element
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
```

* **Note that** quick sort practices divide and conquer algorithm in a genius way that every number before pivot is less than the pivot and thus the sorting problem can be divided recursively

## Counting Sort

```go
func CountSort(arr []int, m int) {
	res := make([]int, 0)

	count := make([]int, m)

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	for i := 1; i < m; i++ {
		count[i] += count[i-1]
	}

	for i := len(arr)-1; i >= 0; i-- {
		res[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = res[i]
	}
}
```

* **Note that** counting sort can be viewed as a form of dynamic programming, which accumulates counts of numbers and the ranks are thus yielded finally

## Radix Sort

```go
func RadixSort(arr []int, n int) {
	max := 0
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(arr, n, exp)
	}
}

func countSort(arr []int, n int, exp int) {
	res := make([]int, n)
	count := make([]int, Base)

	for i := 0; i < n; i++ {
		count[(arr[i]/exp)%Base]++
	}

	for i := 1; i < Base; i++ {
		count[i] += count[i-1]
	}

	for i := n-1; i >= 0; i-- {
		res[count[(arr[i]/exp)%Base]-1] = arr[i]
		count[(arr[i]/exp)%Base]--
	}

	for i := 0; i < n; i++ {
		arr[i] = res[i]
	}
}
```

* **Note that** radix sort works since counting sort is stable and thus from the least significant digit to most significant bit the numbers are sorted non-decreasingly
* Also **note that** radix sort is useful when the time or space complexity is restricted to linear level

## Bucket Sort

```go
func BucketSort(arr []float64, n int) {
	// 1) Create n empty buckets
	buckets := make([][]float64, SlotNum)
	for i := 0; i < SlotNum; i++ {
		buckets[i] = make([]float64, 0)
	}

	// 2) Put array elements in different buckets
	for i := 0; i < n; i++ {
		idx := int(float64(SlotNum) * arr[i])
		buckets[idx] = append(buckets[idx], arr[i])
	}

	// 3) Sort individual buckets
	for i := 0; i < n; i++ {
		sort.Float64s(buckets[i])
	}

	// 4) Concatenate all buckets into arr[]
	idx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			arr[idx] = buckets[i][j]
			idx++
		}
	}
}
```

* **Note that** before using bucket sort we have to figure out the distribution of numbers given the range and it is efficient if the data are uniformly distributed across the range
* Also **note that** we can take advantage of concurrency techniques to boost the sorting algorithm if the size of each bucket is considerably high
* Think about the distribution of birthdays over months and each month can be viewed as a bucket
