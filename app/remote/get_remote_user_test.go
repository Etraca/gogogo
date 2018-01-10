package remote

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	GetUser()
}

func TestPostUser(t *testing.T) {
	PostUser("{\"userName\":\"Kobe Bryant\",\"logonName\":\"kobe\",\"passWord\":\"780823\"}")
}