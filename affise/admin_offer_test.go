package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestAdminOfferService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("GetCountOffers", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.offers.count@get.json"
			method  = "GET"
			path    = "/3.0/offers/count"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOffer.GetCountOffers(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v == 8)
	})

	t.Run("CreateOffer", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.offer@post.json"
			method  = "POST"
			path    = "/3.0/admin/offer"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferCreateOfferOpts{
			Title:      "test",
			Advertiser: "573c69a33b7d9b0e638b4576",
			URL:        "http://example.com",
			URLPreview: "http://preview.example.com",
		}
		v, resp, err := env.Client.AdminOffer.CreateOffer(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.ID == 936)
	})

	t.Run("UpdateOffer", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = 936
			fixture = "3.0.admin.offer.{id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/offer/%d", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferUpdateOfferOpts{Title: "test_edit"}
		v, resp, err := env.Client.AdminOffer.UpdateOffer(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, opts.Title, v.Title)
	})

	t.Run("DeleteOffer", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.offer.delete@post.json"
			method  = "POST"
			path    = "/3.0/admin/offer/delete"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferDeleteOfferOpts{OfferID: []int{936}}
		resp, err := env.Client.AdminOffer.DeleteOffer(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("ListSources", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.offer.sources@get.json"
			method  = "GET"
			path    = "/3.0/admin/offer/sources"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOffer.ListSources(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("CreateSource", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.offer.source@post.json"
			method  = "POST"
			path    = "/3.0/admin/offer/source"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferCreateSourceOpts{
			TitleLang: map[string]string{
				"en": "api-test-en-3",
				"ru": "api-test-ru-3",
				"es": "api-test-es-2",
				"ka": "api-test-ka-2",
				"vi": "api-test-vi-3",
			},
		}
		v, resp, err := env.Client.AdminOffer.CreateSource(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, opts.TitleLang, v.TitleLang)
	})

	t.Run("UpdateSource", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "5b7e6d350f0e5a001c7bb4d4"
			fixture = "3.0.admin.offer.source.{id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/offer/source/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferUpdateSourceOpts{
			TitleLang: map[string]string{
				"en": "api-test-en-3",
				"ru": "api-test-ru-3",
				"es": "api-test-es-2",
				"ka": "api-test-ka-2",
				"vi": "api-test-vi-3",
			},
		}
		v, resp, err := env.Client.AdminOffer.UpdateSource(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, opts.TitleLang, v.TitleLang)
	})

	t.Run("DeleteSource", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "5b7e6d350f0e5a001c7bb4d5"
			fixture = "3.0.admin.offer.source.{id}@delete.json"
			method  = "DELETE"
			path    = fmt.Sprintf("/3.0/admin/offer/source/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOffer.DeleteSource(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, id, v.ID)
	})

	t.Run("CreateCategory", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.category@post.json"
			method  = "POST"
			path    = "/3.0/admin/category"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferCreateCategoryOpts{Title: "test_category"}
		v, resp, err := env.Client.AdminOffer.CreateCategory(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, opts.Title, v.Title)
	})

	t.Run("UpdateCategory", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "59440f427e28feff5c8b4567"
			fixture = "3.0.admin.category.{id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/category/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferUpdateCategoryOpts{Title: "test_category2"}
		v, resp, err := env.Client.AdminOffer.UpdateCategory(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, id, v.ID)
	})

	t.Run("EnableAffiliate", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.offer.enable-affiliate@post.json"
			method  = "POST"
			path    = "/3.0/offer/enable-affiliate"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferEnableAffiliateOpts{
			OfferID:     []int{935},
			AffiliateID: 610,
			Notice:      0,
		}
		resp, err := env.Client.AdminOffer.EnableAffiliate(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("DisableAffiliate", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.offer.disable-affiliate@post.json"
			method  = "POST"
			path    = "/3.0/offer/disable-affiliate"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferDisableAffiliateOpts{
			OfferID:     []int{935},
			AffiliateID: 610,
			Notice:      0,
		}
		resp, err := env.Client.AdminOffer.DisableAffiliate(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("MassUpdateOffers", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.offer.mass-update@post.json"
			method  = "POST"
			path    = "/3.0/admin/offer/mass-update"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOfferMassUpdateOffersOpts{
			OfferID: []int{1, 2},
			Status:  "active",
		}
		resp, err := env.Client.AdminOffer.MassUpdateOffers(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("DisableAffiliates", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "5943f7307e28fe9a1f8b456d"
			fixture = "3.0.admin.offer.{id}.disable-affiliates@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/offer/%s/disable-affiliates", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		resp, err := env.Client.AdminOffer.DisableAffiliates(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("DisableOffers", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      uint64 = 10
			fixture        = "3.0.admin.affiliate.{id}.disable-offers@post.json"
			method         = "POST"
			path           = fmt.Sprintf("/3.0/admin/affiliate/%d/disable-offers", id)
			status         = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		resp, err := env.Client.AdminOffer.DisableOffers(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("RemoveOfferCreatives", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "5943f7307e28fe9a1f8b456d"
			fixture = "3.0.admin.offer.{id}.remove-creative@delete.json"
			method  = "DELETE"
			path    = fmt.Sprintf("/3.0/admin/offer/%s/remove-creative", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminRemoveOfferCreativesOpts{Creatives: []int{1, 2}}
		v, resp, err := env.Client.AdminOffer.RemoveOfferCreatives(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, opts.Creatives, v)
	})
}
