/*
Qovery API

Testing JobMainCallsAPIService

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

func Test_qovery_JobMainCallsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test JobMainCallsAPIService DeleteJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		httpRes, err := apiClient.JobMainCallsAPI.DeleteJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobMainCallsAPIService EditJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.JobMainCallsAPI.EditJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobMainCallsAPIService GetJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.JobMainCallsAPI.GetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobMainCallsAPIService GetJobStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.JobMainCallsAPI.GetJobStatus(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobMainCallsAPIService ListJobCommit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.JobMainCallsAPI.ListJobCommit(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}