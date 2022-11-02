package mikrotik

import (
	"context"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpRoute() *schema.Resource {
	return &schema.Resource{
		Description: "Add an IP Route.",

		CreateContext: resourceIpRouteCreate,
		ReadContext:   resourceIpRouteRead,
		UpdateContext: resourceIpRouteUpdate,
		DeleteContext: resourceIpRouteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"dst_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Destination address.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comment for the IP Route.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to disable IP Route.",
			},
			"gateway": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The interface on which the IP Route.",
			},
		},
	}
}

func resourceIpRouteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	ipRoute := prepareIpRoute(d)

	c := m.(*client.Mikrotik)

	iproute, err := c.AddIpRoute(ipRoute)

	if err != nil {
		return diag.FromErr(err)
	}

	return routeToData(iproute, d)
}

func resourceIpRouteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	iproute, err := c.FindIpRoute(d.Id())

	// Clear the state if the error represents that the resource no longer exists
	_, resourceMissing := err.(*client.NotFound)
	if resourceMissing && err != nil {
		d.SetId("")
		return nil
	}

	// Make sure all other errors are propagated
	if err != nil {
		return diag.FromErr(err)
	}

	return routeToData(iproute, d)
}

func resourceIpRouteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	ipRoute := prepareIpRoute(d)
	ipRoute.Id = d.Id()

	iproute, err := c.UpdateIpRoute(ipRoute)

	if err != nil {
		return diag.FromErr(err)
	}

	return routeToData(iproute, d)
}

func resourceIpRouteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	err := c.DeleteIpRoute(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func routeToData(iproute *client.IpRoute, d *schema.ResourceData) diag.Diagnostics {
	values := map[string]interface{}{
		"dst_address":   iproute.DstAddress,
		"comment":   iproute.Comment,
		"disabled":  iproute.Disabled,
		"gateway": iproute.Gateway,
	}

	d.SetId(iproute.Id)

	var diags diag.Diagnostics

	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			diags = append(diags, diag.Errorf("failed to set %s: %v", key, err)...)
		}
	}

	return diags
}

func prepareIpRoute(d *schema.ResourceData) *client.IpRoute {
	ipaddr := new(client.IpRoute)

	ipaddr.Comment = d.Get("comment").(string)
	ipaddr.DstAddress = d.Get("dst_address").(string)
	ipaddr.Disabled = d.Get("disabled").(bool)
	ipaddr.Gateway = d.Get("gateway").(string)

	return ipaddr
}
