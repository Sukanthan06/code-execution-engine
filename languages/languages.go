package languages

type Language struct {
	Name      string
	Extension string
	Command   string
}

// language resgistry
var SupportedLanguages = map[string]Language{
	"python": {
		Name:      "Python",
		Extension: ".py",
		Command:   "python",
	},
	"javascript": {
		Name:      "JavaScript",
		Extension: ".js",
		Command:   "node",
	},
}
