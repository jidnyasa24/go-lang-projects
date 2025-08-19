package main

// --- Modular condition functions ---
func odd(n int) bool        { return n%2 != 0 }
func even(n int) bool       { return n%2 == 0 }
func greaterThan(k int) func(int) bool {
	return func(n int) bool { return n > k }
}
func lessThan(k int) func(int) bool {
	return func(n int) bool { return n < k }
}
func multipleOf(k int) func(int) bool {
	return func(n int) bool { return n%k == 0 }
}
func prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// --- Orchestration: ANY condition must pass ---
func filterAny(nums []int, conditions ...func(int) bool) []int {
	var result []int
	for _, n := range nums {
		for _, cond := range conditions {
			if cond(n) {
				result = append(result, n)
				break // once one condition passes, keep the number
			}
		}
	}
	return result
}
