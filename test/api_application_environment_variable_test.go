/*
Qovery API

Testing ApplicationEnvironmentVariableAPIService

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

func Test_qovery_ApplicationEnvironmentVariableAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationEnvironmentVariableAPIService CreateApplicationEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationEnvironmentVariableAPI.CreateApplicationEnvironmentVariable(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationEnvironmentVariableAPIService CreateApplicationEnvironmentVariableAlias", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string
		var environmentVariableId string

		resp, httpRes, err := apiClient.ApplicationEnvironmentVariableAPI.CreateApplicationEnvironmentVariableAlias(context.Background(), applicationId, environmentVariableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationEnvironmentVariableAPIService CreateApplicationEnvironmentVariableOverride", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string
		var environmentVariableId string

		resp, httpRes, err := apiClient.ApplicationEnvironmentVariableAPI.CreateApplicationEnvironmentVariableOverride(context.Background(), applicationId, environmentVariableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationEnvironmentVariableAPIService DeleteApplicationEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string
		var environmentVariableId string

		httpRes, err := apiClient.ApplicationEnvironmentVariableAPI.DeleteApplicationEnvironmentVariable(context.Background(), applicationId, environmentVariableId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationEnvironmentVariableAPIService EditApplicationEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string
		var environmentVariableId string

		resp, httpRes, err := apiClient.ApplicationEnvironmentVariableAPI.EditApplicationEnvironmentVariable(context.Background(), applicationId, environmentVariableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationEnvironmentVariableAPIService ImportEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationEnvironmentVariableAPI.ImportEnvironmentVariable(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationEnvironmentVariableAPIService ListApplicationEnvironmentVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationEnvironmentVariableAPI.ListApplicationEnvironmentVariable(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}