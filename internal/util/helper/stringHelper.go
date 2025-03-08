package helper

func ClearPhone(phone string) string {
	var newStr []byte
	for i := 0; i < len(phone); i++ {
		if phone[i] == '(' || phone[i] == ')' || phone[i] == '-' {
			continue
		}

		newStr = append(newStr, phone[i])
	}

	return string(newStr)
}
