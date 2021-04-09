package terraform_bitbucket

import (
	"context"
	"log"

	"github.com/briancabbott/bitbucket_client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	log.Println("BCA-DBG: Into => terraform_bitbucket.Provider()")
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
	log.Println("BCA-DBG: Into => terraform_bitbucket.providerConfigure()")
	log.Printf("[INFO] Initializing BitBucket Client")

	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// if password !=
	auth := bitbucket_client.BasicAuth{
		UserName: username,
		Password: password,
	}

	// var err error = nil
	return auth, nil
}
