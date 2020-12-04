package http_service

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFindUser(t *testing.T) {
	username := "root"
	user, err := FindUser(username)
	if err != nil {
		t.Error(err.Error())
	}
	println(user.Username.String)
	assert.NotEqual(t, user.Id, 0)
}
