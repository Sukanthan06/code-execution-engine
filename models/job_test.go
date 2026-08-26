package models

import (
	"testing"
)

func TestNewExecutionJob(t *testing.T) {
	id := "test-uuid-123"
	lang := "python"
	code := "print('hello')"

	job := NewExecutionJob(id, lang, code)

	if job.ID != id {
		t.Errorf("expected ID %q, got %q", id, job.ID)
	}
	if job.Language != lang {
		t.Errorf("expected Language %q, got %q", lang, job.Language)
	}
	if job.Code != code {
		t.Errorf("expected Code %q, got %q", code, job.Code)
	}
	if job.Status != JobStatusQueued {
		t.Errorf("expected Status %q, got %q", JobStatusQueued, job.Status)
	}
	if job.CreatedAt.IsZero() {
		t.Errorf("expected CreatedAt timestamp to be initialized")
	}
	if job.StartedAt != nil {
		t.Errorf("expected StartedAt to be nil initially")
	}
	if job.CompletedAt != nil {
		t.Errorf("expected CompletedAt to be nil initially")
	}
}

func TestJobStatusConstants(t *testing.T) {
	statuses := []string{
		JobStatusQueued,
		JobStatusRunning,
		JobStatusCompleted,
		JobStatusFailed,
		JobStatusCompilationError,
		JobStatusTimeout,
		JobStatusCancelled,
	}

	for _, s := range statuses {
		if s == "" {
			t.Errorf("expected non-empty status constant")
		}
	}
}
