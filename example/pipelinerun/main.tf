terraform {
  required_providers {
    ozone = {
      source = "ozone.com/com/ozone"
    }
  }
}

provider "ozone" {  
}


resource "oz_pipelinerun" "pipeline" {
  provider = ozone

  application_id = "6461fff7cc9f9877d3470022"
  pipeline_id = "646216a5864db62ed5200213"
  cluster_id = "642da86cbb3287bdb1a9fda2"
  env_id = "642da8c9d3491b96aaae8657"
  kind = "application"
  trigger_params = [
    {  
      name = "REPO_BRANCH"
      type = 14
      description = "Repo branch"
      value = "6461ff09cf6d42d6b01ed3b3:::hotfix/swagger-specs-api"
      default = ""
      required=true
      type_name = "Repository Branch"
      overrideField = ""
      image_tag_config = {
        pick_latest_branch_commit = false
      }
      resource_id = ""
    },

    {  
      name = "IMAGE_TAG"
      type = 11
      description = "image tag"
      value = "iter-2"
      default = ""
      required=true
      type_name = "Registry Image Tag"
      overrideField = ""
      image_tag_config = {
        pick_latest_branch_commit = false
      }
      resource_id = ""
    },
    
  ]  
}

output "pipelune2_output" {
  value = oz_pipelinerun.pipeline
}
