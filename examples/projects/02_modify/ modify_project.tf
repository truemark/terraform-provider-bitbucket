terraform {
  required_providers {
    truemark-bitbucket = {
      source = "truemark.io/terraform/truemark-bitbucket"
      version = "1.0.0"
    }
  }
}


resource "truemark-bitbucket_project" "myLittlePony_Project" {
  name             = "test-terraform2_aSecondName"
  key              = "PROJK"
  is_private       = true
  description      = "a thesis length (or shorter) overview of this project."
  owner            = "briancabbott@gmail.com"
}
