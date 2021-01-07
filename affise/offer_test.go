package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestOfferService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("List", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.offers@get.json"
			method  = "GET"
			path    = "/3.0/offers"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.OfferListOpts{IDs: []string{"331"}}
		v, resp, err := env.Client.Offer.List(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = 906
			fixture = "3.0.offer.{id}@get.json"
			method  = "GET"
			path    = fmt.Sprintf("/3.0/offer/%d", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Offer.Get(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, id, v.ID)
		require.True(t, len(v.DescriptionLang) == 5)
		require.True(t, len(v.KPI) == 5)
	})

	t.Run("ListCategories", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.offer.categories@get.json"
			method  = "GET"
			path    = "/3.0/offer/categories"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.OfferListCategoriesOpts{Page: 1}
		v, resp, err := env.Client.Offer.ListCategories(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})
}
