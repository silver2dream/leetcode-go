package s342

func isPowerOfFour(n int) bool {
	if n <= 0 || n&(n-1) != 0 {
		return false
	}
	return (n & 0x55555555) != 0
}
