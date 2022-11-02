package client

import (
	"reflect"
	"testing"
)

func TestAddIpRouteAndDeleteIpRoute(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	dstAddress := "10.1.2.0/24"
	gateway := "l2tp-out2"
	comment := "terraform-acc-test"
	disabled := false
	updatedComment := "terraform acc test updated"

	expectedIpRoute := &IpRoute{
		DstAddress:   dstAddress,
		Comment:   comment,
		Disabled:  disabled,
		Gateway: gateway,
	}

	iproute, err := c.AddIpRoute(expectedIpRoute)

	if err != nil {
		t.Errorf("Error creating an ip route with: %v", err)
	}

	expectedIpRoute.Id = iproute.Id

	if !reflect.DeepEqual(iproute, expectedIpRoute) {
		t.Errorf("The ip route does not match what we expected. actual: %v expected: %v", iproute, expectedIpRoute)
	}

	expectedIpRoute.Comment = updatedComment
	iproute, err = c.UpdateIpRoute(expectedIpRoute)

	if err != nil {
		t.Errorf("Error updating an ip route with: %v", err)
	}
	if !reflect.DeepEqual(iproute, expectedIpRoute) {
		t.Errorf("The ip route does not match what we expected. actual: %v expected: %v", iproute, expectedIpRoute)
	}

	foundIpRoute, err := c.FindIpRoute(iproute.Id)

	if err != nil {
		t.Errorf("Error getting ip route with: %v", err)
	}

	if !reflect.DeepEqual(iproute, foundIpRoute) {
		t.Errorf("Created ip address and found ip address do not match. actual: %v expected: %v", foundIpRoute, iproute)
	}

	err = c.DeleteIpRoute(iproute.Id)

	if err != nil {
		t.Errorf("Error deleting ip route with: %v", err)
	}
}
