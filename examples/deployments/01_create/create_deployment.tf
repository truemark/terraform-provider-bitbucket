terraform {
  required_providers {
    truemark-bitbucket = {
      source = "truemark.io/terraform/truemark-bitbucket"
      version = "1.0.0"
    }
  }
}


resource "truemark-bitbucket_deployment" "myLittlePony_Deployment" {
  
}