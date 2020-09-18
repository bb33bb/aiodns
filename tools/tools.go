package tools

import (
	"strconv"
)

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return v
}

func String(s string) string {
	buffer := []byte(s)
	data := make([]byte, len(buffer))
	copy(data, buffer)

	return string(data)
}
