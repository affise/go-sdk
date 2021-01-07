package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestAdminUserService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("List", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.users@get.json"
			method  = "GET"
			path    = "/3.0/admin/users"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminUserListOpts{}
		v, resp, err := env.Client.AdminUser.List(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "5f515aefa1ceda82eed06518"
			fixture = "3.0.admin.user.{id}@get.json"
			method  = "GET"
			path    = fmt.Sprintf("/3.0/admin/user/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminUser.Get(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.ID == id)
	})

	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.user@post.json"
			method  = "POST"
			path    = "/3.0/admin/user"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminUserCreateOpts{
			Email:     "user@affise.com",
			Password:  "123456",
			FirstName: "User",
			LastName:  "Affise",
			Roles:     []string{affise.RoleAdmin},
		}
		v, resp, err := env.Client.AdminUser.Create(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, v.FirstName, opts.FirstName)
	})

	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "594927bd7e28fe1c4a8b4569"
			fixture = "3.0.admin.user.{id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/user/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminUserUpdateOpts{
			Email:     "user@affise.com",
			Password:  "123456",
			FirstName: "User2",
			LastName:  "Affise2",
			Roles:     []string{affise.RoleAdmin},
		}
		v, resp, err := env.Client.AdminUser.Update(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, v.FirstName, opts.FirstName)
	})

	t.Run("ChangeAPIKey", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "594927bd7e28fe1c4a8b4569"
			fixture = "3.0.admin.user.api_key.{id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/user/api_key/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminUser.ChangeAPIKey(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, v.ID, id)
	})

	t.Run("ChangePassword", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "594927bd7e28fe1c4a8b4569"
			fixture = "3.0.admin.user.{id}.password@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/user/%s/password", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminUserChangePasswordOpts{
			Password: "c740955e768795098c8b91ef40ec008526f3f884",
		}
		v, resp, err := env.Client.AdminUser.ChangePassword(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, v.Password, opts.Password)
	})

	t.Run("UpdatePermissions", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "594927bd7e28fe1c4a8b4569"
			fixture = "3.1.user.{id}.permissions@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.1/user/%s/permissions", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		perm := env.testPermissions()
		opts := &affise.AdminUserUpdatePermissionsOpts{Permissions: perm}
		v, resp, err := env.Client.AdminUser.UpdatePermissions(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotEmpty(t, v)
	})
}
