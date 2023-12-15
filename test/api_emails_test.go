/*
Elastic Email REST API

Testing EmailsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ElasticEmail

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/elasticemail/elasticemail-go"
)

func Test_ElasticEmail_EmailsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EmailsApiService EmailsByMsgidViewGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var msgid string

		resp, httpRes, err := apiClient.EmailsApi.EmailsByMsgidViewGet(context.Background(), msgid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailsApiService EmailsMergefilePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailsApi.EmailsMergefilePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailsApiService EmailsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailsApi.EmailsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailsApiService EmailsTransactionalPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailsApi.EmailsTransactionalPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
