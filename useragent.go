package usragnt

type UAType int

const (
	UATypeUnknown UAType = iota
	UABrowser
	UABot
)

type Entity struct {
	Name    string
	Version string
}

type UserAgent struct {
	original string
	Platform Entity
	Browser  Entity
	Device   Entity
	Type     UAType
}
