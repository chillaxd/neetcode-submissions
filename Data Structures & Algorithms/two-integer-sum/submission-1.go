func twoSum(nums []int, target int) []int {
    if len(nums) == 2 {
        if nums[0] + nums[1] == target {
            return []int{0,1}
        }
        
        return []int{}
    }

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
