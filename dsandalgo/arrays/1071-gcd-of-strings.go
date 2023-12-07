package arrays

func EuclideanGcd(i, j int) (int, int) {
	smaller := i
	bigger := j
	if j < i {
		smaller, bigger = j, i
	}
	if smaller == 0 {
		return bigger, smaller
	}
	return EuclideanGcd(smaller, bigger%smaller)
}

func GcdOfStrings(str1 string, str2 string) (result string) {
	if str1+str2 == str2+str1 {
		i, _ := EuclideanGcd(len(str1), len(str2))
		return str1[:i]
	}
	return ""
}
