package checker

import (
	"fmt"

	"gopkg.in/src-d/go-billy.v4/osfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/format/gitignore"
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

	// WORKAROUND: handle global gitignore patterns.
	// ref: https://github.com/src-d/go-git/issues/760#issuecomment-523189734
	globalFs := osfs.New("/")
	gp, err := gitignore.LoadGlobalPatterns(globalFs)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	w.Excludes = append(w.Excludes, gp...)
	sp, err := gitignore.LoadSystemPatterns(globalFs)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	w.Excludes = append(w.Excludes, sp...)

	status, err := w.Status()
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	return !status.IsClean(), nil
}
