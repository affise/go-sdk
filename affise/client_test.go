package affise_test

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

type testEnv struct {
	Server     *httptest.Server
	Mux        *http.ServeMux
	Client     *affise.Client
	Ctx        context.Context
	FixtureDir string
}

func newTestEnv(t *testing.T) testEnv {
	t.Helper()

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client, err := affise.NewClient(
		affise.WithBaseURL(server.URL),
		affise.WithAdminURL(server.URL),
	)
	require.NoError(t, err)

	return testEnv{
		Server:     server,
		Mux:        mux,
		Client:     client,
		Ctx:        context.Background(),
		FixtureDir: "../test/testdata",
	}
}

func (env *testEnv) teardown() {
	env.Server.Close()
	env.Server = nil
	env.Mux = nil
	env.Client = nil
	env.FixtureDir = ""
}

func (env *testEnv) mockHandle(t *testing.T, fixture, method, path string, statusCode int) {
	t.Helper()

	f := filepath.Join(env.FixtureDir, fixture)
	data, err := ioutil.ReadFile(f)
	require.NoError(t, err)

	env.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		require.Equal(t, method, r.Method)

		_, err := io.Copy(w, bytes.NewBuffer(data))
		require.NoError(t, err)
	})
}

func (env *testEnv) testPermissions() *affise.Permissions {
	f := filepath.Join(env.FixtureDir, "_permissions.json")
	data, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	return &affise.Permissions{data}
}
