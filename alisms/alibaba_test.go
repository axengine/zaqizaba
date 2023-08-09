package alisms

import (
	"testing"
)

func TestSendSMS(t *testing.T) {
	c := NewAliSMS("---", "---",
		"MyCompany", "SMS_136075095", `{"code":"%s"}`)
	id, err := c.Send("+8618011111111", "1356")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
