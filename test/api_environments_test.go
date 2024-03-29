/*
Qovery API

Testing EnvironmentsAPIService

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

func Test_qovery_EnvironmentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnvironmentsAPIService CreateEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsAPIService GetProjectEnvironmentServiceNumber", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.EnvironmentsAPI.GetProjectEnvironmentServiceNumber(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsAPIService GetProjectEnvironmentsStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.EnvironmentsAPI.GetProjectEnvironmentsStatus(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentsAPIService ListEnvironment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.EnvironmentsAPI.ListEnvironment(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
