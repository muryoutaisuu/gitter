package gitter

import (
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Get a new gitter client
func New(url, token string, cs *object.Signature) (*client, error) {
	return &client{
		url:             url,
		token:           token,
		commitsignature: cs,
	}, nil
}

func CreateCommitSignature(name, email string) *object.Signature {
	return &object.Signature{
		Name:  name,
		Email: email,
	}
}

//
func (c *client) CommitNewFile(filename string, content []byte) error {
	return nil
}
