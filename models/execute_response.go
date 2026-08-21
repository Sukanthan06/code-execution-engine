package models

type ExecuteResponse struct {
	Output    string `json:"output"`
	Status    string `json:"status,omitempty"`
	ExitCode  int    `json:"exit_code"`
	RuntimeMS int64  `json:"runtime_ms,omitempty"`
	Error     string `json:"error,omitempty"`
}
