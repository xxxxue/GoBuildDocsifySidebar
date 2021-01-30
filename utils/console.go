package utils

import "fmt"

type Console struct {
}

func (Console) ReadLen_string(tips string) string {
	print(tips)
	var res string
	_, _ = fmt.Scanln(&res)
	return res
}

func (Console) ReadLen_int(tips string) int {
	print(tips)
	var res int
	_, _ = fmt.Scanln(&res)
	return res
}
