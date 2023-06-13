package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the DynamicResource resource defined in the Terraform configuration
func DynamicResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"cluster_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"group": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"kind": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"misc_data": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
			ForceNew: true,
		},

		"object": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying DynamicResource resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDynamicResourceResourceData(d *schema.ResourceData, m *models.DynamicResource) {
	d.Set("account_id", m.AccountID)
	d.Set("cluster_id", m.ClusterID)
	d.Set("group", m.Group)
	d.Set("id", m.ID)
	d.Set("kind", m.Kind)
	d.Set("misc_data", m.MiscData)
	d.Set("object", m.Object)
}

// Iterate throught and update the DynamicResource resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDynamicResourceSubResourceData(m []*models.DynamicResource) (d []*map[string]interface{}) {
	for _, dynamicResource := range m {
		if dynamicResource != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = dynamicResource.AccountID
			properties["cluster_id"] = dynamicResource.ClusterID
			properties["group"] = dynamicResource.Group
			properties["id"] = dynamicResource.ID
			properties["kind"] = dynamicResource.Kind
			properties["misc_data"] = dynamicResource.MiscData
			properties["object"] = dynamicResource.Object
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DynamicResource resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DynamicResourceModel(d *schema.ResourceData) *models.DynamicResource {
	accountID := d.Get("account_id").(string)
	clusterID := d.Get("cluster_id").(string)
	group := d.Get("group").(string)
	id, _ := d.Get("id").(string)
	kind := d.Get("kind").(string)
	miscData := d.Get("misc_data")
	object := d.Get("object")

	return &models.DynamicResource{
		AccountID: accountID,
		ClusterID: clusterID,
		Group:     group,
		ID:        id,
		Kind:      kind,
		MiscData:  miscData,
		Object:    object,
	}
}

// Retrieve property field names for updating the DynamicResource resource
func GetDynamicResourcePropertyFields() (t []string) {
	return []string{
		"account_id",
		"cluster_id",
		"group",
		"id",
		"kind",
		"misc_data",
		"object",
	}
}
