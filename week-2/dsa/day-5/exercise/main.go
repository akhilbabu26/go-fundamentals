package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	deque := []int{} 
	result := []int{}

	for i := 0; i < len(nums); i++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if deque[0] <= i-k {
			deque = deque[1:]
		}

		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	result := maxSlidingWindow(nums, k)
	fmt.Println("Sliding Window Maximum:", result)
}
