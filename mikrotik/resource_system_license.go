package mikrotik

import (
	"context"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemLicense() *schema.Resource {
	return &schema.Resource{
		Description: "Add an system license.",

		CreateContext: resourceSystemLicenseCreate,
		ReadContext:   resourceSystemLicenseRead,
		UpdateContext: resourceSystemLicenseUpdate,
		DeleteContext: resourceSystemLicenseDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"account": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Mikrotik user.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "User password.",
			},
			"level": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "License level.",
			},
		},
	}
}

func resourceSystemLicenseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	systemLicense := prepareSystemLicense(d)

	c := m.(*client.Mikrotik)

	license, err := c.AddSystemLicense(systemLicense)

	if err != nil {
		return diag.FromErr(err)
	}

	return licenseToData(license, d)
}

func resourceSystemLicenseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	systemLicense := prepareSystemLicense(d)
	return licenseToData(systemLicense, d)
}

func resourceSystemLicenseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceSystemLicenseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func licenseToData(license *client.SystemLicense, d *schema.ResourceData) diag.Diagnostics {
	values := map[string]interface{}{
		"account":   license.Account,
		"level":   license.Level,
	}

	d.SetId(license.Level)

	var diags diag.Diagnostics

	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			diags = append(diags, diag.Errorf("failed to set %s: %v", key, err)...)
		}
	}

	return diags
}

func prepareSystemLicense(d *schema.ResourceData) *client.SystemLicense {
	license := new(client.SystemLicense)

	license.Account = d.Get("account").(string)
	license.Password = d.Get("password").(string)
	license.Level = d.Get("level").(string)

	return license
}