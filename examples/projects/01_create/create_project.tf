terraform {
  required_providers {
    truemark-bitbucket = {
      source = "truemark.io/terraform/truemark-bitbucket"
      version = "1.0.0"
    }
  }
}


### 
## TerraForm Bitbucket Properties: 
##    - "name"        - the name of the project. 
##    - "key"         - a short-code reference to the project.
##    - "is_private"  - weather or not the repository is publically accessible. 
##    - "description" - a description of the project.
##    - "owner"       - the owner of the project.
resource "truemark-bitbucket_project" "myLittlePony_Project" {
  name             = "test-terraform2"
  key              = "PROJK"
  is_private       = true
  description      = "a thesis length (or shorter) overview of this project."
  owner            = "briancabbott@gmail.com"
}
