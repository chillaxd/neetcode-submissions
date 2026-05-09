func hasDuplicate(nums []int) bool {
    entries := map[int]struct{}{}

    for _, n := range nums {
        if _, found := entries[n]; found {
            return true
        }
        entries[n] = struct{}{}
    }

    return false
}
