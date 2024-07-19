package utils

func IsPrime_1(no int64) bool {
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime_2(no int64) bool {
	for i := int64(2); i <= (no - 1); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimes(start, end int64) []int64 {
	result := make([]int64, 0, 20) //pre-allocating memory
	// var result []int64
	for no := start; no <= end; no++ {
		if IsPrime_1(no) {
			result = append(result, no)
		}
	}
	return result
}
