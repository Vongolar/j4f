package jauthority

import (
	"JFFun/data/Dcommand"
	"fmt"
	"testing"
)

func Test_CommandAuthority(t *testing.T) {
	fmt.Println(WithAuthority(Dcommand.Command_guestLogin, Guest))
}

func Test_s2(t *testing.T) {
	a := 6
	a >>= 1
	fmt.Println(a)
}
