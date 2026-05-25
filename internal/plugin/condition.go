package plugin

import (
	"fmt"
	"os"
	"strings"
)

type Condition struct {
	env func(string) string
}

func New() *Condition { return &Condition{env: os.Getenv} }

func NewWithEnv(env func(string) string) *Condition { return &Condition{env: env} }

func (c *Condition) Check() error {
	var errs []string

	if c.env("GITLAB_CI") != "true" {
		errs = append(errs, "GITLAB_CI is not set to \"true\"; this plugin requires a GitLab CI environment")
	}

	if c.env("CI_JOB_TOKEN") == "" && c.env("GITLAB_TOKEN") == "" {
		errs = append(errs, "neither CI_JOB_TOKEN nor GITLAB_TOKEN is set")
	}

	if branch := c.env("SEMREL_PLUGIN_BRANCH"); branch != "" {
		gotBranch := c.env("CI_COMMIT_REF_NAME")
		if gotBranch == "" {
			gotBranch = c.env("CI_COMMIT_BRANCH")
		}
		if gotBranch != branch {
			errs = append(errs, fmt.Sprintf("branch mismatch: want %q got %q", branch, gotBranch))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, "; "))
	}
	return nil
}
