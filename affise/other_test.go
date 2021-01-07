package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestOtherService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("ListISP", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.isp@get.json"
			method  = "GET"
			path    = "/3.1/isp"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.OtherListISPOpts{Country: "KZ"}
		v, resp, err := env.Client.Other.ListISP(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 3)

		want := &affise.ISP{Country: "KZ", Name: "betting office olimp kz llc"}
		require.Equal(t, v[2], want)
	})

	t.Run("ListCountries", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.countries@get.json"
			method  = "GET"
			path    = "/3.1/countries"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Other.ListCountries(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 5)

		want := &affise.Country{Code: "GI", Name: "Gibraltar"}
		require.Equal(t, v[4], want)
	})

	t.Run("ListRegions", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.regions@get.json"
			method  = "GET"
			path    = "/3.1/regions"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.OtherListRegionsOpts{Country: "US"}
		v, resp, err := env.Client.Other.ListRegions(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)

		want := &affise.Region{ID: 1, Name: "Alabama", CountryCode: "US"}
		require.Equal(t, v[1], want)
	})

	t.Run("ListConnectionTypes", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.connection-types@get.json"
			method  = "GET"
			path    = "/3.1/connection-types"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Other.ListConnectionTypes(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 3)

		want := "other"
		require.Equal(t, v[2], want)
	})

	t.Run("ListVendors", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.vendors@get.json"
			method  = "GET"
			path    = "/3.1/vendors"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.OtherListVendorsOpts{Q: "next"}
		v, resp, err := env.Client.Other.ListVendors(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 8)

		want := "NextWolf"
		require.Equal(t, v[7], want)
	})

	t.Run("ListOS", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.oses@get.json"
			method  = "GET"
			path    = "/3.1/oses"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Other.ListOS(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 10)

		key, val := "10", "Apple TV Software"
		require.Equal(t, v[key], val)
	})

	t.Run("ListOSVersions", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			os      = "macOS"
			fixture = "3.1.oses.{os}@get.json"
			method  = "GET"
			path    = fmt.Sprintf("/3.1/oses/%s", os)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Other.ListOSVersions(env.Ctx, os)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 21)

		want := "10.13.6"
		require.Equal(t, v[20], want)
	})
}
