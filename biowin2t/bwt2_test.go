package biowin2t

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	url := "http://localhost/api/1.0/lookup"
	username := "usr"
	password := "passwd"
	c := NewWindhagerClient(url,
		username, password)
	if c.MesEndpoint != url {
		t.Errorf("NewWindhagerClient MesEndpoint = %v, want %v", c.MesEndpoint, url)
	}
	if c.MesUsername != username {
		t.Errorf("NewWindhagerClient MesUsername = %v, want %v", c.MesUsername, username)
	}
	if c.MesPassword != password {
		t.Errorf("NewWindhagerClient MesPassword = %v, want %v", c.MesPassword, password)
	}
}
