---
page_title: "oz_releaserun user guide - ozone"
subcategory: ""
description: |-
  documentation on how to use ozone releaserun provider
---

# Usage Guide

To trigger a release run on ozone using terraform, we use the `oz_releaserun` resource with a set of required inputs.
given below is a sample terraform document which uses `ozone` provider to trigger a release run on ozone.

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
}


resource "oz_releaserun" "release" {
  provider = ozone

    name = "<name of the releaserun>"
    release_id = "<release id for which release has to be triggered>"
    steps = [
        {
          name = "<step name>"
          description = "<step description>"
          run_after = []
          jira_approval = []
          slack_approval = []
          status = []
          timeout = 0
          type_id = 1
          workspaces = []
          pipeline = [{
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
          }]
        }
    ]

}

output "release_output" {
  value = oz_releaserun.release
}

```