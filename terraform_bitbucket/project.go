package terraform_bitbucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// // No Read or Update for Service-Account
// func resourceServiceAccount() *schema.Resource {
// 	return &schema.Resource{
// 		CreateContext: ServiceAccountCreate,
// 		ReadContext:   ServiceAccountRead,
// 		// UpdateContext: ServiceAccountUpdate,
// 		DeleteContext: ServiceAccountDelete,
// 		Importer: &schema.ResourceImporter{
// 			StateContext: schema.ImportStatePassthroughContext,
// 		},
// 		Schema: map[string]*schema.Schema{
// 			"name": {
// 				Type:        schema.TypeString,
// 				Required:    true,
// 				ForceNew:    true,
// 				Description: "",
// 			},
// 			"description": {
// 				Type:        schema.TypeString,
// 				Required:    true,
// 				ForceNew:    true,
// 				Description: "Service Account Description",
// 			},
// 		},
// 	}
// }

func ResourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: ProjectCreate,
		ReadContext:   ProjectRead,
		UpdateContext: ProjectUpdate,
		DeleteContext: ProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		// A Bitbucket project. Projects are used by teams to organize repositories.
		Schema: map[string]*schema.Schema{
			"links": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "",
				Elem: &map[string]*schema.Schema{
					"html": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
						// ValidateFunc: ValidateResourceType,

						Description: "A link to a resource related to this object.",
						Elem: &map[string]*schema.Schema{
							"href": {
								Type:     schema.TypeString,
								Required: true,
								ForceNew: true,
								// ValidateFunc: ValidateResourceType,
							},
							"name": {
								Type:     schema.TypeString,
								Required: true,
								ForceNew: true,
								// ValidateFunc: ValidateResourceType,
							},
						},
					},
					"avatar": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "A link to a resource related to this object.",
						// Elem:
						// "href": {
						// 	Type: schema.TypeString,
						// 	Required:     true,
						// 	ForceNew:     true,
						// 	ValidateFunc: ValidateResourcePatternType,
						// },
						// "name": {
						// 	Type:         schema.TypeString,
						// 	Required:     true,
						// 	ForceNew:     true,
						// 	ValidateFunc: ValidateResourcePatternType,
						// },
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The project's immutable id.",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The project's key.",
			},
			"owner": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the project.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "",
			},
			"is_private": {
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Indicates whether the project is publicly accessible, or whether it is private to the team and consequently only visible to team members. Note that private projects cannot contain public repositories.",
			},
			"created_on": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "",
			},
			"updated_on": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "",
			},
		},
	}
}

func projectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func projectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func projectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func projectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
