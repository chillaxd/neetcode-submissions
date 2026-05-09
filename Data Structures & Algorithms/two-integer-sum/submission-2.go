func twoSum(nums []int, target int) []int {
    seen := make(map[int]int, len(nums))

    for i, n := range nums {
        m := target - n
        if j, found := seen[m]; found {
            return []int{j, i}
        }
        seen[n]=i
    }

    return []int{}
}
