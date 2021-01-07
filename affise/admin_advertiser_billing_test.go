package affise_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestAdminAdvertiserBillingService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("List", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.advertiser-invoices@get.json"
			method  = "GET"
			path    = "/3.0/admin/advertiser-invoices"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAdvertiserBillingListOpts{Page: 1}
		v, resp, err := env.Client.AdminAdvertiserBilling.List(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			number  = 1
			fixture = "3.0.admin.advertiser-invoice.{number}@get.json"
			method  = "GET"
			path    = fmt.Sprintf("/3.0/admin/advertiser-invoice/%d", number)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.AdminAdvertiserBilling.Get(env.Ctx, number)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, number, v.Number)
	})

	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.admin.advertiser-invoice@post.json"
			method  = "POST"
			path    = "/3.0/admin/advertiser-invoice"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAdvertiserBillingCreateOpts{
			Status:     "unpaid",
			SupplierID: "5a37c01cbf0b6b18008b4567",
			StartDate:  "2017-12-05",
			EndDate:    "2017-12-07",
			Currency:   "USD",
			Comment:    "222",
			Details: []affise.DetailOpts{
				{
					OfferID:    1,
					PayoutType: "RPA",
					Actions:    100,
					Amount:     100,
					Comment:    "foo",
				},
			},
		}
		resp, err := env.Client.AdminAdvertiserBilling.Create(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})

	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			number  = 1
			fixture = "3.0.admin.advertiser-invoice.{number}@post.json"
			method  = "POST"
			path    = fmt.Sprintf("/3.0/admin/advertiser-invoice/%d", number)
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.AdminAdvertiserBillingUpdateOpts{
			Status: "unpaid",
		}
		resp, err := env.Client.AdminAdvertiserBilling.Update(env.Ctx, number, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
	})
}
