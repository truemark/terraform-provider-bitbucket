package terraform_bitbucket

import (
	"context"
	"log"

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
		CreateContext: projectCreate,
		ReadContext:   projectRead,
		UpdateContext: projectUpdate,
		DeleteContext: projectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		// A Bitbucket project. Projects are used by teams to organize repositories.
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Required:    false,
				ForceNew:    true,
				Description: "The project's immutable id.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the project.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				ForceNew:    true,
				Description: "",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The project's key.",
			},
			"team_workspace_member": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Workspace or Team that this Project Belongs to. Note: this is a TrueMark created value for signalling the correct assocation for creation/deletion/listing, etc and not a BitBucket suppported value.",
				Elem: &map[string]*schema.Schema{
					"type": {
						Type:         schema.TypeString,
						Required:     true,
						ForceNew:     true,
						ValidateFunc: validateTeamWorkspaceMembership,
						Description:  "('team'|'workspace').",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "The name of the project or workspace.",
					},
				},
			},
			"owner": {
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
				Required:    false,
				ForceNew:    true,
				Description: "",
			},
			"updated_on": {
				Type:        schema.TypeString,
				Required:    false,
				ForceNew:    true,
				Description: "",
			},
			"links": {
				Type:        schema.TypeList,
				Required:    false,
				ForceNew:    true,
				Description: "",
				Elem: &map[string]*schema.Schema{
					"html": {
						Type:     schema.TypeList,
						Required: false,
						ForceNew: true,
						// ValidateFunc: ValidateResourceType,
						Description: "A link to a resource related to this object.",
						Elem: &map[string]*schema.Schema{
							"href": {
								Type:     schema.TypeString,
								Required: false,
								ForceNew: true,
								// ValidateFunc: ValidateResourceType,
							},
							"name": {
								Type:     schema.TypeString,
								Required: false,
								ForceNew: true,
								// ValidateFunc: ValidateResourceType,
							},
						},
					},
					"avatar": {
						Type:        schema.TypeList,
						Required:    false,
						ForceNew:    true,
						Description: "A link to a resource related to this object.",
						Elem: &map[string]*schema.Schema{
							"href": {
								Type:     schema.TypeString,
								Required: false,
								ForceNew: true,
								// ValidateFunc: ValidateResourcePatternType,
							},
							"name": {
								Type:     schema.TypeString,
								Required: false,
								ForceNew: true,
								// ValidateFunc: ValidateResourcePatternType,
							},
						},
					},
				},
			},
		},
	}
}

func projectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	membership := d.Get("team_workspace_member")
	log.Println(membership)
	// workspace := "workspace_example"       // string | This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: `{workspace UUID}`.
	// body := *bitbucket_client.NewProject() // Project |

	// configuration := openapiclient.NewConfiguration()
	// api_client := openapiclient.NewAPIClient(configuration)
	// resp, r, err := api_client.ProjectsApi.WorkspacesWorkspaceProjectsPost(context.Background(), workspace).Body(body).Execute()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.WorkspacesWorkspaceProjectsPost``: %v\n", err)
	// 	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	// }
	// // response from `WorkspacesWorkspaceProjectsPost`: Project
	// fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.WorkspacesWorkspaceProjectsPost`: %v\n", resp)
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

func validateTeamWorkspaceMembership(i interface{}, k string) (warnings []string, errors []error) {
	return nil, nil
}
