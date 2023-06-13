package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the TemplateBootstrapValues resource defined in the Terraform configuration
func TemplateBootstrapValuesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"deployment_replica": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"dns": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"namespace_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"service_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"service_port": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"service_target_port": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"service_type": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying TemplateBootstrapValues resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTemplateBootstrapValuesResourceData(d *schema.ResourceData, m *models.TemplateBootstrapValues) {
	d.Set("deployment_replica", m.DeploymentReplica)
	d.Set("dns", m.DNS)
	d.Set("namespace_name", m.NamespaceName)
	d.Set("service_name", m.ServiceName)
	d.Set("service_port", m.ServicePort)
	d.Set("service_target_port", m.ServiceTargetPort)
	d.Set("service_type", m.ServiceType)
}

// Iterate throught and update the TemplateBootstrapValues resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTemplateBootstrapValuesSubResourceData(m []*models.TemplateBootstrapValues) (d []*map[string]interface{}) {
	for _, templateBootstrapValues := range m {
		if templateBootstrapValues != nil {
			properties := make(map[string]interface{})
			properties["deployment_replica"] = templateBootstrapValues.DeploymentReplica
			properties["dns"] = templateBootstrapValues.DNS
			properties["namespace_name"] = templateBootstrapValues.NamespaceName
			properties["service_name"] = templateBootstrapValues.ServiceName
			properties["service_port"] = templateBootstrapValues.ServicePort
			properties["service_target_port"] = templateBootstrapValues.ServiceTargetPort
			properties["service_type"] = templateBootstrapValues.ServiceType
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate TemplateBootstrapValues resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TemplateBootstrapValuesModel(d *schema.ResourceData) *models.TemplateBootstrapValues {
	deploymentReplica := int64(d.Get("deployment_replica").(int))
	dns := d.Get("dns").(string)
	namespaceName := d.Get("namespace_name").(string)
	serviceName := d.Get("service_name").(string)
	servicePort := int64(d.Get("service_port").(int))
	serviceTargetPort := int64(d.Get("service_target_port").(int))
	serviceType := d.Get("service_type").(string)

	return &models.TemplateBootstrapValues{
		DeploymentReplica: deploymentReplica,
		DNS:               dns,
		NamespaceName:     namespaceName,
		ServiceName:       serviceName,
		ServicePort:       servicePort,
		ServiceTargetPort: serviceTargetPort,
		ServiceType:       serviceType,
	}
}

// Retrieve property field names for updating the TemplateBootstrapValues resource
func GetTemplateBootstrapValuesPropertyFields() (t []string) {
	return []string{
		"deployment_replica",
		"dns",
		"namespace_name",
		"service_name",
		"service_port",
		"service_target_port",
		"service_type",
	}
}
