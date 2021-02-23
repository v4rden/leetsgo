package ClimbingStairs

func climbStairs(n int) int {
	previousStair, currentStair := 1, 1

	for i := 2; i <= n; i++ {
		sum := previousStair + currentStair
		previousStair = currentStair
		currentStair = sum
	}

	return currentStair
}
