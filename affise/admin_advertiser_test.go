package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestAdminAdvertiserService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "5b5f415035752723008b456a"
			fixture = "3.0.admin.advertiser.{id}@get.json"
			method  = "GET"
			path    = fmt.Sprintf("/3.0/admin/advertiser/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminAdvertiser.Get(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v.SubAccounts) == 2)
	})

	t.Run("List", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.advertisers@get.json"
			method  = "GET"
			path    = "/3.0/admin/advertisers"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAdvertiserListOpts{Limit: 1, Order: "title"}
		v, resp, err := env.Client.AdminAdvertiser.List(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.advertiser@post.json"
			method  = "POST"
			path    = "/3.0/admin/advertiser"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAdvertiserCreateOpts{
			Title:   "MyTitle",
			Contact: "ThePerson",
			Skype:   "MySkype",
			Manager: "5747f68c3b7d9be4018b4570",
		}
		v, resp, err := env.Client.AdminAdvertiser.Create(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.ManagerObj.ID == opts.Manager)
	})

	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "59490d317e28febe1e8b456c"
			fixture = "3.0.admin.advertiser.{id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/advertiser/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAdvertiserUpdateOpts{
			Title:   "MyTitle2",
			Contact: "ThePerson2",
			Skype:   "MySkype2",
			Manager: "5747f68c3b7d9be4018b4570",
		}
		v, resp, err := env.Client.AdminAdvertiser.Update(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.ManagerObj.ID == opts.Manager)
	})

	t.Run("SendPassword", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "59490d317e28febe1e8b456c"
			fixture = "3.0.admin.advertiser.{id}.sendpass@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/advertiser/%s/sendpass", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		resp, err := env.Client.AdminAdvertiser.SendPassword(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("EnableAffiliate", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.advertiser.enable-affiliate@post.json"
			method  = "POST"
			path    = "/3.0/admin/advertiser/enable-affiliate"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAdvertiserEnableAffiliateOpts{
			AdvertisersID: []string{"56fce8ab3b7d9b95588b4568&pid=610"},
		}
		resp, err := env.Client.AdminAdvertiser.EnableAffiliate(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("DisableAffiliate", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.advertiser.disable-affiliate@post.json"
			method  = "POST"
			path    = "/3.0/admin/advertiser/disable-affiliate"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAdvertiserDisableAffiliateOpts{
			AdvertisersID: []string{"56fce8ab3b7d9b95588b4568&pid=610"},
		}
		resp, err := env.Client.AdminAdvertiser.DisableAffiliate(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})
}
