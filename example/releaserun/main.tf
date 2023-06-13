terraform {
  required_providers {
    ozone = {
      source = "ozone.com/com/ozone"
    }
  }
}

provider "ozone" {  
}


resource "oz_releaserun" "release" {
  provider = ozone

    name = "releaserun"
    release_id = "64785676b227a3b040cfff8c"
    steps = [
        {
          name = "build-ozone-ctl-swagger-api-6a006758"
          description = "Build"
          run_after = []
          jira_approval = []
          slack_approval = []
          status = []
          timeout = 0
          type_id = 1
          workspaces = []
          pipeline = [{
            application_id = "6461d9ae864db62ed52001e8"
            pipeline_id = "6461e79ccc9f9877d347000a"
            cluster_id = "642da86cbb3287bdb1a9fda2"
            env_id = "642da8c9d3491b96aaae8657"
            kind = "application"
            timeout = 0
            trigger_params = [
              {  
                name = "REPO_BRANCH"
                type = 14
                description = "Repo branch"
                value = "6461d913cf6d42d6b01ed39b:::main"
                default = ""
                required=true
                type_name = "Repository Branch"
                overrideField = ""
                image_tag_config = {
                  pick_latest_branch_commit = false
                }
                resource_id = ""
              },
  ] 
          }]
        }
    ]

}

output "release_output" {
  value = oz_releaserun.release
}
