package hollow_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.metalkube.net/hollow/internal/db"
	"go.metalkube.net/hollow/internal/hollowserver"
	hollow "go.metalkube.net/hollow/pkg/api/v1"
)

type testServer struct {
	h      http.Handler
	Client *hollow.Client
}

func serverTest(t *testing.T) *testServer {
	db.DatabaseTest(t)

	hs := hollowserver.Server{Logger: zap.NewNop()}
	s := hs.NewServer()

	ts := &testServer{
		h: s.Handler,
	}

	c, err := hollow.NewClient("testToken", "http://test.hollow.com", ts)
	require.NoError(t, err)

	ts.Client = c

	return ts
}

func (s *testServer) Do(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	s.h.ServeHTTP(w, req)

	return w.Result(), nil
}
