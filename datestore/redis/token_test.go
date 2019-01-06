package redis

import (
	"fmt"
	"testing"
)

func TestGetTokenValue(t *testing.T) {
	sess := &SessionInfo{
		UserId: 1,
	}
	err := SetTokenValue("lyl_token", sess)
	fmt.Println("err1:", err)

	info, _ := GetTokenValue("lyl_token")
	fmt.Println(info)
}
