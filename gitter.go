package gitter

import (
	"github.com/go-git/go-billy/v5/util"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Get a new gitter client
func New(url, token string, cs *object.Signature) (*client, error) {
	return initClient(url, token, cs)
}

// Helper function to create correctly structured Signatures
func CreateCommitSignature(name, email string) *object.Signature {
	return &object.Signature{
		Name:  name,
		Email: email,
	}
}

// Add, commit and push new file to repository
func (c *client) CommitNewFile(filename, message string, content []byte) error {
	// write the file, and return if got an error
	err := util.WriteFile(*c.fs, filename, content, 0440)
	if err != nil {
		return err
	}

	// add file to worktree
	_, err = c.worktree.Add(filename)
	if err != nil {
		return err
	}

	// commit stuff
	_, err = c.commit(message)
	if err != nil {
		return err
	}

	// push stuff
	err = c.push()
	if err != nil {
		return err
	}

	return nil
}
