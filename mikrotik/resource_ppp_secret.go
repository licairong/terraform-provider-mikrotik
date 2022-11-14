package mikrotik

import (
	"context"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePppSecret() *schema.Resource {
	return &schema.Resource{
		Description: "Add an PPP Secret.",

		CreateContext: resourcePppSecretCreate,
		ReadContext:   resourcePppSecretRead,
		UpdateContext: resourcePppSecretUpdate,
		DeleteContext: resourcePppSecretDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the user.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comment for the PPP Secret.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to disable PPP Secret.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "User password.",
			},
			"profile": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Profile name for the user.",
			},
			"routes": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Routes that appear on the server when the client is connected.",
			},
			"service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies service that will use this user.",
			},
		},
	}
}

func resourcePppSecretCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pppSecret := preparePppSecret(d)

	c := m.(*client.Mikrotik)

	secret, err := c.AddPppSecret(pppSecret)

	if err != nil {
		return diag.FromErr(err)
	}

	return secretToData(secret, d)
}

func resourcePppSecretRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	pppSecret, err := c.FindPppSecret(d.Id())

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

	return secretToData(pppSecret, d)
}

func resourcePppSecretUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	pppSecret := preparePppSecret(d)
	pppSecret.Id = d.Id()

	secret, err := c.UpdatePppSecret(pppSecret)

	if err != nil {
		return diag.FromErr(err)
	}

	return secretToData(secret, d)
}

func resourcePppSecretDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	err := c.DeletePppSecret(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func secretToData(secret *client.PppSecret, d *schema.ResourceData) diag.Diagnostics {
	values := map[string]interface{}{
		"name":   secret.Name,
		"comment":   secret.Comment,
		"disabled":  secret.Disabled,
		"password": secret.Password,
		"service": secret.Service,
		"routes": secret.Routes,
		"profile": secret.Profile,
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

func preparePppSecret(d *schema.ResourceData) *client.PppSecret {
	ipaddr := new(client.PppSecret)

	ipaddr.Comment = d.Get("comment").(string)
	ipaddr.Name = d.Get("name").(string)
	ipaddr.Disabled = d.Get("disabled").(bool)
	ipaddr.Password = d.Get("password").(string)
	ipaddr.Service = d.Get("service").(string)
	ipaddr.Routes = d.Get("routes").(string)
	ipaddr.Profile = d.Get("profile").(string)

	return ipaddr
}
