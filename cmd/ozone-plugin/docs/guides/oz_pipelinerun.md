---
page_title: "oz_pipelinerun user guide - ozone"
subcategory: ""
description: |-
  documentation on how to use ozone provider
---

# Usage Guide

To trigger a pipeline run on ozone using terraform, we use the `oz_pipelinerun` resource with a set of required inputs.
given below is a sample terraform document which uses `ozone` provider to trigger a pipeline run on ozone.

```hcl
terraform {
  required_providers {
    ozone = {
      source = "hashicorp/ozone"
      version = "1.0.0"
    }
  }
}

provider "ozone" {
    # api_key = "values to be imported from env variables"
    # host = "values to be imported from env variables"
    # workspace_id = "values to be imported from env variables" 
}


resource "oz_pipelinerun" "pipeline" {
  provider = ozone

  application_id = "<application id>"
  pipeline_id = "<pipeline id>"
  cluster_id = "<cluster id>"
  env_id = "<environment id>"
  kind = "<pipeline kind eg: application>"
  trigger_params = [
    {  
      name = <param name>
      type = <param type>
      description = "param description"
      value = "value of the param"
      default = "default value if any"
      required="true if the param is required"
      type_name = "Repository Branch"
      overrideField = ""
      image_tag_config = {
        pick_latest_branch_commit = false
      }
      resource_id = "resource id if the param is being selected from the variable resources"
    }
    ...
    ...
    
  ]  
}

output "pipelune2_output" {
  value = oz_pipelinerun.pipeline
}
```