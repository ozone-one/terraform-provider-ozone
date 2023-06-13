package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the MongoObjectIDStruct resource defined in the Terraform configuration
func MongoObjectIDStructSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// Update the underlying MongoObjectIDStruct resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMongoObjectIDStructResourceData(d *schema.ResourceData, m *models.MongoObjectIDStruct) {
	d.Set("id", m.ID)
}

// Iterate throught and update the MongoObjectIDStruct resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMongoObjectIDStructSubResourceData(m []*models.MongoObjectIDStruct) (d []*map[string]interface{}) {
	for _, mongoObjectIdStruct := range m {
		if mongoObjectIdStruct != nil {
			properties := make(map[string]interface{})
			properties["id"] = mongoObjectIdStruct.ID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MongoObjectIDStruct resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MongoObjectIDStructModel(d *schema.ResourceData) *models.MongoObjectIDStruct {
	id, _ := d.Get("id").(string)

	return &models.MongoObjectIDStruct{
		ID: id,
	}
}

// Retrieve property field names for updating the MongoObjectIDStruct resource
func GetMongoObjectIDStructPropertyFields() (t []string) {
	return []string{
		"id",
	}
}
