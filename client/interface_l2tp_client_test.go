package client

import (
	"reflect"
	"testing"
)

func TestAddL2tpClientAndDeleteL2tpClient(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	name := "l2tp-out3"
	password := "password"
	connect_to := "192.168.1.1"
	ipsec_secret := "password"
	use_ipsec := "yes"
	user := "admin"
	comment := "terraform acc test"
	updatedComment := "terraform acc test update"

	expectedL2tpClient := &L2tpClient{
		Name:       name,
		Password:   password,
		ConnectTo:    connect_to,
		IpsecSecret:     ipsec_secret,
		UseIpsec:    use_ipsec,
		User: user,
		Comment:    comment,
	}

	secret, err := c.AddL2tpClient(expectedL2tpClient)

	if err != nil {
		t.Errorf("Error creating an L2tp Client with: %v", err)
	}

	expectedL2tpClient.Id = secret.Id

	if !reflect.DeepEqual(secret, expectedL2tpClient) {
		t.Errorf("The L2tp Client does not match what we expected. actual: %v expected: %v", secret, expectedL2tpClient)
	}

	expectedL2tpClient.Comment = updatedComment
	secret, err = c.UpdateL2tpClient(expectedL2tpClient)

	if err != nil {
		t.Errorf("Error updating an L2tp Client with: %v", err)
	}
	if !reflect.DeepEqual(secret, expectedL2tpClient) {
		t.Errorf("The L2tp Client does not match what we expected. actual: %v expected: %v", secret, expectedL2tpClient)
	}

	foundL2tpClient, err := c.FindL2tpClient(secret.Id)

	if err != nil {
		t.Errorf("Error getting L2tp Client with: %v", err)
	}

	if !reflect.DeepEqual(secret, foundL2tpClient) {
		t.Errorf("Created L2tp Client and found L2tp Client do not match. actual: %v expected: %v", foundL2tpClient, secret)
	}

	err = c.DeleteL2tpClient(secret.Id)

	if err != nil {
		t.Errorf("Error deleting L2tp Client with: %v", err)
	}
}
