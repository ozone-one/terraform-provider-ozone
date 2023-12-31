---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oz_pipeline Data Source - ozone"
subcategory: ""
description: |-
  
---

# oz_pipeline (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `workspace_id` (String)

### Optional

- `params` (List of Object) (see [below for nested schema](#nestedatt--params))

### Read-Only

- `account_id` (String)
- `created_at` (String)
- `created_by` (String)
- `deleted_at` (String)
- `id` (String) The ID of this resource.
- `updated_at` (String)
- `updated_by` (String)

<a id="nestedatt--params"></a>
### Nested Schema for `params`

Optional:

- `default` (String)
- `description` (String)
- `image_tag_config` (Set of Object) (see [below for nested schema](#nestedobjatt--params--image_tag_config))
- `name` (String)
- `required` (Boolean)
- `resource_id` (String)
- `type` (Number)
- `type_name` (String)
- `value` (String)

<a id="nestedobjatt--params--image_tag_config"></a>
### Nested Schema for `params.image_tag_config`

Optional:

- `pick_latest_branch_commit` (Boolean)
