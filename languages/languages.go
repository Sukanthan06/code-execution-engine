package languages

type Language struct {
	Name      string
	Extension string

	LocalInterpreter  string
	DockerInterpreter string

	LocalCompiler  string
	DockerCompiler string
}

// language registry
var supportedLanguages = map[string]Language{
	"python": {
		Name:              "Python",
		Extension:         ".py",
		LocalInterpreter:  "python",
		DockerInterpreter: "python3",
	},
	"javascript": {
		Name:              "JavaScript",
		Extension:         ".js",
		LocalInterpreter:  "node",
		DockerInterpreter: "node",
	},
	"c": {
		Name:           "C",
		Extension:      ".c",
		LocalCompiler:  "gcc",
		DockerCompiler: "gcc",
	},
	"cpp": {
		Name:           "C++",
		Extension:      ".cpp",
		LocalCompiler:  "g++",
		DockerCompiler: "g++",
	},
}

// GetLanguage returns the Language configuration for the given language name.
func GetLanguage(name string) (Language, bool) {
	lang, exists := supportedLanguages[name]
	return lang, exists
}

// IsSupported checks if the given language name is registered.
func IsSupported(name string) bool {
	_, exists := supportedLanguages[name]
	return exists
}
