package helpers

import "github.com/solo-io/glue/pkg/module"

func NewTestSecrets() module.SecretMap {
	return map[string]map[string][]byte{
		"user": {"username": []byte("me@example.com"), "password": []byte("foobar")},
	}
}
