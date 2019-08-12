func consecutiveNumbersSum(N int) int {
    var count int
    for k := 1; (k - 1) * k / 2 < N;k++ {
        tmp := N - (k - 1) * k / 2
        if tmp % k == 0 {
            count ++
        }
    }
    return count
}
/**
N = (a + a + k - 1) * k / 2
N = ak + (k - 1) * k / 2
N - ak = (k - 1) * k / 2
*/
