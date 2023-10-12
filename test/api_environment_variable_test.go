/*
Qovery API

Testing EnvironmentVariableAPIService

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

func Test_qovery_EnvironmentVariableAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnvironmentVariableAPIService CreateEnvironmentEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariable(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentVariableAPIService CreateEnvironmentEnvironmentVariableAlias", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string
		var environmentVariableId string

		resp, httpRes, err := apiClient.EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariableAlias(context.Background(), environmentId, environmentVariableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentVariableAPIService CreateEnvironmentEnvironmentVariableOverride", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string
		var environmentVariableId string

		resp, httpRes, err := apiClient.EnvironmentVariableAPI.CreateEnvironmentEnvironmentVariableOverride(context.Background(), environmentId, environmentVariableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentVariableAPIService DeleteEnvironmentEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string
		var environmentVariableId string

		httpRes, err := apiClient.EnvironmentVariableAPI.DeleteEnvironmentEnvironmentVariable(context.Background(), environmentId, environmentVariableId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentVariableAPIService EditEnvironmentEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string
		var environmentVariableId string

		resp, httpRes, err := apiClient.EnvironmentVariableAPI.EditEnvironmentEnvironmentVariable(context.Background(), environmentId, environmentVariableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentVariableAPIService ListEnvironmentEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.EnvironmentVariableAPI.ListEnvironmentEnvironmentVariable(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}