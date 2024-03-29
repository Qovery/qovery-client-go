/*
Qovery API

Testing OrganizationWebhookAPIService

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

func Test_qovery_OrganizationWebhookAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrganizationWebhookAPIService CreateOrganizationWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OrganizationWebhookAPI.CreateOrganizationWebhook(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationWebhookAPIService DeleteOrganizationWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var webhookId string

		httpRes, err := apiClient.OrganizationWebhookAPI.DeleteOrganizationWebhook(context.Background(), organizationId, webhookId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationWebhookAPIService EditOrganizationWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var webhookId string

		resp, httpRes, err := apiClient.OrganizationWebhookAPI.EditOrganizationWebhook(context.Background(), organizationId, webhookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationWebhookAPIService GetOrganizationWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string
		var webhookId string

		resp, httpRes, err := apiClient.OrganizationWebhookAPI.GetOrganizationWebhook(context.Background(), organizationId, webhookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrganizationWebhookAPIService ListOrganizationWebHooks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.OrganizationWebhookAPI.ListOrganizationWebHooks(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
