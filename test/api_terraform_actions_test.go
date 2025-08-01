/*
Qovery API

Testing TerraformActionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package qovery

import (
	"context"
	openapiclient "github.com/qovery/qovery-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_qovery_TerraformActionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TerraformActionsAPIService DeployTerraform", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var terraformId string

		resp, httpRes, err := apiClient.TerraformActionsAPI.DeployTerraform(context.Background(), terraformId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TerraformActionsAPIService RedeployTerraform", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var terraformId string

		resp, httpRes, err := apiClient.TerraformActionsAPI.RedeployTerraform(context.Background(), terraformId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TerraformActionsAPIService UninstallTerraform", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var terraformId string

		resp, httpRes, err := apiClient.TerraformActionsAPI.UninstallTerraform(context.Background(), terraformId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
