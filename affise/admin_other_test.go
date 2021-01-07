package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestAdminOtherService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("ListCities", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.cities@get.json"
			method  = "GET"
			path    = "/3.1/cities"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOtherListCitiesOpts{Country: []string{"DE"}}
		v, resp, err := env.Client.AdminOther.ListCities(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 3)
	})

	t.Run("ListDevices", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.devices@get.json"
			method  = "GET"
			path    = "/3.1/devices"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOther.ListDevices(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 8)
	})

	t.Run("ListBrowsers", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.browsers@get.json"
			method  = "GET"
			path    = "/3.1/browsers"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOther.ListBrowsers(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 7)
	})

	t.Run("ListCurrencies", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.currency@get.json"
			method  = "GET"
			path    = "/3.0/admin/currency"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOtherListCurrenciesOpts{GetOnlyActive: 1}
		v, resp, err := env.Client.AdminOther.ListCurrencies(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 8)
	})

	t.Run("ListCurrenciesExtended", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.currency_extended@get.json"
			method  = "GET"
			path    = "/3.0/admin/currency"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOtherListCurrenciesOpts{GetOnlyActive: 1}
		v, resp, err := env.Client.AdminOther.ListCurrenciesExtended(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("ListPaymentSystems", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.payment_systems@get.json"
			method  = "GET"
			path    = "/3.0/admin/payment_systems"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOther.ListPaymentSystems(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)
	})

	t.Run("ListCustomFields", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.custom_fields@get.json"
			method  = "GET"
			path    = "/3.0/admin/custom_fields"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOther.ListCustomFields(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("ListDomains", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.domains@get.json"
			method  = "GET"
			path    = "/3.0/admin/domains"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOther.ListDomains(env.Ctx)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)
	})

	t.Run("GetTicket", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "1"
			fixture = "3.0.admin.ticket.{id}@get.json"
			method  = "GET"
			path    = fmt.Sprintf("/3.0/admin/ticket/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOther.GetTicket(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.ID == id)
	})

	t.Run("ListTickets", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.tickets@get.json"
			method  = "GET"
			path    = "/3.0/admin/tickets"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOtherListTicketsOpts{Page: 1}
		v, resp, err := env.Client.AdminOther.ListTickets(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)
	})

	t.Run("ApproveTicket", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = 17340
			fixture = "3.0.admin.ticket.{id}.offer@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/ticket/%d/offer", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOtherApproveTicketOpts{Do: "approve"}
		resp, err := env.Client.AdminOther.ApproveTicket(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("ListPixels", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      uint64 = 610
			fixture        = "3.0.partner.pixels.{id}@get.json"
			method         = "GET"
			path           = fmt.Sprintf("/3.0/partner/pixels/%d", id)
			status         = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOther.ListPixels(env.Ctx, id)
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

		opts := &affise.AdminOtherCreatePixelOpts{
			AffiliateID: 610,
			OfferID:     906,
			Name:        "test",
			Code:        "<script>test</script>",
			CodeType:    "javascript",
		}
		v, resp, err := env.Client.AdminOther.CreatePixel(env.Ctx, opts)
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

		opts := &affise.AdminOtherUpdatePixelOpts{
			Name:     "test2",
			Code:     "<script>test</script>",
			CodeType: "javascript",
		}
		v, resp, err := env.Client.AdminOther.UpdatePixel(env.Ctx, id, opts)
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

		v, resp, err := env.Client.AdminOther.DeletePixel(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.OfferID == "906")
	})

	t.Run("ListSmartLinkCategories", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.smartlink.categories@get.json"
			method  = "GET"
			path    = "/3.0/admin/smartlink/categories"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOtherListSmartLinkCategoriesOpts{Name: "test1"}
		v, resp, err := env.Client.AdminOther.ListSmartLinkCategories(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("CreateSmartLinkCategory", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.smartlink.category@post.json"
			method  = "POST"
			path    = "/3.0/admin/smartlink/category"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOtherCreateSmartLinkCategoryOpts{
			Name:        "test",
			DomainID:    5,
			Description: "test",
		}
		v, resp, err := env.Client.AdminOther.CreateSmartLinkCategory(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.Name == opts.Name)
	})

	t.Run("UpdateSmartLinkCategory", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "595fd4877e28fee8428b459f"
			fixture = "3.0.admin.smartlink.category.{id}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/smartlink/category/%s", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminOtherUpdateSmartLinkCategoryOpts{
			Name:        "test",
			DomainID:    5,
			Description: "test",
		}
		v, resp, err := env.Client.AdminOther.UpdateSmartLinkCategory(env.Ctx, id, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.ID == id)
	})

	t.Run("DeleteSmartLinkCategory", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "595fd4877e28fee8428b459f"
			fixture = "3.0.admin.smartlink.category.{id}.remove@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/smartlink/category/%s/remove", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOther.DeleteSmartLinkCategory(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v.ID == id)
	})

	t.Run("GetSmartLinkOffersCount", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "500"
			fixture = "3.0.admin.smartlink.category.{id}.offers-count@get.json"
			method  = "GET"
			path    = fmt.Sprintf("/3.0/admin/smartlink/category/%s/offers-count", id)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminOther.GetSmartLinkOffersCount(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, v == 2)
	})
}
