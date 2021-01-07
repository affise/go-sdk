package affise_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestAdminConversionService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("Edit", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.conversion.edit@post.json"
			method  = "POST"
			path    = "/3.0/admin/conversion/edit"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminConversionEditOpts{IDs: []string{"59359e1d7e28feb7568b456a"}}
		v, resp, err := env.Client.AdminConversion.Edit(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, opts.IDs, v.IDs)
	})

	t.Run("Import", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.conversion.import@post.json"
			method  = "POST"
			path    = "/3.0/admin/conversion/import"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminConversionImportOpts{Offer: 1000, AffiliateID: 500}
		v, resp, err := env.Client.AdminConversion.Import(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, opts.Offer, v.Offer)
		require.Equal(t, opts.AffiliateID, v.AffiliateID)
	})

	t.Run("ImportList", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.conversions.import@post.json"
			method  = "POST"
			path    = "/3.0/admin/conversions/import"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminConversionImportListOpts{
			List: []affise.AdminConversionImportOpts{{Offer: 1000, AffiliateID: 500}},
		}
		v, resp, err := env.Client.AdminConversion.ImportList(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, len(opts.List), len(v))
		require.Equal(t, opts.List[0].Offer, v[0].Offer)
		require.Equal(t, opts.List[0].AffiliateID, v[0].AffiliateID)
	})
}
