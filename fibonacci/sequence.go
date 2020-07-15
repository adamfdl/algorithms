package fibonacci

func Sequence(n int) []int {
	switch n {
	case 0:
		return []int{}
	case 1:
		return []int{0}
	case 2:
		return []int{0, 1}
	case 3:
		return []int{0, 1, 1}
	default:
		sequence := []int{0, 1, 1}
		for i := 2; i < n; i++ {
			sequence = append(sequence, sequence[i]+sequence[i-1])
		}
		return sequence
	}

}
