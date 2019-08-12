func numRollsToTarget(d int, f int, target int) int {
    c := 0
    f64 := int64(f)
    d64 := int64(d)
    for i:=int64(0);i<=pow(f64, d64);i++{
        v := i
        t := int64(0)
        for j:=int64(d-1);j>=0;j++{
            p := pow(f64, j)
            fmt.Println(f, j, p)
            t += v / p
            v = v % p
        }    
        if t == int64(target) {
            c ++
        }
    }
    return c
}
func pow(a int64, b int64) int64 {
    if b == 0 {
        return 1
    }
    c := int64(1)
    for i:=int64(0); i< b;i++{
        c = c*a
    }
    return c
}
