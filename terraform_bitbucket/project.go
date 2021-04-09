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
	log.Println("BCA-DBG: Into => terraform_bitbucket.ResourceProject()")

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
				Optional:    true,
				Description: "The project's immutable id.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the project.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project's key.",
			},
			"team_workspace_member": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The name of the Workspace or Team that this Project Belongs to. Note: this is a TrueMark created value for signalling the correct assocation for creation/deletion/listing, etc and not a BitBucket suppported value.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							// ValidateFunc: validateTeamWorkspaceMembership,
							Description: "('team'|'workspace').",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the project or workspace.",
						},
					},
				},
			},
			"owner": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			"is_private": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Indicates whether the project is publicly accessible, or whether it is private to the team and consequently only visible to team members. Note that private projects cannot contain public repositories.",
			},
			"created_on": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "",
			},
			"updated_on": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "",
			},
			"links": {
				Type:        schema.TypeList,
				Required:    false,
				Optional:    true,
				Description: "",
				// Elem: &map[string]*schema.Schema{
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"html": {
							Type:     schema.TypeList,
							Required: false,
							Optional: true,
							// ValidateFunc: ValidateResourceType,
							Description: "A link to a resource related to this object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:     schema.TypeString,
										Required: false,
										Optional: true,
										// ValidateFunc: ValidateResourceType,
									},
									"name": {
										Type:     schema.TypeString,
										Required: false,
										Optional: true,
										// ValidateFunc: ValidateResourceType,
									},
								},
							},
						},
						"avatar": {
							Type:        schema.TypeList,
							Required:    false,
							Optional:    true,
							Description: "A link to a resource related to this object.",
							// Elem: &map[string]*schema.Schema{
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:     schema.TypeString,
										Required: false,
										Optional: true,
										// ValidateFunc: ValidateResourcePatternType,
									},
									"name": {
										Type:     schema.TypeString,
										Required: false,
										Optional: true,
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
	log.Println("BCA-DBG: Into => terraform_bitbucket.projectCreate()")

	fmt.Fprintf(os.Stderr, "Entering projectCreate()\n")

	membership := d.Get("team_workspace_member")
	fmt.Fprintf(os.Stderr, "membership: %v\n", membership)

	project := *bitbucket_client.NewProject()

	nameParam := d.Get("name").(string)
	project.Name = &nameParam

	description := d.Get("description").(string)
	project.Description = &description

	key := d.Get("key").(string)
	project.Key = &key

	isPrivate := d.Get("is_private").(bool)
	project.IsPrivate = &isPrivate

	// project.Owner = d.Get("owner")
	// links := d.Get("links")

	configuration := bitbucket_client.NewConfiguration()
	configuration.Debug = true
	api_client := bitbucket_client.NewAPIClient(configuration)

	auth := context.WithValue(ctx, bitbucket_client.ContextBasicAuth, meta)
	projectRequest := api_client.ProjectsApi.WorkspacesWorkspaceProjectsPost(auth, "bitbucket-tfprovider-workspace")

	newProject, response, err := projectRequest.Body(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.WorkspacesWorkspaceProjectsPost``: %v\n", err)
	}
	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", response)
	fmt.Fprintf(os.Stderr, "Response from `ProjectsApi.WorkspacesWorkspaceProjectsPost`: %v\n", newProject)

	return nil // diag.FromErr(err)
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
