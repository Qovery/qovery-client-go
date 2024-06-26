/*
Qovery API

Testing HelmRepositoriesAPIService

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

func Test_qovery_HelmRepositoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HelmRepositoriesAPIService CreateHelmRepository", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.HelmRepositoriesAPI.CreateHelmRepository(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HelmRepositoriesAPIService DeleteHelmRepository", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var helmRepositoryId string

		httpRes, err := apiClient.HelmRepositoriesAPI.DeleteHelmRepository(context.Background(), organizationId, helmRepositoryId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HelmRepositoriesAPIService EditHelmRepository", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var helmRepositoryId string

		resp, httpRes, err := apiClient.HelmRepositoriesAPI.EditHelmRepository(context.Background(), organizationId, helmRepositoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HelmRepositoriesAPIService GetHelmCharts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var helmRepositoryId string

		resp, httpRes, err := apiClient.HelmRepositoriesAPI.GetHelmCharts(context.Background(), organizationId, helmRepositoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HelmRepositoriesAPIService GetHelmRepository", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var helmRepositoryId string

		resp, httpRes, err := apiClient.HelmRepositoriesAPI.GetHelmRepository(context.Background(), organizationId, helmRepositoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HelmRepositoriesAPIService ListAvailableHelmRepository", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HelmRepositoriesAPI.ListAvailableHelmRepository(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HelmRepositoriesAPIService ListHelmRepository", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.HelmRepositoriesAPI.ListHelmRepository(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
