terraform {
  required_providers {
    ozone = {
      source = "ozone.com/com/ozone"
    }
  }
}

provider "ozone" {
}

data "oz_project" "default_project"{
    provider = ozone
    name = "default"
}

data "oz_pipeline" "gitpull-pipeline" {
  provider = ozone
  name = "ozone-git-pull"
  workspace_id = data.oz_project.default_project.id
}


output "pipeline_params" {
    value = data.oz_pipeline.gitpull-pipeline.params
}

