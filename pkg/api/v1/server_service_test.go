package hollow_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	hollow "go.metalkube.net/hollow/pkg/api/v1"
)

func TestServerServiceCreate(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		srv := hollow.Server{UUID: uuid.New(), FacilityCode: "Test1"}
		jsonResponse := json.RawMessage([]byte(`{"message": "resource created", "slug":"00000000-0000-0000-0000-000000001234"}`))

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.Server.Create(ctx, srv)
		if !expectError {
			assert.Equal(t, "00000000-0000-0000-0000-000000001234", res.String())
		}

		return err
	})
}

func TestServerServiceDelete(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		jsonResponse := json.RawMessage([]byte(`{"message": "resource deleted"}`))
		c := mockClient(string(jsonResponse), respCode)
		_, err := c.Server.Delete(ctx, hollow.Server{UUID: uuid.New()})

		return err
	})
}
func TestServerServiceGet(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		srv := hollow.Server{UUID: uuid.New(), FacilityCode: "Test1"}
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Record: srv})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.Server.Get(ctx, srv.UUID)
		if !expectError {
			assert.Equal(t, srv.UUID, res.UUID)
			assert.Equal(t, srv.FacilityCode, res.FacilityCode)
		}

		return err
	})
}

func TestServerServiceList(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		srv := []hollow.Server{{UUID: uuid.New(), FacilityCode: "Test1"}}
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Records: srv})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.Server.List(ctx, nil)
		if !expectError {
			assert.ElementsMatch(t, srv, res)
		}

		return err
	})
}

func TestServerServiceUpdate(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Message: "resource updated"})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		_, err = c.Server.Update(ctx, uuid.UUID{}, hollow.Server{Name: "new-name"})

		return err
	})
}

func TestServerServiceCreateAttributes(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		attr := hollow.Attributes{Namespace: "unit-test", Data: json.RawMessage([]byte(`{"test":"unit"}`))}
		jsonResponse := json.RawMessage([]byte(`{"message": "resource created"}`))

		c := mockClient(string(jsonResponse), respCode)
		_, err := c.Server.CreateAttributes(ctx, uuid.New(), attr)

		return err
	})
}
func TestServerServiceGetAttributes(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		attr := &hollow.Attributes{Namespace: "unit-test", Data: json.RawMessage([]byte(`{"test":"unit"}`))}
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Record: attr})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.Server.GetAttributes(ctx, uuid.UUID{}, "unit-test")
		if !expectError {
			assert.Equal(t, attr, res)
		}

		return err
	})
}

func TestServerServiceDeleteAttributes(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Message: "resource deleted"})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		_, err = c.Server.DeleteAttributes(ctx, uuid.UUID{}, "unit-test")

		return err
	})
}

func TestServerServiceListAttributes(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		attrs := []hollow.Attributes{{Namespace: "unit-test", Data: json.RawMessage([]byte(`{"test":"unit"}`))}}
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Records: attrs})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.Server.ListAttributes(ctx, uuid.UUID{}, nil)
		if !expectError {
			assert.ElementsMatch(t, attrs, res)
		}

		return err
	})
}

func TestServerServiceUpdateAttributes(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Message: "resource updated"})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		_, err = c.Server.UpdateAttributes(ctx, uuid.UUID{}, "unit-test", json.RawMessage([]byte(`{"test":"unit"}`)))

		return err
	})
}

func TestServerServiceListComponents(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		sc := []hollow.ServerComponent{{Name: "unit-test", Serial: "1234"}}
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Records: sc})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.Server.ListComponents(ctx, uuid.UUID{}, nil)
		if !expectError {
			assert.ElementsMatch(t, sc, res)
		}

		return err
	})
}

func TestServerServiceVersionedAttributeCreate(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		va := hollow.VersionedAttributes{Namespace: "unit-test", Data: json.RawMessage([]byte(`{"test":"unit"}`))}
		jsonResponse := json.RawMessage([]byte(`{"message": "resource created", "slug":"the-namespace"}`))

		c := mockClient(string(jsonResponse), respCode)
		resp, err := c.Server.CreateVersionedAttributes(ctx, uuid.New(), va)
		if !expectError {
			assert.Equal(t, "the-namespace", resp.Slug)
		}

		return err
	})
}

func TestServerServiceGetVersionedAttributess(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		va := []hollow.VersionedAttributes{{Namespace: "test", Data: json.RawMessage([]byte(`{}`))}}
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Records: va})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.Server.GetVersionedAttributes(ctx, uuid.New(), "namespace")
		if !expectError {
			assert.ElementsMatch(t, va, res)
		}

		return err
	})
}

func TestServerServiceListVersionedAttributess(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		va := []hollow.VersionedAttributes{{Namespace: "test", Data: json.RawMessage([]byte(`{}`))}}
		jsonResponse, err := json.Marshal(hollow.ServerResponse{Records: va})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.Server.ListVersionedAttributes(ctx, uuid.New())
		if !expectError {
			assert.ElementsMatch(t, va, res)
		}

		return err
	})
}
