terraform {
  required_providers {
    truemark-bitbucket = {
      source = "truemark.io/terraform/truemark-bitbucket"
      version = "1.0.0"
    }
  }
}

provider "truemark-bitbucket" {
  username = "briancabbott@gmail.com"
  password = ""
}

### 
## TerraForm Bitbucket Properties: 
##    - "name"        - the name of the project. 
##    - "key"         - a short-code reference to the project.
##    - "is_private"  - weather or not the repository is publically accessible. 
##    - "description" - a description of the project.
##    - "owner"       - the owner of the project.
resource "truemark-bitbucket_project" "terraform-project-11" {
  name             = "terraform-project-11"
  key              = "TFPROJ11"
  is_private       = true
  description      = "an overview of the project."
  owner            = "babbott@truemark.io"
  team_workspace_member {
    type = "workspace"
    name = "bitbucket-tfprovider-workspace"
  }
}
