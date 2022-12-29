package gitter

import (
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"

	billy "github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
)

// Client represents a Gitter Client
type client struct {
	url             string
	token           string
	commitsignature *object.Signature
	storer          *memory.Storage
	fs              *billy.Filesystem
}

func initClient(url, token string, cs *object.Signature) *client {
	s := memory.NewStorage()
	fs := memfs.New()
	return &client{
		url:             url,
		token:           token,
		commitsignature: cs,
		storer:          s,
		fs:              &fs,
	}
}

func (c *client) repository() (*git.Repository, error) {
	return git.Clone(c.storer, *c.fs, c.cloneOptions())
}

func (c *client) cloneOptions() *git.CloneOptions {
	return &git.CloneOptions{
		URL:      c.url,
		Progress: os.Stdout,
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: "abc", // yes, this can be anything except an empty string
			Password: c.token,
		},
	}
}

func (c *client) writeFile(filename string, data []byte, perm os.FileMode) error {
	return util.WriteFile(*c.fs, filename, data, perm)
}
