package models

// ExecutionResult represents the outcome of a code execution run.
type ExecutionResult struct {
	Output    string `json:"output"`
	Status    string `json:"status,omitempty"`
	ExitCode  int    `json:"exit_code,omitempty"`
	RuntimeMS int64  `json:"runtime_ms,omitempty"`
	Error     string `json:"error,omitempty"`
}
