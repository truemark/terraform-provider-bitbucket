package terraform_bitbucket

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/truemark/terraform-provider-bitbucket/terraform_bitbucket/bitbucket_client"
)

func GetConfiguredAPIClient(ctx context.Context, d *schema.ResourceData, meta interface{}) *bitbucket_client.APIClient {

	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// configuration := bitbucket_client.NewConfiguration()
	// configuration.Auth = BasicAuth
	// api_client := bitbucket_client.NewAPIClient(configuration)

	auth := context.WithValue(context.Background(), bitbucket_client.ContextBasicAuth, bitbucket_client.BasicAuth{
		UserName: username,
		Password: password,
	})
	log.Println(auth)
	// r, err := bitbucket_client.Service.Operation(auth, args)
	return nil
}
