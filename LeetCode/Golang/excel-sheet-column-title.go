func convertToTitle(n int) string {
	var res string
	for n > 0 {
		remainder := n % 26
		n /= 26
		if remainder == 0 {
			remainder = 26
			n -= 1
		}
		res = string('A'+remainder-1) + res
	}
	return res
}

