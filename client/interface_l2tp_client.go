package client

import (
	"fmt"
	"log"
)

type L2tpClient struct {
	Id                    string `mikrotik:".id"`
	Name                  string `mikrotik:"name"`
	Password              string `mikrotik:"password"`
	ConnectTo             string `mikrotik:"connect-to"`
	IpsecSecret           string `mikrotik:"ipsec-secret"`
	KeepaliveTimeout      string `mikrotik:"keepalive-timeout"`
	UseIpsec              string `mikrotik:"use-ipsec"`
	User                  string `mikrotik:"user"`
	Comment               string `mikrotik:"comment"`
	Disabled              bool   `mikrotik:"disabled"`
}

func (client Mikrotik) AddL2tpClient(l2tpClient *L2tpClient) (*L2tpClient, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/l2tp-client/add", l2tpClient)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] l2tp-client creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	id := r.Done.Map["ret"]

	return client.FindL2tpClient(id)
}

func (client Mikrotik) ListL2tpClient() ([]L2tpClient, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/interface/l2tp-client/print"}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] found interface l2tp-client: %v", r)

	l2tpClient := []L2tpClient{}

	err = Unmarshal(*r, &l2tpClient)

	if err != nil {
		return nil, err
	}

	return l2tpClient, nil
}

func (client Mikrotik) FindL2tpClient(id string) (*L2tpClient, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/interface/l2tp-client/print", "?.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] interface l2tp-client response: %v", r)

	if err != nil {
		return nil, err
	}

	l2tpClient := L2tpClient{}
	err = Unmarshal(*r, &l2tpClient)

	if err != nil {
		return nil, err
	}

	if l2tpClient.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("interface l2tp-client `%s` not found", id))
	}

	return &l2tpClient, nil
}

func (client Mikrotik) UpdateL2tpClient(l2tpClient *L2tpClient) (*L2tpClient, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/interface/l2tp-client/set", l2tpClient)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return client.FindL2tpClient(l2tpClient.Id)
}

func (client Mikrotik) DeleteL2tpClient(id string) error {
	c, err := client.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{"/interface/l2tp-client/remove", "=.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)
	return err
}
