package maxsum

func MaxSumSubArray(k int, arr []int) int {
	windowSum := int(0)
	maxSum := int(0)
	windowStart := 0
	for windowEnd := range arr {
		windowSum += arr[windowEnd]
		if windowEnd >= k+1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= arr[windowStart]
			windowStart = windowStart + 1
		}
	}
	return maxSum
}
