package test

func mergeSort(nums []int, start int, end int) {
	if start == end {
		return
	}

	mid := end / 2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)

	temps := []int{}
	left := 0
	leftTemps := nums[:mid]

	right := mid + 1
	rightTemps := nums[mid+1:]

	for {
		if left > mid {
			temps = append(temps, rightTemps[right:]...)
			break
		} else if right > end {
			temps = append(temps, rightTemps[left:]...)
			break
		}

		if leftTemps[left] <= rightTemps[right] {
			temps = append(temps, leftTemps[left])
			left++
		} else if leftTemps[left] > rightTemps[right] {
			temps = append(temps, rightTemps[right])
			right++
		}
	}

	for i := 0; i < len(temps); i++ {
		nums[i] = temps[i]
	}
}
