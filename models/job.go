package models

import "time"

// Job status constants representing the lifecycle of an ExecutionJob.
const (
	JobStatusQueued           = "queued"
	JobStatusRunning          = "running"
	JobStatusCompleted        = "completed"
	JobStatusFailed           = "failed"
	JobStatusCompilationError = "compilation_error"
	JobStatusTimeout          = "timeout"
	JobStatusCancelled        = "cancelled"
)

// ExecutionJob represents an application-level execution unit.
type ExecutionJob struct {
	ID          string     `json:"id"`
	Language    string     `json:"language"`
	Code        string     `json:"code"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	RuntimeMS   int64      `json:"runtime_ms,omitempty"`
	ExitCode    int        `json:"exit_code"`
	Output      string     `json:"output"`
	Error       string     `json:"error,omitempty"`
	Cancelled   bool       `json:"cancelled,omitempty"`
}

// NewExecutionJob creates and initializes a new ExecutionJob.
func NewExecutionJob(id, language, code string) *ExecutionJob {
	return &ExecutionJob{
		ID:        id,
		Language:  language,
		Code:      code,
		Status:    JobStatusQueued,
		CreatedAt: time.Now().UTC(),
	}
}
