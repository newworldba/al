func minSwaps(data []int) int {
    m := make(map[int]int)
    l := len(data)
    for k:=1;k<=l;k++ {
        m[k] = m[k-1] + data[k-1]
    }
    t := m[l]
    rst := l
    for k:=t;k<l;k++ {
        i := t - (m[k] - m[k-t])
        if rst > i {
            rst = i
        }
    }
    return rst
}
