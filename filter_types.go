package proton

type Filter struct {
	ID       string
	Name     string
	Status   FilterStatus
	Priority int
	Version  int
	Simple   string
	Sieve    string
}

type FilterUpdateReq struct {
	Name   string
	Status FilterStatus
	Simple string
	Sieve  string
}

func (f Filter) IsEnabled() bool {
	return f.Status == FilterEnabled
}

type FilterStatus int

const (
	FilterDisabled = iota
	FilterEnabled
)

type FilterPos struct {
	Line   int
	Ch     int
	Sticky bool
}

type FilterIssue struct {
	Message  string
	Severity string
	From     FilterPos
	To       FilterPos
}
