package prime

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func firstNPrimeSequence(n int) []int {
	sequence := []int{}

	var counter int = 2
	for len(sequence) <= 4 {
		if isPrime(counter) {
			sequence = append(sequence, counter)
		}
		counter++
	}

	return sequence
}
