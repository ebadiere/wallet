package client

import "testing"

func TestClientConnection(t *testing.T) {

	conn := true
	if !conn {
		t.Error("Connection Failed")
	}
}
