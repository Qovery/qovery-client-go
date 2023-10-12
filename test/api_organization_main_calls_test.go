/*
Qovery API

Testing OrganizationMainCallsAPIService

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

func Test_qovery_OrganizationMainCallsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrganizationMainCallsAPIService CreateOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OrganizationMainCallsAPI.CreateOrganization(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationMainCallsAPIService DeleteOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		httpRes, err := apiClient.OrganizationMainCallsAPI.DeleteOrganization(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationMainCallsAPIService EditOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OrganizationMainCallsAPI.EditOrganization(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationMainCallsAPIService GetOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OrganizationMainCallsAPI.GetOrganization(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationMainCallsAPIService ListOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OrganizationMainCallsAPI.ListOrganization(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationMainCallsAPIService ListOrganizationAvailableRoles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OrganizationMainCallsAPI.ListOrganizationAvailableRoles(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}