func compareVersion(version1 string, version2 string) int {
    var revisionNum1, revisionNum2 int
    i, j := 0, 0
    for i < len(version1) && j < len(version2) {
        for k := i; k < len(version1); k++ {
            if k+1 >= len(version1) || version1[k+1] == '.' {
                revisionNum1, _ = strconv.Atoi(version1[i:k+1])
                i = k+2
                break
            }
        }
        for l := j; l < len(version2); l++ {
            if l+1 >= len(version2) || version2[l+1] == '.' {
                revisionNum2, _ = strconv.Atoi(version2[j:l+1])
                j = l+2
                break
            }
        }
        if revisionNum1 < revisionNum2 {
            return -1
        } else if revisionNum1 > revisionNum2 {
            return 1
        }
    }
    if i < len(version1) {
        for i < len(version1) {
            for k := i; k < len(version1); k++ {
                if k+1 >= len(version1) || version1[k+1] == '.' {
                    revisionNum1, _ = strconv.Atoi(version1[i:k+1])
                    i = k+2
                    break
                }
            }
            if revisionNum1 > 0 {
                return 1
            }
        }
    }
    if j < len(version2) {
        for j < len(version2) {
            for l := j; l < len(version2); l++ {
                if l+1 >= len(version2) || version2[l+1] == '.' {
                    revisionNum2, _ = strconv.Atoi(version2[j:l+1])
                    j = l+2
                    break
                }
            }
            if revisionNum2 > 0 {
                return -1
            }
        }
    }
    return 0
}
