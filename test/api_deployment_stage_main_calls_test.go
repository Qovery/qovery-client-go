/*
Qovery API

Testing DeploymentStageMainCallsAPIService

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

func Test_qovery_DeploymentStageMainCallsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeploymentStageMainCallsAPIService AttachServiceToDeploymentStage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deploymentStageId string
		var serviceId string

		resp, httpRes, err := apiClient.DeploymentStageMainCallsAPI.AttachServiceToDeploymentStage(context.Background(), deploymentStageId, serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentStageMainCallsAPIService CreateEnvironmentDeploymentStage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.DeploymentStageMainCallsAPI.CreateEnvironmentDeploymentStage(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentStageMainCallsAPIService DeleteDeploymentStage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deploymentStageId string

		httpRes, err := apiClient.DeploymentStageMainCallsAPI.DeleteDeploymentStage(context.Background(), deploymentStageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentStageMainCallsAPIService EditDeploymentStage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deploymentStageId string

		resp, httpRes, err := apiClient.DeploymentStageMainCallsAPI.EditDeploymentStage(context.Background(), deploymentStageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentStageMainCallsAPIService GetDeploymentStage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deploymentStageId string

		resp, httpRes, err := apiClient.DeploymentStageMainCallsAPI.GetDeploymentStage(context.Background(), deploymentStageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentStageMainCallsAPIService GetServiceDeploymentStage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.DeploymentStageMainCallsAPI.GetServiceDeploymentStage(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentStageMainCallsAPIService ListEnvironmentDeploymentStage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.DeploymentStageMainCallsAPI.ListEnvironmentDeploymentStage(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentStageMainCallsAPIService MoveAfterDeploymentStage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deploymentStageId string
		var stageId string

		resp, httpRes, err := apiClient.DeploymentStageMainCallsAPI.MoveAfterDeploymentStage(context.Background(), deploymentStageId, stageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentStageMainCallsAPIService MoveBeforeDeploymentStage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var deploymentStageId string
		var stageId string

		resp, httpRes, err := apiClient.DeploymentStageMainCallsAPI.MoveBeforeDeploymentStage(context.Background(), deploymentStageId, stageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
