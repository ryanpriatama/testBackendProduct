package app

import (
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	client := NewRedisClient()

	pong, err := client.Ping().Result()
	if err != nil {
		t.Errorf("NewRedisClient() error = %v", err)
	}

	if pong != "PONG" {
		t.Errorf("NewRedisClient() returned unexpected pong response = %s", pong)
	}

	// Close the connection after testing
	err = client.Close()
	if err != nil {
		t.Errorf("NewRedisClient() failed to close connection: %v", err)
	}
}
