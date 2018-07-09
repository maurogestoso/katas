package katas

// ConfirmEnding returns a boolean of wether str ends with end
func ConfirmEnding(str, end string) bool {
	start := len(str) - len(end)
	if len(end) == 0 || start < 0 {
		return false
	}
	offset := len(end) + 1
	for i := start; i < len(str); i++ {
		if str[i] != end[i-offset] {
			return false
		}
	}
	return true
}
