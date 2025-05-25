package main

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	stat := 0
	maxN := 0
	for i := range grumpy {
		if grumpy[i] == 1 {
			stat += customers[i]
		}
	}

	for i := range grumpy {
		if grumpy[i] == 0 {
			stat += customers[i]
		}

		if i < minutes-1 {
			continue
		}

		maxN = max(maxN, stat)

		if grumpy[i-minutes+1] == 0 {
			stat -= customers[i-minutes+1]
		}
	}

	return maxN
}
