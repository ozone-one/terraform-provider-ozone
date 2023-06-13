package schemata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStepParams(t *testing.T) {
	step := map[string]interface{}{

		"name":        "build-ozone-ctl-swagger-api-6a006758",
		"type_id":     1,
		"description": "",
		"pipeline": []interface{}{
			map[string]interface{}{
				"pipeline_id":    "6461e79ccc9f9877d347000a",
				"cluster_id":     "642da86cbb3287bdb1a9fda2",
				"application_id": "6461d9ae864db62ed52001e8",
				"env_id":         "642da8c9d3491b96aaae8657",
				"trigger_params": []interface{}{
					map[string]interface{}{
						"name":        "REPO_BRANCH",
						"type":        14,
						"typeName":    "Repository Branch",
						"description": "Repo branch\n",
						"value":       "6461d913cf6d42d6b01ed39b:::main",
						"default":     "",
						"required":    true,
						"index":       0,
					}},
				"timeout": 2,
			},
		},
		"workspaces": map[string]interface{}{
			"items": []interface{}{},
		},
	}
	result := stepParams(step)
	assert.Equal(t, "build-ozone-ctl-swagger-api-6a006758", result.Name)
	assert.Equal(t, "", result.Description)
	assert.Equal(t, "6461e79ccc9f9877d347000a", result.Pipeline.PipelineID)
	assert.Equal(t, "642da86cbb3287bdb1a9fda2", result.Pipeline.ClusterID)
	assert.Equal(t, "6461d9ae864db62ed52001e8", result.Pipeline.ApplicationID)
	assert.Equal(t, "642da8c9d3491b96aaae8657", result.Pipeline.EnvID)
	assert.Equal(t, int64(2), result.Pipeline.Timeout)
	assert.Len(t, result.Pipeline.TriggerParams, 1)
	assert.Equal(t, "REPO_BRANCH", result.Pipeline.TriggerParams[0].Name)
	assert.Equal(t, int64(1), result.TypeID)

}
