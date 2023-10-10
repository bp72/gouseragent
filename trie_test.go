package usragnt

import "testing"

func TestBasicTrie(t *testing.T) {
	/*
		Mozilla/5.0 (Linux; arm_64; Android 13; SM-F731B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.8.36.00 SA/3 Mobile Safari/537.36
		Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.34 (KHTML, like Gecko) PhantomJS/1.9.0 Safari/534.34 Siteimprove (Accessibility)
		Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.1.0
	*/
	T := NewTrie()
	T.Add("YaBrowser", NodeTypeBrowser)
	T.Add("PhantomJS", NodeTypeBrowser)
	T.Add("115Browser", NodeTypeBrowser)
	T.Add("Windows 7", NodeTypePlatform)

	node, err := T.Get("YaBrowser")
	if err != nil {
		t.Errorf("err: %s", err)
	}

	if node.Type != NodeTypeBrowser {
		t.Errorf("expected NodeTypeBrowser got %d", node.Type)
	}

	node, err = T.Get("Windows")
	if err != nil {
		t.Errorf("err: %s", err)
	}

	if node.Type != NodeTypeNone {
		t.Errorf("expected NodeTypeBrowser got %d", node.Type)
	}
}
