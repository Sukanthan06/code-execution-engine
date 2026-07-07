package languages

type Language struct {
	Name        string
	Extension   string
	Interpreter string
	Compiler    string
}

// language resgistry
var SupportedLanguages = map[string]Language{
	"python": {
		Name:        "Python",
		Extension:   ".py",
		Interpreter: "python",
	},
	"javascript": {
		Name:        "JavaScript",
		Extension:   ".js",
		Interpreter: "node",
	},
	"c": {
		Name:      "C",
		Extension: ".c",
		Compiler:  "gcc",
	},
}
