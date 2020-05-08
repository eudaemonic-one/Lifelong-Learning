func canCompleteCircuit(gas []int, cost []int) int {
    var tank int
    var flag bool
    for i := 0; i < len(gas); i++ {
        tank = 0
        j := i
        flag = true
        for ; j < len(gas); j++ {
            tank += gas[j]
            tank -= cost[j]
            if (tank < 0) {
                flag = false
                break
            }
        }
        if (!flag) {
            continue
        }
        for j = 0; j < i; j++ {
            tank += gas[j]
            tank -= cost[j]
            if (tank < 0) {
                flag = false
                break
            }
        }
        if (flag) {
            return i
        }
    }
    return -1
}
