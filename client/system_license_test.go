package client

import (
	"testing"
)

func TestAddSystemLicense(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	account := "account"
	password := "password"
	level := "p1"

	expectedLicense := &SystemLicense{
		Account:    account,
		Password:   password,
		Level:      level,
	}

	_, err := c.AddSystemLicense(expectedLicense)

	if err != nil {
		t.Errorf("Error creating an system license with: %v", err)
	}
}
