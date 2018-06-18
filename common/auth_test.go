package common

import (
	"testing"
	"fmt"
)

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT("raspMachine","AC:25:D5:E8:F9:AA","admin")
	fmt.Println(token)
	if err!= nil {
		t.Error("Expected no error, got ",err)
	}
}
