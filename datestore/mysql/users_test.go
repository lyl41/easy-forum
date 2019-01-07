package mysql

import (
	"fmt"
	"testing"
)

func TestFirstOrCreateUser(t *testing.T) {
	data := &Users{
		Phone:  "123",
		Avatar: "123",
		Name:   "123",
	}
	//paramsNeed := []string{"phone", "avatar", "name"}
	err := FirstOrCreateUserByPhone(data.Phone, data)
	fmt.Println(err)
	fmt.Printf("%+v",data)
}
