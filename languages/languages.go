package languages

type Language struct {
	Name      string
	Extension string

	LocalInterpreter  string
	DockerInterpreter string

	LocalCompiler  string
	DockerCompiler string
}

// language resgistry
var SupportedLanguages = map[string]Language{
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
