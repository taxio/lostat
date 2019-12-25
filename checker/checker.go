package checker

import (
	"fmt"

	git "gopkg.in/src-d/go-git.v4"
)

// Checker is a repository status checker
type Checker struct {
	path string
	repo *git.Repository
}

// New returns a Checker object
func New(repoPath string) (*Checker, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &Checker{
		path: repoPath,
		repo: repo,
	}, nil
}

// HasChanges returns whether the repo has changes or not
func (c *Checker) HasChanges() (bool, error) {
	w, err := c.repo.Worktree()
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	status, err := w.Status()
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	return len(status) != 0, nil
}
