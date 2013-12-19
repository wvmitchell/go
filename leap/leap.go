package leap

func IsLeapYear(year int) bool {
	switch {
	case divisibleBy(400, year):
		return true
	case divisibleBy(100, year):
		return false
	case divisibleBy(4, year):
		return true
	}
	return false
}

func divisibleBy(divisor int, year int) bool {
	return year%divisor == 0
}
