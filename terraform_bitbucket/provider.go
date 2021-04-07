package terraform_bitbucket

import (
	"context"
	"log"

	"github.com/briancabbott/bitbucket_client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	// TODO: log.Printf("[INFO] Creating Provider")
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				// Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TRUEMARK_CONFLUENTCLOUD_USERNAME", ""),
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
				// Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TRUEMARK_CONFLUENTCLOUD_PASSWORD", ""),
			},
			"configuration": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			// "truemark-bitbucket_deployment": ResourceDeployment(),
			// "truemark-bitbucket_group":      ResourceGroup(),
			"truemark-bitbucket_project": ResourceProject(),
			// "truemark-bitbucket_repository": ResourceRepository(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	log.Printf("[INFO] Initializing BitBucket Client")

	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// if password !=
	auth_ctx := context.WithValue(ctx, bitbucket_client.ContextBasicAuth, bitbucket_client.BasicAuth{
		UserName: username,
		Password: password,
	})

	// var diags diag.Diagnostics
	// var c

	// err := resource.RetryContext(ctx, 30*time.Minute, func() *resource.RetryError {
	// 	err := c.Login()
	// 	if strings.Contains(err.Error(), "Exceeded rate limit") {
	// 		// TODO: log.Printf("[INFO] ConfluentCloud API rate limit exceeded, retrying.")
	// 		return resource.RetryableError(err)
	// 	}
	// 	return resource.NonRetryableError(err)
	// })
	return auth_ctx, nil // c, diag.FromErr(err)
}
