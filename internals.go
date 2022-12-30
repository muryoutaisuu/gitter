package gitter

import (
	"os"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"

	billy "github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
)

// Client represents a Gitter Client
type client struct {
	url             string
	token           string
	commitSignature *object.Signature
	storer          *memory.Storage
	fs              *billy.Filesystem
	repo            *git.Repository
	worktree        *git.Worktree
}

func initClient(url, token string, cs *object.Signature) (*client, error) {
	s := memory.NewStorage()
	fs := memfs.New()

	r, err := git.Clone(s, fs, cloneOptions(url, token))
	if err != nil {
		return nil, err
	}

	w, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	return &client{
		url:             url,
		token:           token,
		commitSignature: cs,
		storer:          s,
		fs:              &fs,
		repo:            r,
		worktree:        w,
	}, nil
}

func cloneOptions(url, token string) *git.CloneOptions {
	return &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: "abc", // yes, this can be anything except an empty string
			Password: token,
		},
	}
}

func pushOptions(token string) *git.PushOptions {
	return &git.PushOptions{
		Progress: os.Stdout,
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: "abc", // yes, this can be anything except an empty string
			Password: token,
		},
	}
}

func (c *client) repository() (*git.Repository, error) {
	return git.Clone(c.storer, *c.fs, c.cloneOptions())
}

func (c *client) cloneOptions() *git.CloneOptions {
	return cloneOptions(c.url, c.token)
}

func (c *client) pushOptions() *git.PushOptions {
	return pushOptions(c.token)
}

func (c *client) commit(message string) (plumbing.Hash, error) {
	c.commitSignature.When = time.Now()
	return c.worktree.Commit(message, &git.CommitOptions{
		Author: c.commitSignature,
	})
}

func (c *client) push() error {
	return c.repo.Push(c.pushOptions())
}
