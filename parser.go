package usragnt

import (
	"bytes"
	"fmt"
)

type Parser struct {
	t *Trie
}

func (p *Parser) loadTrie() {
	p.t.Add("YaBrowser", NodeTypeBrowser)
	p.t.Add("PhantomJS", NodeTypeBrowser)
	p.t.Add("115Browser", NodeTypeBrowser)
	p.t.Add("Windows 7", NodeTypePlatform)
}

func NewParser() *Parser {
	p := &Parser{t: NewTrie()}
	p.loadTrie()
	return p
}

func (p *Parser) Parse(s string) UserAgent {
	tokens := tonekise(s)
	ua := UserAgent{original: s, Platform: Entity{}, Browser: Entity{}, Device: Entity{}}

	for _, dirtyToken := range tokens {
		token, version := tokenAndVersion(dirtyToken)
		node, err := p.t.Get(token)
		if err != nil {
			continue
		}
		switch node.Type {
		case NodeTypeBrowser:
			ua.Browser.Name = token
			ua.Browser.Version = version
		case NodeTypePlatform:
			ua.Platform.Name = token
			ua.Platform.Version = version
		default:
			fmt.Printf("unknow token %s", dirtyToken)
		}
	}

	return ua
}

func tokenAndVersion(s string) (string, string) {
	for pos, char := range s {
		if char == '/' {
			return s[0:pos], s[pos+1:]
		}
	}
	return s, ""
}

func tonekise(s string) []string {
	res := make([]string, 0)
	var b bytes.Buffer

	startBrace := false

	for _, char := range s {
		switch char {
		case 40:
			startBrace = true
		case 41:
			startBrace = false
		case 32:
			if !startBrace {
				res = append(res, b.String())
				b.Reset()
				startBrace = false
			}
			continue
		}
		b.WriteRune(char)
	}

	return res
}
