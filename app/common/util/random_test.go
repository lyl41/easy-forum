package util

import (
	"fmt"
	"testing"
)

func TestRandomStr(t *testing.T) {
	str := RandomStr(10)
	fmt.Println(str)

	numStr := RandomNumberStr(10)
	fmt.Println(numStr)
}
