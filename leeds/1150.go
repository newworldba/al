func isMajorityElement(nums []int, target int) bool {
    l := len(nums)
    c := l / 2
    t := 0
    for i:=0; i<l; i++ {
        if nums[i] == target {
            t ++
            if t > c {
                return true
            }
        }
    }
    return false
}
