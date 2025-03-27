/*
Qovery API

Testing ClustersAPIService

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

func Test_qovery_ClustersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClustersAPIService CreateCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ClustersAPI.CreateCluster(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService DeleteCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		httpRes, err := apiClient.ClustersAPI.DeleteCluster(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService DeployCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.DeployCluster(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService EditCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.EditCluster(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService EditClusterAdvancedSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.EditClusterAdvancedSettings(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService EditClusterKubeconfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		httpRes, err := apiClient.ClustersAPI.EditClusterKubeconfig(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService EditRoutingTable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.EditRoutingTable(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetClusterAdvancedSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.GetClusterAdvancedSettings(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetClusterKubeconfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.GetClusterKubeconfig(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetClusterKubernetesEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.GetClusterKubernetesEvents(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetClusterReadinessStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.GetClusterReadinessStatus(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetClusterStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.GetClusterStatus(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetDefaultClusterAdvancedSettings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ClustersAPI.GetDefaultClusterAdvancedSettings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetInstallationHelmValues", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.GetInstallationHelmValues(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetOrganizationCloudProviderInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.GetOrganizationCloudProviderInfo(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetOrganizationClusterStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ClustersAPI.GetOrganizationClusterStatus(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService GetRoutingTable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.GetRoutingTable(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService ListClusterLogs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.ListClusterLogs(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService ListOrganizationCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ClustersAPI.ListOrganizationCluster(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService LockCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.LockCluster(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService SpecifyClusterCloudProviderInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.SpecifyClusterCloudProviderInfo(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService StopCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.StopCluster(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService UnlockCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		httpRes, err := apiClient.ClustersAPI.UnlockCluster(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService UpdateKarpenterPrivateFargateSubnetIds", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var clusterId string

		httpRes, err := apiClient.ClustersAPI.UpdateKarpenterPrivateFargateSubnetIds(context.Background(), organizationId, clusterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClustersAPIService UpgradeCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterId string

		resp, httpRes, err := apiClient.ClustersAPI.UpgradeCluster(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
