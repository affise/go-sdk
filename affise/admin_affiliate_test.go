package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestAdminAffiliateService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      uint64 = 1
			fixture        = "3.0.admin.partner.{id}@get.json"
			method         = "GET"
			path           = fmt.Sprintf("/3.0/admin/partner/%d", id)
			status         = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminAffiliate.Get(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, id, v.ID)
	})

	t.Run("ListPartners", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.partners@get.json"
			method  = "GET"
			path    = "/3.0/admin/partners"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateListPartnersOpts{}
		v, resp, err := env.Client.AdminAffiliate.ListPartners(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)
	})

	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.partner@post.json"
			method  = "POST"
			path    = "/3.0/admin/partner"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateCreateOpts{
			Email:             "affise@gmail.com",
			Password:          "qwerty123456",
			Login:             "ivan.ivanov",
			RefPercent:        "2",
			Notes:             "none",
			Status:            "active",
			ManagerID:         "5cd5530ad596c1c0008b4567",
			CustomFields:      []string{"skype"},
			Ref:               2,
			SubAccount1:       "sub1",
			SubAccount2:       "sub2",
			SubAccount1Except: 0,
			SubAccount2Except: 1,
			PaymentSystems: []affise.PaymentSystemOpts{
				{
					SystemID: 1,
					Currency: "USD",
					Fields: map[string]string{
						"1": "BA731035962466786892",
						"2": "PK83DELLCTnbVB5RMU5TL1X4",
					},
				},
			},
		}
		v, resp, err := env.Client.AdminAffiliate.Create(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.ID == 5)
	})

	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      uint64 = 5
			fixture        = "3.0.admin.partner.{id}@post.json"
			method         = "POST"
			path           = fmt.Sprintf("/3.0/admin/partner/%d", id)
			status         = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateUpdateOpts{
			Password:          "qwerty123456",
			Login:             "ivan.ivanov",
			RefPercent:        "2",
			Notes:             "none",
			Status:            "active",
			ManagerID:         "5cd5530ad596c1c0008b4567",
			CustomFields:      []string{"skype"},
			Ref:               2,
			SubAccount1:       "sub1",
			SubAccount2:       "sub2",
			SubAccount1Except: 0,
			SubAccount2Except: 1,
			PaymentSystems: []affise.PaymentSystemOpts{
				{
					SystemID: 1,
					Currency: "USD",
					Fields: map[string]string{
						"1": "BA731035962466786892",
						"2": "PK83DELLCTnbVB5RMU5TL1X4",
					},
				},
			},
		}
		v, resp, err := env.Client.AdminAffiliate.Update(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, id == v.ID)
	})

	t.Run("MassUpdate", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.partners.mass-update@post.json"
			method  = "POST"
			path    = "/3.0/admin/partners/mass-update"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateMassUpdateOpts{ID: []uint64{11, 12}}
		resp, err := env.Client.AdminAffiliate.MassUpdate(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("ChangePassword", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      uint64 = 10117
			fixture        = "3.0.admin.partner.password.{id}@post.json"
			method         = "POST"
			path           = fmt.Sprintf("/3.0/admin/partner/password/%d", id)
			status         = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminAffiliate.ChangePassword(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, id == v.ID)
	})

	t.Run("AddPostback", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.partner.postback@post.json"
			method  = "POST"
			path    = "/3.0/partner/postback"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateAddPostbackOpts{
			Status:      "by_creating",
			AffiliateID: 610,
			OfferID:     960,
		}
		v, resp, err := env.Client.AdminAffiliate.AddPostback(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, opts.OfferID == v.ID)
	})

	t.Run("EditPostback", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = 960
			fixture = "3.0.partner.postback.{id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/partner/postback/%d", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateEditPostbackOpts{
			Status: "confirmed",
		}
		v, resp, err := env.Client.AdminAffiliate.EditPostback(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, id == v.ID)
	})

	t.Run("DeletePostback", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = 960
			fixture = "3.0.partner.postback.{id}.remove@delete.json"
			method  = "DELETE"
			path    = fmt.Sprintf("/3.0/partner/postback/%d/remove", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminAffiliate.DeletePostback(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, id == v.ID)
	})

	t.Run("DeletePostbacksByAffiliates", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.partner.postbacks.by-affiliates@delete.json"
			method  = "DELETE"
			path    = "/3.0/partner/postbacks/by-affiliates"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateDeletePostbacksByAffiliatesOpts{IDs: []int{1, 2, 3}}
		resp, err := env.Client.AdminAffiliate.DeletePostbacksByAffiliates(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("DeletePostbacksByOffers", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.partner.postbacks.by-offers@delete.json"
			method  = "DELETE"
			path    = "/3.0/partner/postbacks/by-offers"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateDeletePostbacksByOffersOpts{IDs: []int{1, 2, 3}}
		resp, err := env.Client.AdminAffiliate.DeletePostbacksByOffers(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("ListPostbacks", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.postbacks@get.json"
			method  = "GET"
			path    = "/3.0/admin/postbacks"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateListPostbacksOpts{Limit: 3}
		v, resp, err := env.Client.AdminAffiliate.ListPostbacks(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, opts.Limit == len(v))
	})

	t.Run("ChangeAPIKey", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.partner.api_key@post.json"
			method  = "POST"
			path    = "/3.1/partner/api_key"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminAffiliate.ChangeAPIKey(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotEmpty(t, v.APIKey)
	})

	t.Run("UpdateLocale", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      uint64 = 1
			fixture        = "3.0.admin.partner.{id}.locale@post.json"
			method         = "POST"
			path           = fmt.Sprintf("/3.0/admin/partner/%d/locale", id)
			status         = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAffiliateUpdateLocaleOpts{
			Lang:     "en",
			Timezone: "Europe/Minsk",
		}
		resp, err := env.Client.AdminAffiliate.UpdateLocale(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("GetReferrals", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      uint64 = 2
			fixture        = "3.0.admin.partner.{id}.referrals@get.json"
			method         = "GET"
			path           = fmt.Sprintf("/3.0/admin/partner/%d/referrals", id)
			status         = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminAffiliate.GetReferrals(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})
}
