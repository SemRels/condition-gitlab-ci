package plugin

import (
	"strings"
	"testing"
)

func env(kv map[string]string) func(string) string {
	return func(key string) string { return kv[key] }
}

func TestCheck_HappyPath(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITLAB_CI":    "true",
		"CI_JOB_TOKEN": "token",
	})).Check()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_AllowsGitLabToken(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITLAB_CI":    "true",
		"GITLAB_TOKEN": "token",
	})).Check()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_RequiresGitLabCI(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{"CI_JOB_TOKEN": "token"})).Check()
	if err == nil || !strings.Contains(err.Error(), "GITLAB_CI") {
		t.Fatalf("expected gitlab ci error, got: %v", err)
	}
}

func TestCheck_RequiresToken(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{"GITLAB_CI": "true"})).Check()
	if err == nil || !strings.Contains(err.Error(), "CI_JOB_TOKEN") {
		t.Fatalf("expected token error, got: %v", err)
	}
}

func TestCheck_BranchMatch(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITLAB_CI":            "true",
		"CI_JOB_TOKEN":         "token",
		"SEMREL_PLUGIN_BRANCH": "main",
		"CI_COMMIT_REF_NAME":   "main",
	})).Check()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_BranchFallback(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITLAB_CI":            "true",
		"CI_JOB_TOKEN":         "token",
		"SEMREL_PLUGIN_BRANCH": "main",
		"CI_COMMIT_BRANCH":     "main",
	})).Check()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_BranchMismatch(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"GITLAB_CI":            "true",
		"CI_JOB_TOKEN":         "token",
		"SEMREL_PLUGIN_BRANCH": "main",
		"CI_COMMIT_REF_NAME":   "release",
	})).Check()
	if err == nil || !strings.Contains(err.Error(), "branch mismatch") {
		t.Fatalf("expected branch mismatch, got: %v", err)
	}
}

func TestCheck_MultipleErrors(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{})).Check()
	if err == nil || !strings.Contains(err.Error(), "GITLAB_CI") || !strings.Contains(err.Error(), "CI_JOB_TOKEN") {
		t.Fatalf("expected combined errors, got: %v", err)
	}
}
