/*
Qovery API

Testing ProjectMainCallsAPIService

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

func Test_qovery_ProjectMainCallsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectMainCallsAPIService DeleteProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		httpRes, err := apiClient.ProjectMainCallsAPI.DeleteProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectMainCallsAPIService EditProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectMainCallsAPI.EditProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectMainCallsAPIService GetProject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectMainCallsAPI.GetProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
