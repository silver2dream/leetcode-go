package s1720

func decode(encoded []int, first int) []int {
	n := len(encoded)
	original := make([]int, n+1)
	original[0] = first
	for i := 1; i < len(original); i++ {
		original[i] = encoded[i-1] ^ original[i-1]
	}
	return original
}
