package terraform_bitbucket

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/briancabbott/bitbucket_client"
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
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				// ForceNew:    true,
				Description: "The project's immutable id.",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				// ForceNew:    true,
				Description: "The name of the project.",
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				// ForceNew:    true,
				Description: "",
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				// ForceNew:    true,
				Description: "The project's key.",
			},
			"team_workspace_member": {
				Type:     schema.TypeList,
				Required: true,
				// ForceNew:    true,
				Description: "The name of the Workspace or Team that this Project Belongs to. Note: this is a TrueMark created value for signalling the correct assocation for creation/deletion/listing, etc and not a BitBucket suppported value.",
				// Elem: &map[string]*schema.Schema{
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							// ForceNew: true,
							// ValidateFunc: validateTeamWorkspaceMembership,
							Description: "('team'|'workspace').",
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
							// ForceNew:    true,
							Description: "The name of the project or workspace.",
						},
					},
				},
			},
			"owner": {
				Type:     schema.TypeString,
				Required: true,
				// ForceNew:    true,
				Description: "",
			},
			"is_private": {
				Type:     schema.TypeBool,
				Required: true,
				// ForceNew:    true,
				Description: "Indicates whether the project is publicly accessible, or whether it is private to the team and consequently only visible to team members. Note that private projects cannot contain public repositories.",
			},
			"created_on": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				// ForceNew:    true,
				Description: "",
			},
			"updated_on": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				// ForceNew:    true,
				Description: "",
			},
			"links": {
				Type:     schema.TypeList,
				Required: false,
				Optional: true,
				// ForceNew:    true,
				Description: "",
				// Elem: &map[string]*schema.Schema{
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"html": {
							Type:     schema.TypeList,
							Required: false,
							Optional: true,
							// ForceNew: true,
							// ValidateFunc: ValidateResourceType,
							Description: "A link to a resource related to this object.",
							// Elem: &map[string]*schema.Schema{
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:     schema.TypeString,
										Required: false,
										Optional: true,
										// ForceNew: true,
										// ValidateFunc: ValidateResourceType,
									},
									"name": {
										Type:     schema.TypeString,
										Required: false,
										Optional: true,
										// ForceNew: true,
										// ValidateFunc: ValidateResourceType,
									},
								},
							},
						},
						"avatar": {
							Type:     schema.TypeList,
							Required: false,
							Optional: true,
							// ForceNew:    true,
							Description: "A link to a resource related to this object.",
							// Elem: &map[string]*schema.Schema{
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:     schema.TypeString,
										Required: false,
										Optional: true,
										// ForceNew: true,
										// ValidateFunc: ValidateResourcePatternType,
									},
									"name": {
										Type:     schema.TypeString,
										Required: false,
										Optional: true,
										// ForceNew: true,
										// ValidateFunc: ValidateResourcePatternType,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// type Project struct {
// 	Links *ProjectLinks `json:"links,omitempty"`

// 	// The project's immutable id.
// 	Uuid *string `json:"uuid,omitempty"`

// 	// The project's key.
// 	Key *string `json:"key,omitempty"`

// 	Owner *Team `json:"owner,omitempty"`

// 	// The name of the project.
// 	Name *string `json:"name,omitempty"`

// 	Description *string `json:"description,omitempty"`

// 	// Indicates whether the project is publicly accessible, or whether it is private
// 	// to the team and consequently only visible to team members.
// 	// Note that private projects cannot contain public repositories.
// 	IsPrivate *bool      `json:"is_private,omitempty"`
// 	CreatedOn *time.Time `json:"created_on,omitempty"`
// 	UpdatedOn *time.Time `json:"updated_on,omitempty"`
// }

func projectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	// username := d.Get("username").(string)
	// password := d.Get("password").(string)

	log.Println("Entering projectCreate()")

	membership := d.Get("team_workspace_member")
	log.Println(membership)

	body := *bitbucket_client.NewProject() // Project |

	tmp_name := d.Get("name").(string)
	body.Name = &tmp_name

	tmp_description := d.Get("description").(string)
	body.Description = &tmp_description

	tmp_key := d.Get("key").(string)
	body.Key = &tmp_key

	tmp_isprivate := d.Get("is_private").(bool)
	body.IsPrivate = &tmp_isprivate

	// body.Owner = d.Get("owner")
	// links := d.Get("links")

	configuration := bitbucket_client.NewConfiguration()
	configuration.Debug = true
	api_client := bitbucket_client.NewAPIClient(configuration)
	// auth_ctx := context.WithValue(context.Background(), bitbucket_client.ContextBasicAuth, bitbucket_client.BasicAuth{
	// 	UserName: username,
	// 	Password: password,
	// })

	log.Println("")
	log.Println("Meta: ")
	log.Println(meta)
	c := meta.(context.Context)
	resp, r, err := api_client.ProjectsApi.WorkspacesWorkspaceProjectsPost(c, "bitbucket-tfprovider-workspace").Body(body).Execute()
	log.Println("completed execute call for HTTP ")
	log.Println(resp)
	log.Println(r)
	log.Println(err)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error when calling `ProjectsApi.WorkspacesWorkspaceProjectsPost``: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Full HTTP response: %v\n", r)
	// }
	// response from `WorkspacesWorkspaceProjectsPost`: Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.WorkspacesWorkspaceProjectsPost`: %v\n", resp)

	return diag.FromErr(err)
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
