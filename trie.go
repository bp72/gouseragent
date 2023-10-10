package usragnt

import "fmt"

type NodeType int

const (
	NodeTypeNone NodeType = iota
	NodeTypeBrowser
	NodeTypePlatform
)

type Trie struct {
	Nodes map[rune]*Trie
	Type  NodeType
}

func NewTrie() *Trie {
	return &Trie{Nodes: make(map[rune]*Trie)}
}

func (t *Trie) Add(s string, tp NodeType) {
	head := t

	for _, char := range s {
		if head.Nodes[char] == nil {
			head.Nodes[char] = NewTrie()
		}
		head = head.Nodes[char]
		head.Type = NodeTypeNone
	}

	head.Type = tp
}

func (t *Trie) Get(s string) (*Trie, error) {
	head := t

	for _, char := range s {
		if head.Nodes[char] == nil {
			return nil, fmt.Errorf("not found %s of %s", string(char), s)
		}
		head = head.Nodes[char]
	}

	return head, nil
}
