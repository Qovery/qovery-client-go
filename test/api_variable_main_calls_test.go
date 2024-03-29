/*
Qovery API

Testing VariableMainCallsAPIService

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

func Test_qovery_VariableMainCallsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VariableMainCallsAPIService CreateVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VariableMainCallsAPI.CreateVariable(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VariableMainCallsAPIService CreateVariableAlias", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var variableId string

		resp, httpRes, err := apiClient.VariableMainCallsAPI.CreateVariableAlias(context.Background(), variableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VariableMainCallsAPIService CreateVariableOverride", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var variableId string

		resp, httpRes, err := apiClient.VariableMainCallsAPI.CreateVariableOverride(context.Background(), variableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VariableMainCallsAPIService DeleteVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var variableId string

		httpRes, err := apiClient.VariableMainCallsAPI.DeleteVariable(context.Background(), variableId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VariableMainCallsAPIService EditVariable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var variableId string

		resp, httpRes, err := apiClient.VariableMainCallsAPI.EditVariable(context.Background(), variableId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VariableMainCallsAPIService ImportEnvironmentVariables", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VariableMainCallsAPI.ImportEnvironmentVariables(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VariableMainCallsAPIService ListVariables", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VariableMainCallsAPI.ListVariables(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
