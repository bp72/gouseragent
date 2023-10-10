package usragnt

import "testing"

func TestUser(t *testing.T) {
	s := "Mozilla/5.0 (Linux; arm_64; Android 13; SM-F731B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.8.36.00 SA/3 Mobile Safari/537.36"
	ua := NewParser().Parse(s)
	if &ua == nil {
		t.Errorf("parser error")
	}
}
