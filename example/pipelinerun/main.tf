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

data "oz_microservice" "hello-world-app" {
    provider = ozone
    name = "hello-world"
    workspace_id = data.oz_project.default_project.id
}

data "oz_registry" "hello-world-registry" {
    provider = ozone
    name = "hello-world"
    workspace_id = data.oz_project.default_project.id
}

data "oz_repository" "hello-world-repo" {
    provider = ozone
    name = "go-hello-world"
    workspace_id = data.oz_project.default_project.id
}


data "oz_cluster" "default_cluster"{
    provider = ozone
    name = "localdev"
    workspace_id = data.oz_project.default_project.id
}

data "oz_environment" "dev_env"{
    provider = ozone
    name = "Development"
    workspace_id = data.oz_project.default_project.id
}

data "oz_variable" "test-var"{
    provider = ozone
    name = "test-var"
    workspace_id = data.oz_project.default_project.id
}

data "oz_pipeline" "gitpull-pipeline" {
  provider = ozone
  name = "ozone-git-pull"
  workspace_id = data.oz_project.default_project.id
}

resource "oz_pipelinerun" "pipeline" {
  provider = ozone

  application_id = data.oz_microservice.hello-world-app.id
  pipeline_id = data.oz_pipeline.gitpull-pipeline.id
  cluster_id = data.oz_cluster.default_cluster.id
  env_id = data.oz_environment.dev_env.id
  kind = "application"
  workspace_id = data.oz_project.default_project.id
  trigger_params = [
    {
      name = "REPO_BRANCH"
      type = 14
      description = "Repo branch"
      value = "master"
      default = ""
      required=true
      type_name = "Repository Branch"
      overrideField = ""
      image_tag_config = [{
        pick_latest_branch_commit = false
      }]
      resource_id = ""
    },
  ]  
}

output "pipelune2_output" {
  value = oz_pipelinerun.pipeline
}
