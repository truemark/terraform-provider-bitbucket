package terraform_bitbucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

///
// Deployment Schema:
//    - uuid
//    - state
//    - environment
//    - release

func ResourceDeployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: DeploymentCreate,
		ReadContext:   DeploymentRead,
		UpdateContext: DeploymentUpdate,
		DeleteContext: DeploymentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		////
		// TypeBool
		// TypeInt
		// TypeFloat
		// TypeString
		//   // Date & Time Data
		// 		Type:         schema.TypeString,
		// 		ValidateFunc: validation.ValidateRFC3339TimeString,
		// TypeList
		// TypeMap
		// TypeSet

		// A Bitbucket project. Projects are used by teams to organize repositories.
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Description: "The UUID identifying the deployment.",
			},
			"state": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "",
				// "$ref": "#/definitions/deployment_state"
				Elem: &map[string]*schema.Schema{
					// "deployment_state": {
					// 	Type:         schema.TypeString,
					// 	Required:     true,
					// 	ForceNew:     true,
					// 	Description:  "The representation of the progress state of a deployment.",
					// 	ValidateFunc: validation.ValidateRFC3339TimeString,
					// },
					"deployment_state_completed_status": {
						Type:         schema.TypeString,
						Required:     true,
						ForceNew:     true,
						Description:  "The status of a completed deployment.",
						ValidateFunc: validation.ValidateRFC3339TimeString,
					},
					"deployment_state_in_progress": {
						Type:        schema.TypeList,
						Required:    true,
						ForceNew:    true,
						Description: "A Bitbucket Deployment IN_PROGRESS deployment state.",
						Enum: &map[string]*schema.Schema{
							"name": {
								// "enum": [
								// 	"IN_PROGRESS"
								// ],
								Type:        schema.TypeList,
								Required:    true,
								ForceNew:    true,
								Description: "The name of deployment state (IN_PROGRESS).",
								Enum: &map[string]*schema.Schema{
									"IN_PROGRESS": {},
								},
							},
							"url": {
								Type:        schema.TypeString,
								Required:    true,
								ForceNew:    true,
								Description: "Link to the deployment result.",
							},
							"deployer": {
								// "$ref": "#/definitions/account",
								Type:        schema.TypeList,
								Required:    true,
								ForceNew:    true,
								Description: "The Bitbucket account that was used to perform the deployment.",
								//
							},
							"start_date": {
								// Type: 		 schema.TypeList,
								Type: schema.TypeString,

								Required:     true,
								ForceNew:     true,
								Description:  "The timestamp when the deployment was started.",
								ValidateFunc: validation.ValidateRFC3339TimeString,
								// "type": "string",
								// "format": "date-time",
								// "description": "The timestamp when the deployment was started."
							},
						},
					},
					"deployment_state_undeployed": {
						Type:        schema.TypeList,
						Required:    true,
						ForceNew:    true,
						Description: "A Bitbucket Deployment IN_PROGRESS deployment state.",
						Enum: &map[string]*schema.Schema{
							"name": {
								// "enum": [
								// 	"UNDEPLOYED"
								// ],
								Type:        schema.TypeString,
								Required:    true,
								ForceNew:    true,
								Description: "The name of deployment state (UNDEPLOYED).",
							},
							"trigger_url": {
								Type: schema.TypeString,
								// "format": "uri",
								Required:    true,
								ForceNew:    true,
								Description: "Link to trigger the deployment.",
							},
						},
					},
					"deployment_state_completed_status_stopped": {
						Type: schema.TypeString,
						// "format": "uri",
						Required:    true,
						ForceNew:    true,
						Description: "Link to trigger the deployment.",
						Enum: &map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Required:    true,
								ForceNew:    true,
								Description: "The name of the completed deployment status (STOPPED).",
								// "enum": [
								// 	"STOPPED"
								// ],
							},
						},
					},
					"deployment_state_completed_status_failed": {
						Type: schema.TypeString,
						// "format": "uri",
						Required:    true,
						ForceNew:    true,
						Description: "Link to trigger the deployment.",
						Enum: &map[string]*schema.Schema{
							"name": {
								// "enum": [
								// 	"FAILED"
								// ],
								Type:        schema.TypeString,
								Required:    true,
								ForceNew:    true,
								Description: "The name of the completed deployment status (FAILED).",
							},
						},
					},
				},
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "A deployment environment.",
				// "$ref": "#/definitions/deployment_environment",
				// "description": ""
				Enum: &map[string]*schema.Schema{
					"deployment_environment": {
						Type:        schema.TypeList,
						Required:    true,
						ForceNew:    true,
						Description: "A Bitbucket Deployment Environment.",
						Enum: &map[string]*schema.Schema{
							"uuid": {
								Type:        schema.TypeString,
								Required:    true,
								ForceNew:    true,
								Description: "The UUID identifying the environment.",
							},
							"name": {
								Type:        schema.TypeString,
								Required:    true,
								ForceNew:    true,
								Description: "The name of the environment.",
							},
						},
					},
				},
			},
			"release": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The release associated with this deployment.",
				// "$ref": "#/definitions/deployment_release",
				Enum: &map[string]*schema.Schema{
					"uuid": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "The UUID identifying the release.",
					},
					"name": {
						Type:          schema.TypeString,
						Required:      true,
						ForceNew:      true,
						"description": "The name of the release.",
					},
					"url": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
						// "format": "uri",
						Description: "Link to the pipeline that produced the release.",
					},
					"commit": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
						Description: "The commit associated with this release.",
						Enum: &map[string]*schema.Schema{},
					},
					"created_on": {
						Type:         schema.TypeString,
						Required:     true,
						ForceNew:     true,
						Description:  "The timestamp when the release was created.",
						ValidateFunc: validation.ValidateRFC3339TimeString,
					},
				},
			},
		},
	}
}

func DeploymentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func DeploymentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func DeploymentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func DeploymentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
