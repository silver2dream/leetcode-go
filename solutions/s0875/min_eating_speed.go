package test

func minEatingSpeed(piles []int, h int) int {
	if len(piles) < 1 {
		return 0
	}

	max := getMax(piles)
	speed := getSpeed(piles, 1, max, h)
	return speed
}

func getSpeed(piles []int, start int, end int, h int) int {
	speed := -1
	pivot := (start + end) / 2
	cost := getCost(piles, pivot)
	if cost <= h {
		speed = pivot
		if start < pivot {
			newSpeed := getSpeed(piles, start, pivot-1, h)
			if newSpeed != -1 {
				speed = newSpeed
			}
		}
	} else if cost > h {
		if pivot < end {
			speed = getSpeed(piles, pivot+1, end, h)
		}
	}

	return speed
}

func getMax(piles []int) int {
	max := 0
	for i := range piles {
		if piles[i] > max {
			max = piles[i]
		}
	}
	return max
}

func getCost(piles []int, pivot int) int {
	cost := 0
	for _, banana := range piles {
		cost += banana / pivot
		if banana%pivot != 0 {
			cost++
		}
	}
	return cost
}
