package test

var (
	_ = `go`              // want "enclosed with double quotes \"go\""
	_ = "go\\ngo"         // want "enclosed with back quotes `go\\\\ngo`"
	_ = "{\"go\":\"go\"}" // want "enclosed with back quotes `{\"go\":\"go\"}`"
	_ = "go\ngo"
	_ = "go\tgo"
	_ = "go\rgo"
	_ = "go\bgo"
	_ = "go\fgo"
	_ = "go\vgo"
	_ = "go\ago"
	_ = "go\\\ngo"
	_ = `go
	go`
)

type T struct {
	Tag string `json:"tag"`
}
