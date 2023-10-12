/*
Qovery API

Testing ContainerActionsAPIService

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

func Test_qovery_ContainerActionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContainerActionsAPIService DeployContainer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var containerId string

		resp, httpRes, err := apiClient.ContainerActionsAPI.DeployContainer(context.Background(), containerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerActionsAPIService RebootContainer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var containerId string

		resp, httpRes, err := apiClient.ContainerActionsAPI.RebootContainer(context.Background(), containerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerActionsAPIService RedeployContainer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var containerId string

		resp, httpRes, err := apiClient.ContainerActionsAPI.RedeployContainer(context.Background(), containerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerActionsAPIService RestartContainer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var containerId string

		resp, httpRes, err := apiClient.ContainerActionsAPI.RestartContainer(context.Background(), containerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerActionsAPIService StopContainer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var containerId string

		resp, httpRes, err := apiClient.ContainerActionsAPI.StopContainer(context.Background(), containerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}