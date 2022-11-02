package client

import (
	"fmt"
	"log"
)

type IpRoute struct {
	Id           string `mikrotik:".id"`
	DstAddress   string `mikrotik:"dst-address"`
	Comment      string `mikrotik:"comment"`
	Disabled     bool   `mikrotik:"disabled"`
	Gateway      string `mikrotik:"gateway"`
}

func (client Mikrotik) AddIpRoute(route *IpRoute) (*IpRoute, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/ip/route/add", route)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip route creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	id := r.Done.Map["ret"]

	return client.FindIpRoute(id)
}

func (client Mikrotik) ListIpRoute() ([]IpRoute, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}
	cmd := []string{"/ip/route/print"}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] found ip route: %v", r)

	iproute := []IpRoute{}

	err = Unmarshal(*r, &iproute)

	if err != nil {
		return nil, err
	}

	return iproute, nil
}

func (client Mikrotik) FindIpRoute(id string) (*IpRoute, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{"/ip/route/print", "?.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip route response: %v", r)

	if err != nil {
		return nil, err
	}

	iproute := IpRoute{}
	err = Unmarshal(*r, &iproute)

	if err != nil {
		return nil, err
	}

	if iproute.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("ip route `%s` not found", id))
	}

	return &iproute, nil
}

func (client Mikrotik) UpdateIpRoute(route *IpRoute) (*IpRoute, error) {
	c, err := client.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := Marshal("/ip/route/set", route)
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)

	if err != nil {
		return nil, err
	}

	return client.FindIpRoute(route.Id)
}

func (client Mikrotik) DeleteIpRoute(id string) error {
	c, err := client.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{"/ip/route/remove", "=.id=" + id}
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	_, err = c.RunArgs(cmd)
	return err
}
