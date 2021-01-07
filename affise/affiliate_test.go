package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestAffiliateService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("Me", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.partner.me@get.json"
			method  = "GET"
			path    = "/3.1/partner/me"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Affiliate.Me(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotNil(t, v)
	})

	t.Run("ListOffers", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.partner.offers@get.json"
			method  = "GET"
			path    = "/3.0/partner/offers"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AffiliateListOffersOpts{Limit: 1}
		v, resp, err := env.Client.Affiliate.ListOffers(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotNil(t, resp.Meta.Pagination)
		require.True(t, len(v) == 1)
	})

	t.Run("ListLiveOffers", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.partner.live-offers@get.json"
			method  = "GET"
			path    = "/3.0/partner/live-offers"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AffiliateListLiveOffersOpts{Limit: 1}
		v, resp, err := env.Client.Affiliate.ListLiveOffers(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotNil(t, resp.Meta.Pagination)
		require.True(t, len(v) == 1)
	})

	t.Run("ActivationOffer", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.partner.activation.offer@post.json"
			method  = "POST"
			path    = "/3.0/partner/activation/offer"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AffiliateActivationOfferOpts{OfferID: 123, Comment: "test"}
		resp, err := env.Client.Affiliate.ActivationOffer(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotEmpty(t, 1, resp.Meta.Message)
	})

	t.Run("CreatePostback", func(t *testing.T) {
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

		opts := &affise.AffiliateCreatePostbackOpts{
			URL:     "http://affise.com",
			Status:  "by_creating",
			OfferID: 906,
		}
		v, resp, err := env.Client.Affiliate.CreatePostback(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.Status == opts.Status)
	})

	t.Run("UpdatePostback", func(t *testing.T) {
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

		opts := &affise.AffiliateUpdatePostbackOpts{
			URL:    "http://affise.com",
			Status: "confirmed",
		}
		v, resp, err := env.Client.Affiliate.UpdatePostback(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.Status == opts.Status)
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

		v, resp, err := env.Client.Affiliate.DeletePostback(env.Ctx, id)
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

		opts := &affise.AffiliateDeletePostbacksByAffiliatesOpts{IDs: []int{1, 2, 3}}
		resp, err := env.Client.Affiliate.DeletePostbacksByAffiliates(env.Ctx, opts)
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

		opts := &affise.AffiliateDeletePostbacksByOffersOpts{IDs: []int{1, 2, 3}}
		resp, err := env.Client.Affiliate.DeletePostbacksByOffers(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("ListNews", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.news@get.json"
			method  = "GET"
			path    = "/3.0/news"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AffiliateListNewsOpts{Limit: 1}
		v, resp, err := env.Client.Affiliate.ListNews(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("GetNewsByID", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "57a4914f3b7d9bbd358b45b6"
			fixture = "3.0.news.{id}@get.json"
			method  = "GET"
			path    = fmt.Sprintf("/3.0/news/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Affiliate.GetNewsByID(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.ID.ID == id)
	})

	t.Run("ListPixels", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.partner.pixels@get.json"
			method  = "GET"
			path    = "/3.0/partner/pixels"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Affiliate.ListPixels(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("CreatePixel", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.partner.pixel@post.json"
			method  = "POST"
			path    = "/3.0/partner/pixel"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AffiliateCreatePixelOpts{
			OfferID:  906,
			Name:     "test",
			Code:     "<script>test</script>",
			CodeType: "javascript",
		}
		v, resp, err := env.Client.Affiliate.CreatePixel(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.Name == opts.Name)
	})

	t.Run("UpdatePixel", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = 2
			fixture = "3.0.partner.pixel.{id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/partner/pixel/%d", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AffiliateUpdatePixelOpts{
			Name:     "test2",
			Code:     "<script>test</script>",
			CodeType: "javascript",
		}
		v, resp, err := env.Client.Affiliate.UpdatePixel(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.CodeType == opts.CodeType)
	})

	t.Run("DeletePixel", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = 2
			fixture = "3.0.partner.pixel.{id}.remove@delete.json"
			method  = "DELETE"
			path    = fmt.Sprintf("/3.0/partner/pixel/%d/remove", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Affiliate.DeletePixel(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.OfferID == "906")
	})

	t.Run("GetAffiliateBalance", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.balance@get.json"
			method  = "GET"
			path    = "/3.0/balance"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Affiliate.GetAffiliateBalance(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)

		_, ok := v["USD"]
		require.True(t, ok)
	})

	t.Run("GetSmartLinkCategories", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.partner.smartlink.categories@get.json"
			method  = "GET"
			path    = "/3.0/partner/smartlink/categories"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AffiliateGetSmartLinkCategoriesOpts{
			ID: []string{"595e3b547e28fede7b8b456c"},
		}
		v, resp, err := env.Client.Affiliate.GetSmartLinkCategories(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("GetSmartLinkOfferCount", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "595fd4877e28fee8428b459f"
			fixture = "3.0.partner.smartlink.category.{id}.offers-count@get.json"
			method  = "GET"
			path    = fmt.Sprintf("/3.0/partner/smartlink/category/%s/offers-count", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Affiliate.GetSmartLinkOfferCount(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v == 2)
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

		v, resp, err := env.Client.Affiliate.GetReferrals(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)

		_, ok := v[0].Balance["USD"]
		require.True(t, ok)
	})
}
