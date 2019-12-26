package checker

import (
	"fmt"

	"github.com/taxio/lostat/log"
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
		return nil, fmt.Errorf("failed to open: %w", err)
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
	log.Printf("found %d global gitignore patterns\n", len(gp))
	sp, err := gitignore.LoadSystemPatterns(globalFs)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	w.Excludes = append(w.Excludes, sp...)
	log.Printf("found %d system gitignore patterns\n", len(sp))

	status, err := w.Status()
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	log.Printf("%s status:\n%v", c.path, status)
	return !status.IsClean(), nil
}
