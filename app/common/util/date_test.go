package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGetFirstDateOfMonth(t *testing.T) {
	d := time.Now()
	fmt.Println("now: ", d)
	fmt.Println("first day:", GetFirstDateOfMonth(d))
	fmt.Println("last day:", GetLastDateOfMonth(d))
}
