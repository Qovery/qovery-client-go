/*
Qovery API

Testing ApplicationsAPIService

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

func Test_qovery_ApplicationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationsAPIService CloneApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationsAPI.CloneApplication(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService CreateApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.ApplicationsAPI.CreateApplication(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetDefaultApplicationAdvancedSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationsAPI.GetDefaultApplicationAdvancedSettings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetEnvironmentApplicationStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetEnvironmentApplicationStatus(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService ListApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentId string

		resp, httpRes, err := apiClient.ApplicationsAPI.ListApplication(context.Background(), environmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
