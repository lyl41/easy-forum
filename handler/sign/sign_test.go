package sign

import (
	"fmt"
	"testing"
	"time"
)

func TestDealSign(t *testing.T) {
	//fmt.Println(time.Now().Format(util.StdTimeFormat))
	d := time.Now().AddDate(0, 1, 25)
	ret, err := AddMask(0, d)
	fmt.Println(ret, err)
	fmt.Println(1 << uint32(4))
	for i := 0; i < 32; i++ {
		fmt.Print(ret & (1 << uint8(i)))
		fmt.Print(" ")
	}
	fmt.Println()
}
