package tasks

func romanToInt(s string) int {
	roman_in_int := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	sum := 0
	for i := 0; i <= len(s)-1; i++ {
		if i == len(s)-1 {
			sum += roman_in_int[string(s[i])]
		} else if roman_in_int[string(s[i])] < roman_in_int[string(s[i+1])] {
			sum -= roman_in_int[string(s[i])]
		} else if roman_in_int[string(s[i])] >= roman_in_int[string(s[i+1])] {
			sum += roman_in_int[string(s[i])]
		}
	}
	return sum
}
