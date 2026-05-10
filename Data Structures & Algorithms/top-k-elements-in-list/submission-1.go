func topKFrequent(nums []int, k int) []int {
	repeatedElements := make(map[int]int, len(nums))
	repeatCounts := []int{}
	result := []int{}

	for _, n := range nums {
		repeatedElements[n]++
	}
	if len(repeatedElements) == 1 {
		return []int{nums[0]}
	}

	for _k := range repeatedElements {
		repeatCounts = append(repeatCounts, _k)
	}

	sort.Slice(repeatCounts, func(i,j int) bool{
		return repeatedElements[repeatCounts[i]] > repeatedElements[repeatCounts[j]]
	})

	for _, e := range repeatCounts {
		result = append(result, e)
		k--
		if k == 0 {
			break
		}
	}

	return result
}
