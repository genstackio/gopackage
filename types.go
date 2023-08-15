package gopackage

type Package struct {
	Files  []File      `json:"files"`
	Target Target_Body `json:"target"`
}

type Target_Body struct {
	Type     string            `json:"type"`
	Location string            `json:"location"`
	Params   map[string]string `json:"params"`
}

type File struct {
	Name    string `json:"name"`
	Source  string `json:"source"`
	Package string `json:"package"`
	Content []byte
	Treated bool
}
