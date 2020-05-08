type Intervals [][]int

func (invls Intervals) Len() int { return len(invls) }
func (invls Intervals) Swap(i, j int){ invls[i], invls[j] = invls[j], invls[i] }
func (invls Intervals) Less(i, j int) bool { return invls[i][0] < invls[j][0] }

func canAttendMeetings(intervals [][]int) bool {
    sorted := Intervals(intervals)
    sort.Sort(sorted)
    for i := 0; i < len(sorted)-1; i++ {
        if sorted[i][1] > sorted[i+1][0] {
            return false
        }
    }
    return true
}
