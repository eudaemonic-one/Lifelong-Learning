type Intervals [][]int

func (invls Intervals) Len() int { return len(invls) }
func (invls Intervals) Less(i, j int) bool { return invls[i][0] < invls[j][0] }
func (invls Intervals) Swap(i, j int){ invls[i], invls[j] = invls[j], invls[i] }

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int){ h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    old = old[0:n-1]
    *h = old
    return x
}

func minMeetingRooms(intervals [][]int) int {
    invlLen := len(intervals)
    if invlLen <= 1 {
        return invlLen
    }
    // sort the intervals
    sorted := Intervals(intervals)
    sort.Sort(sorted)
    // build a min-heap for stording the end time of processing meetings
    rooms := IntHeap{sorted[0][1]}
    heap.Init(&rooms)
    for i := 1; i < invlLen; i++ {
        if sorted[i][0] >= rooms[0] {
            heap.Pop(&rooms)
        }
        heap.Push(&rooms, sorted[i][1])
    }
    return len(rooms)
}
