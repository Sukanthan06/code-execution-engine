package models

const (
	StatusSuccess             = "SUCCESS"
	StatusCompilationError    = "COMPILATION_ERROR"
	StatusRuntimeError        = "RUNTIME_ERROR"
	StatusTimeLimitExceeded   = "TIME_LIMIT_EXCEEDED"
	StatusUnsupportedLanguage = "UNSUPPORTED_LANGUAGE"
	StatusInternalError       = "INTERNAL_ERROR"
)

// ExecutionResult represents the outcome of a code execution run.
type ExecutionResult struct {
	Output    string `json:"output"`
	Status    string `json:"status,omitempty"`
	ExitCode  int    `json:"exit_code"`
	RuntimeMS int64  `json:"runtime_ms,omitempty"`
	Error     string `json:"error,omitempty"`
}
