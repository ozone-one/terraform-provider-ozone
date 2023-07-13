---
page_title: "oz_releaserun user guide - ozone"
subcategory: ""
description: |-
  documentation on how to use ozone releaserun provider
---

# Usage Guide

To trigger a release run on ozone using terraform, we use the `oz_releaserun` resource with a set of required inputs.
given below is a sample terraform document which uses `ozone` provider to trigger a release run on ozone.

The expected pipeline inputs can be retrieved via the `oz_release` data source.
The release triggers are mapped from the ozone portal, but can also be passed from the terraform.
here is an example how to extract the pipeline params release.

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
    name = "default"
}

data "oz_release" "first_release" {
  provider = ozone
  name = "ozone-ctl-deploy"
  workspace_id = data.oz_project.default_project.id
}
  
output "step_output" {
  value = data.oz_release.first_release.steps[*].pipeline
}
```
apply the above terraform output to get the pipeline parameters and use them to pass along with its expected step pipeline

```bash
data "oz_project" "default_project" {
    id   = "<project_id>"
    name = "default"
}

# data.oz_release.first_release:
data "oz_release" "first_release" {
    account_id      = "<account_id>"
    created_by_name = "<cread username>"
    id              = "<release id>"
    name            = "ozone-ctl-deploy"
    steps           = [
        {
            description    = ""
            jira_approval  = []
            name           = "build-ozone-ctl-api-0f09af04"
            pipeline       = [
                {
                    application_id = "<app_id>"
                    cluster_id     = "<cluster_id>"
                    env_id         = ""
                    pipeline_id    = "<pipeline_id>"
                    timeout        = 1
                    trigger_params = [
                        {
                            default          = ""
                            description      = <<-EOT
                                Repo branch
                            EOT
                            image_tag_config = [
                                {
                                    pick_latest_branch_commit = false
                                },
                            ]
                            name             = "REPO_BRANCH"
                            required         = true
                            resource_id      = ""
                            type             = 14
                            type_name        = "Repository Branch"
                            value            = "main"
                        },
                    ]
                },
            ]
            run_after      = [
                "ozone-snyk-repo-scan-1bb87ea5",
            ]
            slack_approval = []
            type_id        = 1
            workspaces     = [
                {
                    items = []
                },
            ]
        },
        {
            description    = "Run Snyk repository scans on any git repository synced to Ozone"
            jira_approval  = []
            name           = "ozone-snyk-repo-scan-1bb87ea5"
            pipeline       = [
                {
                    application_id = "<app_id>"
                    cluster_id     = "<cluster_id>"
                    env_id         = ""
                    pipeline_id    = "<pipeline_id>"
                    timeout        = 1
                    trigger_params = [
                        {
                            default          = "test"
                            description      = ""
                            image_tag_config = []
                            name             = "command"
                            required         = false
                            resource_id      = ""
                            type             = 1
                            type_name        = "Text"
                            value            = ""
                        },
                        {
                            default          = ""
                            description      = ""
                            image_tag_config = []
                            name             = "args"
                            required         = true
                            resource_id      = ""
                            type             = 1
                            type_name        = "Text"
                            value            = ""
                        },
                        {
                            default          = ""
                            description      = "The token issued by snyk's API mechanism"
                            image_tag_config = []
                            name             = "SNYK_TOKEN"
                            required         = true
                            resource_id      = ""
                            type             = 3100
                            type_name        = "Snyk Token"
                            value            = ""
                        },
                        {
                            default          = ""
                            description      = "The branch of the repository that Ozone should pull for the repository scan"
                            image_tag_config = [
                                {
                                    pick_latest_branch_commit = false
                                },
                            ]
                            name             = "REPO_BRANCH"
                            required         = true
                            resource_id      = ""
                            type             = 14
                            type_name        = "Repository Branch"
                            value            = "main"
                        },
                    ]
                },
            ]
            run_after      = []
            slack_approval = []
            type_id        = 1
            workspaces     = [
                {
                    items = []
                },
            ]
        },
    ]
    updated_by_name = "<member_name>"
    workspace_id    = "<project_id>"
    workspaces      = [
        {
            items = []
        },
    ]

    last_release_run {}
}


Outputs:

step_output = [
    [
        {
            application_id = "<app_id>"
            cluster_id     = "<cluster_id>"
            env_id         = ""
            # copy this pipeline id and add it in the release runs resource along with the params
            # if the below trigger_params are to be overriden or passed
            pipeline_id    = "<pipeline_id>"
            timeout        = 1
            trigger_params = [
                {
                    default          = ""
                    description      = <<-EOT
                        Repo branch
                    EOT
                    image_tag_config = [
                        {
                            pick_latest_branch_commit = false
                        },
                    ]
                    name             = "REPO_BRANCH"
                    required         = true
                    resource_id      = ""
                    type             = 14
                    type_name        = "Repository Branch"
                    value            = "main"
                },
            ]
        },
    ],
    [
        {
            application_id = "<app_id>"
            cluster_id     = "<cluster_id>"
            env_id         = ""
            # copy this pipeline id and add it in the release runs resource along with the params
            # if the below trigger_params are to be overriden or passed
            pipeline_id    = "<pipeline_id>"
            timeout        = 1
            trigger_params = [
                {
                    default          = "test"
                    description      = ""
                    image_tag_config = []
                    name             = "command"
                    required         = false
                    resource_id      = ""
                    type             = 1
                    type_name        = "Text"
                    value            = ""
                },
                {
                    default          = ""
                    description      = ""
                    image_tag_config = []
                    name             = "args"
                    required         = true
                    resource_id      = ""
                    type             = 1
                    type_name        = "Text"
                    value            = ""
                },
                {
                    default          = ""
                    description      = "The token issued by snyk's API mechanism"
                    image_tag_config = []
                    name             = "SNYK_TOKEN"
                    required         = true
                    resource_id      = ""
                    type             = 3100
                    type_name        = "Snyk Token"
                    value            = ""
                },
                {
                    default          = ""
                    description      = "The branch of the repository that Ozone should pull for the repository scan"
                    image_tag_config = [
                        {
                            pick_latest_branch_commit = false
                        },
                    ]
                    name             = "REPO_BRANCH"
                    required         = true
                    resource_id      = ""
                    type             = 14
                    type_name        = "Repository Branch"
                    value            = "main"
                },
            ]
        },
    ],
]
```


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
    name = "default"
}

data "oz_release" "first_release" {
  provider = ozone
  name = "ozone-ctl-deploy"
  workspace_id = data.oz_project.default_project.id
}

resource "oz_releaserun" "release" {
  provider = ozone

    name = "releaserun"
    workspace_id = data.oz_project.default_project.id
    release_id = data.oz_release.first_release.id
    params = [
      {
        # copy the pipeline id and add desired params which you want to override
        pipeline_id = "<pipeline_id>"
        params = [
         {
          default = ""
          description = "Repo branch"
          image_tag_config = [{
            pick_latest_branch_commit = false
          }]
          name = "REPO_BRANCH"
          required = true
          resource_id = ""
          type = 14
          type_name = "Repository Branch"
          value = "feature/OZN-5416"
        },
        ]
      },
      {
        # copy the pipeline id and add desired params which you want to override
        pipeline_id = "<pipeline_id>"
        params = [{
          default = "test"
          description = ""
          image_tag_config = toset([])
          name = "command"
          required = false
          resource_id = ""
          type = 1
          type_name = "Text"
          value = "sbom"
        },
        {
          default = ""
          description = ""
          image_tag_config = toset([])
          name = "args"
          required = true
          resource_id = ""
          type = 1
          type_name = "Text"
          value = "<desired args>"
        },
        ]
      }
    ]

}

output "release_output" {
  value = oz_releaserun.release
}

```