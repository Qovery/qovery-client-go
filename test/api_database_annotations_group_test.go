/*
Qovery API

Testing DatabaseAnnotationsGroupAPIService

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

func Test_qovery_DatabaseAnnotationsGroupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DatabaseAnnotationsGroupAPIService AddAnnotationsGroupToDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		httpRes, err := apiClient.DatabaseAnnotationsGroupAPI.AddAnnotationsGroupToDatabase(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseAnnotationsGroupAPIService DeleteAnnotationsGroupToDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		httpRes, err := apiClient.DatabaseAnnotationsGroupAPI.DeleteAnnotationsGroupToDatabase(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseAnnotationsGroupAPIService ListDatabaseAnnotationsGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseId string

		resp, httpRes, err := apiClient.DatabaseAnnotationsGroupAPI.ListDatabaseAnnotationsGroup(context.Background(), databaseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
