package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestAdminPreset_Mock(t *testing.T) {
	t.Parallel()
	t.Run("List", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.presets@get.json"
			method  = "GET"
			path    = "/3.1/presets"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminPresetListOpts{Limit: 1}
		v, resp, err := env.Client.AdminPreset.List(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.presets@post.json"
			method  = "POST"
			path    = "/3.1/presets"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminPresetCreateOpts{
			Name:        "Test affiliate_manager 4",
			Permissions: env.testPermissions(),
			Type:        "affiliate_manager",
		}
		v, resp, err := env.Client.AdminPreset.Create(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotNil(t, v.Permissions)
		require.Equal(t, v.CreatedAt, "2020-09-03T22:59:38Z")
	})

	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "5f51755a535bac2217eb7619"
			fixture = "3.1.presets.{preset_id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.1/presets/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminPresetUpdateOpts{
			Name:        "Test affiliate_manager 4 (update)",
			Permissions: env.testPermissions(),
		}
		v, resp, err := env.Client.AdminPreset.Update(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotNil(t, v.Permissions)
		require.Equal(t, v.UpdatedAt, "2020-09-03T23:16:22Z")
	})

	t.Run("Delete", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "5d4c30c610dd212ea8dc98f5"
			fixture = "3.1.presets.{preset_id}@delete.json"
			method  = "DELETE"
			path    = fmt.Sprintf("/3.1/presets/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		resp, err := env.Client.AdminPreset.Delete(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})
}
