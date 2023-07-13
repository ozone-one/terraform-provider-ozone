---
page_title: "oz_pipelinerun user guide - ozone"
subcategory: ""
description: |-
  documentation on how to use ozone provider
---

# Usage Guide

To trigger a pipeline run on ozone using terraform, we use the `oz_pipelinerun` resource with a set of required inputs.
given below is a sample terraform document which uses `ozone` provider to trigger a pipeline run on ozone.
The expected pipeline inputs can be retrieved via the `oz_pipeline` data source.

here is an example how to extract the pipeline params from the pipeline

```hcl
terraform {
  required_providers {
    ozone = {
      source = "hashicorp/ozone"
      version = "1.0.1"
    }
  }
}

provider "ozone" {
    # api_key = "values to be imported from env variables" - required
    # host = "values to be imported from env variables" - required
    # base_path = "the base path to be used" - optional
    # scheme = "http or https (https by default)" - optional
}

data "oz_project" "default_project"{
    provider = ozone
    name = "<your project name here>"
}

data "oz_pipeline" "gitpull-pipeline" {
  provider = ozone
  name = "your pipeline name here"
  workspace_id = data.oz_project.default_project.id
}

output "pipelune_params" {
  value = data.oz_pipeline.gitpull-pipeline.params  
}

```

on the above terraform file, run the following commands

```bash
terraform apply && terraform show
```

the output will be the paramaters which the pipeline has
```bash
# data.oz_pipeline.gitpull-pipeline:
data "oz_pipeline" "gitpull-pipeline" {
    account_id   = "000000000000000000000000"
    id           = "606c44882961cab98f9ee525"
    name         = "ozone-git-pull-docker-image-build-and-template-deploy-v-1"
    params       = [
        {
            default          = ""
            description      = <<-EOT
                Name of registry
            EOT
            image_tag_config = []
            name             = "REGISTRY_NAME"
            required         = true
            resource_id      = ""
            type             = 7
            type_name        = "Registry Name"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                Registry user
            EOT
            image_tag_config = []
            name             = "REGISTRY_USER"
            required         = true
            resource_id      = ""
            type             = 8
            type_name        = "Registry Username"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                Registry password
            EOT
            image_tag_config = []
            name             = "REGISTRY_PASSWORD"
            required         = true
            resource_id      = ""
            type             = 9
            type_name        = "Registry Password"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                Repository directory
            EOT
            image_tag_config = []
            name             = "PROJECT_DIR"
            required         = true
            resource_id      = ""
            type             = 12
            type_name        = "Application Name"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                Repository user
            EOT
            image_tag_config = []
            name             = "REPO_USER"
            required         = true
            resource_id      = ""
            type             = 4
            type_name        = "Repository Username"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                Repository password
            EOT
            image_tag_config = []
            name             = "REPO_PASSWORD"
            required         = true
            resource_id      = ""
            type             = 5
            type_name        = "Repository Password/ Access Token"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                Repository url
            EOT
            image_tag_config = []
            name             = "REPO_URL"
            required         = true
            resource_id      = ""
            type             = 3
            type_name        = "Repository URL"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                Repo branch
            EOT
            image_tag_config = []
            name             = "REPO_BRANCH"
            required         = true
            resource_id      = ""
            type             = 14
            type_name        = "Repository Branch"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                registry image name
            EOT
            image_tag_config = []
            name             = "REGISTRY_IMAGE"
            required         = true
            resource_id      = ""
            type             = 10
            type_name        = "Registry Image"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                image tag
            EOT
            image_tag_config = []
            name             = "IMAGE_TAG"
            required         = true
            resource_id      = ""
            type             = 11
            type_name        = "Registry Image Tag"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                docker file path
            EOT
            image_tag_config = []
            name             = "DOCKER_FILE_PATH"
            required         = true
            resource_id      = ""
            type             = 13
            type_name        = "Dockerfile Path"
            value            = ""
        },
        {
            default          = ""
            description      = <<-EOT
                docker git param
            EOT
            image_tag_config = []
            name             = "DOCKER_NETRC"
            required         = true
            resource_id      = ""
            type             = 6
            type_name        = "Repository NETRC"
            value            = ""
        },
        {
            default          = ""
            description      = "YAML string to apply"
            image_tag_config = []
            name             = "YAML_STRING"
            required         = true
            resource_id      = ""
            type             = 15
            type_name        = "Parsed Templates"
            value            = ""
        },
    ]
    workspace_id = "64a54f77343ea404d2d55f84"
}

# data.oz_project.default_project:
data "oz_project" "default_project" {
    id   = "64a54f77343ea404d2d55f84"
    name = "default"
}


Outputs:

pipelune_params = [
    {
        default          = ""
        description      = <<-EOT
            Name of registry
        EOT
        image_tag_config = []
        name             = "REGISTRY_NAME"
        required         = true
        resource_id      = ""
        type             = 7
        type_name        = "Registry Name"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            Registry user
        EOT
        image_tag_config = []
        name             = "REGISTRY_USER"
        required         = true
        resource_id      = ""
        type             = 8
        type_name        = "Registry Username"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            Registry password
        EOT
        image_tag_config = []
        name             = "REGISTRY_PASSWORD"
        required         = true
        resource_id      = ""
        type             = 9
        type_name        = "Registry Password"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            Repository directory
        EOT
        image_tag_config = []
        name             = "PROJECT_DIR"
        required         = true
        resource_id      = ""
        type             = 12
        type_name        = "Application Name"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            Repository user
        EOT
        image_tag_config = []
        name             = "REPO_USER"
        required         = true
        resource_id      = ""
        type             = 4
        type_name        = "Repository Username"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            Repository password
        EOT
        image_tag_config = []
        name             = "REPO_PASSWORD"
        required         = true
        resource_id      = ""
        type             = 5
        type_name        = "Repository Password/ Access Token"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            Repository url
        EOT
        image_tag_config = []
        name             = "REPO_URL"
        required         = true
        resource_id      = ""
        type             = 3
        type_name        = "Repository URL"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            Repo branch
        EOT
        image_tag_config = []
        name             = "REPO_BRANCH"
        required         = true
        resource_id      = ""
        type             = 14
        type_name        = "Repository Branch"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            registry image name
        EOT
        image_tag_config = []
        name             = "REGISTRY_IMAGE"
        required         = true
        resource_id      = ""
        type             = 10
        type_name        = "Registry Image"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            image tag
        EOT
        image_tag_config = []
        name             = "IMAGE_TAG"
        required         = true
        resource_id      = ""
        type             = 11
        type_name        = "Registry Image Tag"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            docker file path
        EOT
        image_tag_config = []
        name             = "DOCKER_FILE_PATH"
        required         = true
        resource_id      = ""
        type             = 13
        type_name        = "Dockerfile Path"
        value            = ""
    },
    {
        default          = ""
        description      = <<-EOT
            docker git param
        EOT
        image_tag_config = []
        name             = "DOCKER_NETRC"
        required         = true
        resource_id      = ""
        type             = 6
        type_name        = "Repository NETRC"
        value            = ""
    },
    {
        default          = ""
        description      = "YAML string to apply"
        image_tag_config = []
        name             = "YAML_STRING"
        required         = true
        resource_id      = ""
        type             = 15
        type_name        = "Parsed Templates"
        value            = ""
    },
]
```
the above output can be used to trigger a pipeline run.
By default, ozone will take care of linked parameters in the pipeline and only selectable params are required for missing input.
Those can be found on the pipeliune run page of ozone itself.

```hcl
terraform {
  required_providers {
    ozone = {
      source = "hashicorp/ozone"
      version = "1.0.1"
    }
  }
}

provider "ozone" {
    # api_key = "values to be imported from env variables" - required
    # host = "values to be imported from env variables" - required
    # base_path = "the base path to be used" - optional
    # scheme = "http or https (https by default)" - optional
}


data "oz_project" "default_project"{
    provider = ozone
    name = "<add your workspace name>"
}

data "oz_microservice" "hello-world-app" {
    provider = ozone
    name = "<add your microservice name>"
    workspace_id = data.oz_project.default_project.id
}

data "oz_registry" "hello-world-registry" {
    provider = ozone
    name = "<add your registry name>"
    workspace_id = data.oz_project.default_project.id
}

data "oz_repository" "hello-world-repo" {
    provider = ozone
    name = "<add your repository name>"
    workspace_id = data.oz_project.default_project.id
}


data "oz_cluster" "default_cluster"{
    provider = ozone
    name = "<add your cluster name>"
    workspace_id = data.oz_project.default_project.id
}

data "oz_environment" "dev_env"{
    provider = ozone
    name = "add your environment name"
    workspace_id = data.oz_project.default_project.id
}

data "oz_variable" "test-var"{
    provider = ozone
    name = "add your variable name"
    workspace_id = data.oz_project.default_project.id
}

data "oz_pipeline" "gitpull-pipeline" {
  provider = ozone
  name = "<add your pipeline name>"
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
      image_tag_config = {
        pick_latest_branch_commit = false
      }
      resource_id = ""
    },
    {
      # add more params
    },
  ]  
}

output "pipelune2_output" {
  value = oz_pipelinerun.pipeline
}


output "pipelune2_output" {
  value = oz_pipelinerun.pipeline
}
```