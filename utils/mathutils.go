package utils

func IntMin(n1, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

func IntMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func AbsInt(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
