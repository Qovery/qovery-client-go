/*
Qovery API

Testing DatabaseDeploymentHistoryAPIService

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

func Test_qovery_DatabaseDeploymentHistoryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DatabaseDeploymentHistoryAPIService ListDatabaseDeploymentHistory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		resp, httpRes, err := apiClient.DatabaseDeploymentHistoryAPI.ListDatabaseDeploymentHistory(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseDeploymentHistoryAPIService ListDatabaseDeploymentHistoryV2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		resp, httpRes, err := apiClient.DatabaseDeploymentHistoryAPI.ListDatabaseDeploymentHistoryV2(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
