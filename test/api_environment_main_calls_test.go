/*
Qovery API

Testing EnvironmentMainCallsAPIService

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

func Test_qovery_EnvironmentMainCallsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnvironmentMainCallsAPIService DeleteEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		httpRes, err := apiClient.EnvironmentMainCallsAPI.DeleteEnvironment(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentMainCallsAPIService EditEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.EnvironmentMainCallsAPI.EditEnvironment(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentMainCallsAPIService GetEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.EnvironmentMainCallsAPI.GetEnvironment(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentMainCallsAPIService GetEnvironmentStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.EnvironmentMainCallsAPI.GetEnvironmentStatus(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentMainCallsAPIService GetEnvironmentStatuses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.EnvironmentMainCallsAPI.GetEnvironmentStatuses(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentMainCallsAPIService GetEnvironmentStatusesWithStages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.EnvironmentMainCallsAPI.GetEnvironmentStatusesWithStages(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentMainCallsAPIService ListDeploymentRequestByEnvironmentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.EnvironmentMainCallsAPI.ListDeploymentRequestByEnvironmentId(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentMainCallsAPIService ListDeploymentRequestByServiceId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.EnvironmentMainCallsAPI.ListDeploymentRequestByServiceId(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentMainCallsAPIService ListServicesByEnvironmentId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.EnvironmentMainCallsAPI.ListServicesByEnvironmentId(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
