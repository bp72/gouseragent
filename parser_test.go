package usragnt

import "testing"

func TestParser(t *testing.T) {
	s := "Mozilla/5.0 (Linux; arm_64; Android 13; SM-F731B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.8.36.00 SA/3 Mobile Safari/537.36"
	p := NewParser()
	ua := p.Parse(s)

	if ua.Browser.Name != "YaBrowser" {
		t.Errorf("expected YaBrowser got: %s", ua.Browser.Name)
	}
}
