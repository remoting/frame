package email

import (
	"testing"
)

func Test_Example1(t *testing.T) {
	OnInit("", "", "", "", "", 465)
	to := &Address{
		Address: "",
		Name:    "",
	}
	err := SendMail(to, "", "body")
	println("%v", err)
}
