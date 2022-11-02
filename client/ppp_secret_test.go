package client

import (
	"reflect"
	"testing"
)

func TestAddPppSecretAndDeletePppSecret(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	name := "ydd_test"
	password := "Passw0rd@_"
	profile := "l2tp-profile"
	routes := "10.100.80.0/24"
	service := "l2tp"
	comment := "terraform acc test"
	updatedComment := "terraform acc test update"

	expectedPppSecret := &PppSecret{
		Name:       name,
		Password:   password,
		Profile:    profile,
		Routes:     routes,
		Service:    service,
		Comment:    comment,
	}

	secret, err := c.AddPppSecret(expectedPppSecret)

	if err != nil {
		t.Errorf("Error creating an ppp secret with: %v", err)
	}

	expectedPppSecret.Id = secret.Id

	if !reflect.DeepEqual(secret, expectedPppSecret) {
		t.Errorf("The ip route does not match what we expected. actual: %v expected: %v", secret, expectedPppSecret)
	}

	expectedPppSecret.Comment = updatedComment
	secret, err = c.UpdatePppSecret(expectedPppSecret)

	if err != nil {
		t.Errorf("Error updating an ppp secret with: %v", err)
	}
	if !reflect.DeepEqual(secret, expectedPppSecret) {
		t.Errorf("The ppp secret does not match what we expected. actual: %v expected: %v", secret, expectedPppSecret)
	}

	foundPppSecret, err := c.FindPppSecret(secret.Id)

	if err != nil {
		t.Errorf("Error getting ppp secret with: %v", err)
	}

	if !reflect.DeepEqual(secret, foundPppSecret) {
		t.Errorf("Created ppp secret and found ppp secret do not match. actual: %v expected: %v", foundPppSecret, secret)
	}

	err = c.DeletePppSecret(secret.Id)

	if err != nil {
		t.Errorf("Error deleting ppp secret with: %v", err)
	}
}
