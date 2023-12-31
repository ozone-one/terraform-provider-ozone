---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oz_releaserun Resource - ozone"
subcategory: ""
description: |-
  
---

# oz_releaserun (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `params` (List of Object) (see [below for nested schema](#nestedatt--params))
- `release_id` (String)
- `workspace_id` (String)

### Optional

- `account_id` (String)
- `app_release_id` (String)
- `created_at` (String)
- `created_by_id` (String)
- `created_by_name` (String)
- `deleted_at` (String)
- `name` (String)
- `release` (List of Object) (see [below for nested schema](#nestedatt--release))
- `trigger_resource_id` (String)
- `trigger_resource_kind` (String)
- `uid` (String)
- `updated_at` (String)
- `updated_by_id` (String)
- `updated_by_name` (String)
- `workspaces` (List of Object) (see [below for nested schema](#nestedatt--workspaces))

### Read-Only

- `id` (String) The ID of this resource.
- `status` (List of Object) (see [below for nested schema](#nestedatt--status))

<a id="nestedatt--params"></a>
### Nested Schema for `params`

Required:

- `params` (List of Object) (see [below for nested schema](#nestedobjatt--params--params))
- `pipeline_id` (String)

<a id="nestedobjatt--params--params"></a>
### Nested Schema for `params.params`

Required:

- `default` (String)
- `description` (String)
- `image_tag_config` (Set of Object) (see [below for nested schema](#nestedobjatt--params--params--image_tag_config))
- `name` (String)
- `required` (Boolean)
- `resource_id` (String)
- `type` (Number)
- `type_name` (String)
- `value` (String)

<a id="nestedobjatt--params--params--image_tag_config"></a>
### Nested Schema for `params.params.image_tag_config`

Required:

- `pick_latest_branch_commit` (Boolean)




<a id="nestedatt--release"></a>
### Nested Schema for `release`

Optional:

- `name` (String)


<a id="nestedatt--workspaces"></a>
### Nested Schema for `workspaces`

Optional:

- `items` (List of Object) (see [below for nested schema](#nestedobjatt--workspaces--items))

<a id="nestedobjatt--workspaces--items"></a>
### Nested Schema for `workspaces.items`

Optional:




<a id="nestedatt--status"></a>
### Nested Schema for `status`

Read-Only:

- `completion_time` (String)
- `created_time` (String)
- `message` (String)
- `status` (String)
- `updated_time` (String)
