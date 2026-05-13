func productExceptSelf(nums []int) []int {
	zeroIndex, restProduct := -1, 1
	products := make([]int, len(nums))

	for i, n := range nums {
		if n == 0 {
			if zeroIndex != -1 {
				restProduct = 0
			}
			zeroIndex = i
			continue
		}

		restProduct *= n
	}

	if restProduct == 0 {
		return products
	}

	if zeroIndex != -1 {
		products[zeroIndex] = restProduct
		return products
	}

	for i, n := range nums {
		products[i] = restProduct/n
	}
	return products
}
