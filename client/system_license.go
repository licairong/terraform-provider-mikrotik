package client

import (
	"log"
)

type SystemLicense struct {
	Account      string `mikrotik:"account"`
	Password     string `mikrotik:"password"`
	Level        string `mikrotik:"level"`
}

func (client Mikrotik) AddSystemLicense(license *SystemLicense) (*SystemLicense, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/system/license/renew", license)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] system license creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	return license, nil
}

func (client Mikrotik) ListSystemLicense() ([]SystemLicense, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/system/license/print"}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] found system license: %v", r)

	license := []SystemLicense{}

	err = Unmarshal(*r, &license)

	if err != nil {
		return nil, err
	}

	return license, nil
}

func (client Mikrotik) FindSystemLicense(id string) (*SystemLicense, error) {
	license := SystemLicense{}
	return &license, nil
}

func (client Mikrotik) UpdateSystemLicense(license *SystemLicense) (*SystemLicense, error) {
	return license, nil
}

func (client Mikrotik) DeleteSystemLicense(id string) error {
	return nil
}
