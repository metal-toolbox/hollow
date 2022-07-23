package serverservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	serverservice "go.hollow.sh/serverservice/pkg/api/v1"
)

func TestIntegrationServerSecretTypesList(t *testing.T) {
	s := serverTest(t)

	realClientTests(t, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		r, resp, err := s.Client.ListServerSecretTypes(ctx, nil)
		if !expectError {
			require.NoError(t, err)
			assert.Len(t, r, 1)

			assert.EqualValues(t, 1, resp.PageCount)
			assert.EqualValues(t, 1, resp.TotalPages)
			assert.EqualValues(t, 1, resp.TotalRecordCount)
			// We returned everything, so we shouldnt have a next page info
			assert.Nil(t, resp.Links.Next)
			assert.Nil(t, resp.Links.Previous)
		}

		return err
	})
}

func TestIntegrationServerSecretTypesCreate(t *testing.T) {
	ctx := context.TODO()
	s := serverTest(t)

	realClientTests(t, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.CreateServerSecretType(ctx, &serverservice.ServerSecretType{Name: "Test Type"})
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Links.Self)
			assert.Equal(t, "http://test.hollow.com/api/v1/server-secret-types/test-type", resp.Links.Self.Href)
		}

		return err
	})

	s.Client.SetToken(validToken(adminScopes))

	t.Run("creating a duplicate type fails", func(t *testing.T) {
		slug := serverservice.ServerSecretTypeBMC

		// Make sure our server doesn't already have a BMC secret
		_, err := s.Client.CreateServerSecretType(ctx, &serverservice.ServerSecretType{Name: "Test Type", Slug: slug})
		assert.Error(t, err)
		require.Contains(t, err.Error(), "duplicate key")
	})
}
