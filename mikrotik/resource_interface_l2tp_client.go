package mikrotik

import (
	"context"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceL2tpClient() *schema.Resource {
	return &schema.Resource{
		Description: "Add an L2tp Client.",

		CreateContext: resourceL2tpClientCreate,
		ReadContext:   resourceL2tpClientRead,
		UpdateContext: resourceL2tpClientUpdate,
		DeleteContext: resourceL2tpClientDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the client.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comment for the L2tp Client.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to disable L2tp Client.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "User password.",
			},
			"ipsec_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Ipsec secret.",
			},
			"connect_to": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Connect to host.",
			},
			"use_ipsec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Use ipsec.",
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Connect user.",
			},
		},
	}
}

func resourceL2tpClientCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	l2tpClient := prepareL2tpClient(d)

	c := m.(*client.Mikrotik)

	l2tpclient, err := c.AddL2tpClient(l2tpClient)

	if err != nil {
		return diag.FromErr(err)
	}

	return l2tpClientToData(l2tpclient, d)
}

func resourceL2tpClientRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	l2tpClient, err := c.FindL2tpClient(d.Id())

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

	return l2tpClientToData(l2tpClient, d)
}

func resourceL2tpClientUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	L2tpClient := prepareL2tpClient(d)
	L2tpClient.Id = d.Id()

	l2tpClient, err := c.UpdateL2tpClient(L2tpClient)

	if err != nil {
		return diag.FromErr(err)
	}

	return l2tpClientToData(l2tpClient, d)
}

func resourceL2tpClientDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	err := c.DeleteL2tpClient(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func l2tpClientToData(secret *client.L2tpClient, d *schema.ResourceData) diag.Diagnostics {
	values := map[string]interface{}{
		"name":   secret.Name,
		"comment":   secret.Comment,
		"disabled":  secret.Disabled,
		"password": secret.Password,
		"connect_to": secret.ConnectTo,
		"ipsec_secret": secret.IpsecSecret,
		"use_ipsec": secret.UseIpsec,
		"user": secret.User,
	}

	d.SetId(secret.Id)

	var diags diag.Diagnostics

	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			diags = append(diags, diag.Errorf("failed to set %s: %v", key, err)...)
		}
	}

	return diags
}

func prepareL2tpClient(d *schema.ResourceData) *client.L2tpClient {
	l2tpClient := new(client.L2tpClient)

	l2tpClient.Comment = d.Get("comment").(string)
	l2tpClient.Name = d.Get("name").(string)
	l2tpClient.Disabled = d.Get("disabled").(bool)
	l2tpClient.Password = d.Get("password").(string)
	l2tpClient.ConnectTo = d.Get("connect_to").(string)
	l2tpClient.IpsecSecret = d.Get("ipsec_secret").(string)
	l2tpClient.UseIpsec = d.Get("use_ipsec").(string)
	l2tpClient.User = d.Get("user").(string)

	return l2tpClient
}
