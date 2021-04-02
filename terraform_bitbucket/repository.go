package terraform_bitbucket

import (
	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: repositoryCreate,
		ReadContext:   repositoryRead,
		UpdateContext: repositoryUpdate,
		DeleteContext: repositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			// The repository's immutable id. This can be used as a substitute for the slug segment in URLs. Doing this guarantees your URLs will survive renaming of the repository by its owner, or even transfer of the repository to a different user.
			"uuid": {},
			// Uuid *string `json:"uuid,omitempty"`
			// The concatenation of the repository owner's username and the slugified name, e.g. \"evzijst/interruptingcow\". This is the same string used in Bitbucket URLs.
			"full_name": {},
			// FullName    *string     `json:"full_name,omitempty"`
			"is_private": {},
			// IsPrivate   *bool       `json:"is_private,omitempty"`
			"parent": {},
			// Parent      *Repository `json:"parent,omitempty"`
			"scm": {},
			// Scm         *string     `json:"scm,omitempty"`
			"owner": {},
			// Owner       *Account    `json:"owner,omitempty"`
			"name": {},
			// Name        *string     `json:"name,omitempty"`
			"description": {},
			// Description *string     `json:"description,omitempty"`
			"created_on": {},
			// CreatedOn   *time.Time  `json:"created_on,omitempty"`
			"updated_on": {},
			// UpdatedOn   *time.Time  `json:"updated_on,omitempty"`
			"size": {},
			// Size        *int32      `json:"size,omitempty"`
			"language": {},
			// Language    *string     `json:"language,omitempty"`
			"has_issues": {},
			// HasIssues   *bool       `json:"has_issues,omitempty"`
			"has_wiki": {},
			// HasWiki     *bool       `json:"has_wiki,omitempty"`
			//  Controls the rules for forking this repository.  * **allow_forks**: unrestricted forking * **no_public_forks**: restrict forking to private forks (forks cannot   be made public later) * **no_forks**: deny all forking

			"fork_policy": {},
			// ForkPolicy *string  `json:"fork_policy,omitempty"`
			"project": {},
			// Project    *Project `json:"project,omitempty"`
			"mainbranch": {},
			// Mainbranch *Branch  `json:"mainbranch,omitempty"`
			"links": {},
			// Links *RepositoryLinks `json:"links,omitempty"`
		},
	}
}

func repositoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func repositoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func repositoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func repositoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
