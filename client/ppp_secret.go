package client

import (
	"fmt"
	"log"
)

type PppSecret struct {
	Id           string `mikrotik:".id"`
	Name         string `mikrotik:"name"`
	Password     string `mikrotik:"password"`
	Profile      string `mikrotik:"profile"`
	Routes       string `mikrotik:"routes"`
	Service      string `mikrotik:"service"`
	Comment      string `mikrotik:"comment"`
	Disabled     bool   `mikrotik:"disabled"`
}

func (client Mikrotik) AddPppSecret(secret *PppSecret) (*PppSecret, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/ppp/secret/add", secret)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ppp secret creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	id := r.Done.Map["ret"]

	return client.FindPppSecret(id)
}

func (client Mikrotik) ListPppSecret() ([]PppSecret, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/ppp/secret/print"}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] found ppp secret: %v", r)

	secret := []PppSecret{}

	err = Unmarshal(*r, &secret)

	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (client Mikrotik) FindPppSecret(id string) (*PppSecret, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/ppp/secret/print", "?.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ppp secret response: %v", r)

	if err != nil {
		return nil, err
	}

	secret := PppSecret{}
	err = Unmarshal(*r, &secret)

	if err != nil {
		return nil, err
	}

	if secret.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("ppp secret `%s` not found", id))
	}

	return &secret, nil
}

func (client Mikrotik) UpdatePppSecret(secret *PppSecret) (*PppSecret, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/ppp/secret/set", secret)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return client.FindPppSecret(secret.Id)
}

func (client Mikrotik) DeletePppSecret(id string) error {
	c, err := client.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{"/ppp/secret/remove", "=.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)
	return err
}
