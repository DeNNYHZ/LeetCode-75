package Stack

func decodeString(s string) string {
	stack := []string{}
	numStack := []int{}
	num := 0
	currentString := ""

	for _, char := range s {
		if char >= '0' && char <= '9' {
			// Membentuk angka jika ada lebih dari 1 digit
			num = num*10 + int(char-'0')
		} else if char == '[' {
			// Push angka dan currentString ke stack
			numStack = append(numStack, num)
			stack = append(stack, currentString)
			// Reset currentString dan num
			currentString = ""
			num = 0
		} else if char == ']' {
			// Ambil angka terakhir dan ulangi currentString
			repeatCount := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Ulangi currentString dan gabungkan ke temp
			for i := 0; i < repeatCount; i++ {
				temp += currentString
			}
			currentString = temp
		} else {
			// Bangun currentString
			currentString += string(char)
		}
	}

	return currentString
}
