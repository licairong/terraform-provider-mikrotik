package mikrotik

import (
	"context"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIpAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIpRouteRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address ID.",
						},
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address and netmask of the interface using slash notation.",
						},
						"comment": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comment for the IP address assignment.",
						},
						"disabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to disable IP address.",
						},
						"interface": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The interface on which the IP address is assigned.",
						},
						"network": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address for the network.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIpRouteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	ipaddr, err := c.ListIpAddress()

	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	orderItems, ids := flattenOrderItemsData(&ipaddr)
	if err := d.Set("items", orderItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ids)

	return diags
}

func flattenOrderItemsData(orderItems *[]client.IpAddress) ([]interface{}, string) {
	if orderItems != nil {
		ois := make([]interface{}, len(*orderItems), len(*orderItems))
		ids := ""

		for i, orderItem := range *orderItems {
			oi := make(map[string]interface{})

			oi["id"] = orderItem.Id
			oi["address"] = orderItem.Address
			oi["comment"] = orderItem.Comment
			oi["disabled"] = orderItem.Disabled
			oi["interface"] = orderItem.Interface
			oi["network"] = orderItem.Network

			ois[i] = oi

			ids += orderItem.Id
		}

		return ois, ids
	}

	return make([]interface{}, 0), ""
}
