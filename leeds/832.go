func flipAndInvertImage(A [][]int) [][]int {
    for _, v := range A {
        l := len(v)
        for i:=0; i< l/2 ; i++ {
            v[i], v[l - i - 1] = v[l-i-1], v[i]
        }
        for kk, vv := range v {
            if vv == 1 {
                v[kk] = 0
            } else {
                v[kk] = 1
            }
        }
    }
    return A
}
