package main

// 生成markdown中的层级 空格
func GenerateSpace(n int) string {
	var res = ""

	for true {

		if n > 0 {
			res += "  "
		} else {
			break
		}
		n--
	}

	return res
}

