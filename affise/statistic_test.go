package affise_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/clobucks/go-sdk/affise"
)

func TestStatisticService_Mock(t *testing.T) {
	t.Parallel()
	t.Run("Custom", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.custom@get.json"
			method  = "GET"
			path    = "/3.0/stats/custom"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticCustomOpts{
			Slice:           []string{"year", "month", "day"},
			ConversionTypes: []string{"total", "confirmed"},
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}

		v, resp, err := env.Client.Statistic.Custom(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Year == 2017)
		require.True(t, len(v[0].Actions) == 2)
	})

	t.Run("ConversionsByID", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			id      = "5bd00d73901fcf20008b4574"
			fixture = "3.0.stats.conversionsbyid@get.json"
			method  = "GET"
			path    = "/3.0/stats/conversionsbyid"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		v, resp, err := env.Client.Statistic.ConversionsByID(env.Ctx, id)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, id == v.ID)
	})

	t.Run("Conversions", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.conversions@get.json"
			method  = "GET"
			path    = "/3.0/stats/conversions"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticConversionsOpts{
			DateFrom: "2020-11-01",
		}
		v, resp, err := env.Client.Statistic.Conversions(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotEmpty(t, v)
		require.True(t, v[0].Offer.ID == 71)
		require.True(t, v[0].Partner.ID == 12)
		require.True(t, v[0].Advertiser.ID == "5f059c53d346519b154f20d4")
	})

	t.Run("Clicks", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.clicks@get.json"
			method  = "GET"
			path    = "/3.0/stats/clicks"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticClicksOpts{
			DateFrom: "2020-01-01",
			DateTo:   "2021-01-01",
			Limit:    20,
		}
		v, resp, err := env.Client.Statistic.Clicks(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.NotEmpty(t, v)
		require.True(t, v[0].Offer.ID == 104)
		require.True(t, v[0].Partner.ID == 4)
	})

	t.Run("GetByDate", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbydate@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbydate"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByDateOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-01-01",
				DateTo:   "2017-01-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByDate(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Year == 2017)
		require.True(t, len(v[0].Actions) == 6)
	})

	t.Run("GetByHour", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbyhour@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbyhour"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByHourOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-02-01",
				DateTo:   "2017-02-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByHour(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Month == 5)
		require.True(t, len(v[0].Actions) == 6)
	})

	t.Run("GetBySub", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbysub@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbysub"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetBySubOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetBySub(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Sub1 == "")
		require.True(t, v[0].Actions["total"].Count == 12600)
	})

	t.Run("GetByOffer", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbyprogram@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbyprogram"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByOfferOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-04-01",
				DateTo:   "2017-04-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByOffer(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Offer.ID == 906)
		require.True(t, v[0].Actions["hold"].Count == 12599)
	})

	t.Run("GetByAdvertiser", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbyadvertiser@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbyadvertiser"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByAdvertiserOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-05-01",
				DateTo:   "2017-05-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByAdvertiser(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Advertiser.ID == "56cc49dc3b7d9b89058b45f0")
		require.True(t, v[0].Actions["confirmed"].Count == 1)
	})

	t.Run("GetByAccountManager", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbyaccountmanager@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbyaccountmanager"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByAccountManagerOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     2,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-06-01",
				DateTo:   "2017-06-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByAccountManager(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)
		require.True(t, v[0].Slice.AdvertiserManagerID.FirstName == "Undefined")
		require.True(t, v[1].Actions["pending"].Earning == 0.1587)
	})

	t.Run("GetByAffiliateManager", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbyaffiliatemanager@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbyaffiliatemanager"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByAffiliateManagerOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     2,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-07-01",
				DateTo:   "2017-07-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByAffiliateManager(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)
		require.True(t, v[1].Slice.AffiliateManagerID.LastName == "GmbH")
	})

	t.Run("GetByAffiliate", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbypartner@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbypartner"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByAffiliateOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-09-01",
				DateTo:   "2017-09-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByAffiliate(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Affiliate.ID == 610)
	})

	t.Run("GetByAffiliateByDate", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbypartnerbydate@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbypartnerbydate"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByAffiliateByDateOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-08-01",
				DateTo:   "2017-08-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByAffiliateByDate(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Affiliate.Login == "affiliate")
	})

	t.Run("GetByCountries", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbycountries@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbycountries"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByCountriesOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-09",
				DateTo:   "2017-03-10",
			},
		}
		v, resp, err := env.Client.Statistic.GetByCountries(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.Equal(t, 1, resp.Meta.Pagination.Page)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Country == "USA")
	})

	t.Run("GetByBrowsers", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbybrowsers@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbybrowsers"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByBrowsersOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-11",
			},
		}
		v, resp, err := env.Client.Statistic.GetByBrowsers(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Browser == "")
	})

	t.Run("GetByBrowserVersion", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbybrowsersversion@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbybrowsersversion"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByBrowserVersionOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2018-03-01",
				DateTo:   "2018-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByBrowserVersion(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.BrowserVersion == "70")
	})

	t.Run("GetByLanding", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbylanding@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbylanding"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByLandingOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-11-11",
				DateTo:   "2017-11-11",
			},
		}
		v, resp, err := env.Client.Statistic.GetByLanding(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Landing == "1543238303")
		require.True(t, v[0].LandingsInfo["1543238303"].URL == "http://test-url.com")
	})

	t.Run("GetByPrelanding", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbyprelanding@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbyprelanding"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByPrelandingOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-12-01",
				DateTo:   "2017-12-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByPrelanding(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Prelanding == "1543243821")
		require.True(t, v[0].LandingsInfo["1543243821"].Title == "333")
	})

	t.Run("GetByMobileCarrier", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbymobilecarrier@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbymobilecarrier"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByMobileCarrierOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2019-03-01",
				DateTo:   "2019-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByMobileCarrier(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.ISP == "")
	})

	t.Run("GetByConnectionType", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbyconnectiontype@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbyconnectiontype"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByConnectionTypeOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByConnectionType(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.ConnType == "")
	})

	t.Run("GetByOS", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbyos@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbyos"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByOSOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByOS(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.OS == "")
	})

	t.Run("GetByVersions", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbyversions@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbyversions"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByVersionsOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByVersions(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.OSVersion == "")
	})

	t.Run("GetByGoal", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbygoal@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbygoal"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByGoalOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByGoal(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Goal == "1")
	})

	t.Run("GetByCities", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbycities@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbycities"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByCitiesOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByCities(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.City == "0")
	})

	t.Run("GetByDevices", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbydevices@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbydevices"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByDevicesOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByDevices(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.Device == "")
	})

	t.Run("GetByDeviceModels", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbydevicemodels@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbydevicemodels"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByDeviceModelsOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByDeviceModels(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Slice.DeviceModel == "")
	})

	t.Run("GetByReferralPayments", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getreferralpayments@get.json"
			method  = "GET"
			path    = "/3.0/stats/getreferralpayments"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByReferralPaymentsOpts{
			Limit:    1,
			DateFrom: "2017-03-01",
			DateTo:   "2017-03-01",
		}
		v, resp, err := env.Client.Statistic.GetByReferralPayments(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].Ref == "3")
	})

	t.Run("FindSubs", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.find-subs@get.json"
			method  = "GET"
			path    = "/3.0/stats/find-subs"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticFindSubsOpts{Sub1: "test1"}
		v, resp, err := env.Client.Statistic.FindSubs(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
	})

	t.Run("ServerPostbacks", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.serverpostbacks@get.json"
			method  = "GET"
			path    = "/3.0/stats/serverpostbacks"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticServerPostbacksOpts{
			ClickID: "59359dcb7e28fee0558b4567",
			Goal:    "1",
		}
		v, resp, err := env.Client.Statistic.ServerPostbacks(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].IDStruct.ID == "59359e1d7e28feb7568b4569")
		require.True(t, v[0].GetStruct.Clickid == "59359dcb7e28fee0558b4567")
		require.True(t, v[0].Track.Partner.ID == "610")
	})

	t.Run("AffiliatePostbacks", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.affiliatepostbacks@get.json"
			method  = "GET"
			path    = "/3.0/stats/affiliatepostbacks"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticAffiliatePostbacksOpts{
			DateFrom: "2017-11-25",
			DateTo:   "2017-11-28",
			Goal:     "1",
		}
		v, resp, err := env.Client.Statistic.AffiliatePostbacks(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].IDStruct.ID == "5a1d248f1bfa2441008b4567")
		require.True(t, v[0].HTTPCode == 200)
	})

	t.Run("Caps", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.1.stats.caps@get.json"
			method  = "GET"
			path    = "/3.1/stats/caps"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticCapsOpts{
			OfferID: []int{10},
		}
		v, resp, err := env.Client.Statistic.Caps(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, len(v[0].Stats) == 1)
		require.True(t, v[0].Stats[0].AffiliateType == "string")
	})

	t.Run("GetByTrafficback", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.getbytrafficback@get.json"
			method  = "GET"
			path    = "/3.0/stats/getbytrafficback"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticGetByTrafficbackOpts{
			OrderType: "asc",
			Locale:    "en",
			Limit:     1,
			StatFilter: affise.StatFilter{
				DateFrom: "2017-03-01",
				DateTo:   "2017-03-01",
			},
		}
		v, resp, err := env.Client.Statistic.GetByTrafficback(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)
		require.True(t, v[0].Slice.TrafficbackReason == "unknown-affiliate")
	})

	t.Run("RetentionRate", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.retentionrate@get.json"
			method  = "GET"
			path    = "/3.0/stats/retentionrate"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticRetentionRateOpts{
			DateFrom: "2018-10-16",
			DateTo:   "2018-10-19",
		}
		v, resp, err := env.Client.Statistic.RetentionRate(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 2)
		require.True(t, v[0].RrInstall == "66.66")
		require.True(t, v[1].RrOther1 == "100")
	})

	t.Run("TimeToAction", func(t *testing.T) {
		t.Parallel()
		env := newTestEnv(t)
		defer env.teardown()

		var (
			fixture = "3.0.stats.time-to-action@get.json"
			method  = "GET"
			path    = "/3.0/stats/time-to-action"
			status  = 200
		)
		env.mockHandle(t, fixture, method, path, status)

		opts := &affise.StatisticTimeToActionOpts{
			DateFrom: "2018-10-16",
			DateTo:   "2018-10-19",
			Timezone: "Europe/Berlin",
		}
		v, resp, err := env.Client.Statistic.TimeToAction(env.Ctx, opts)
		require.NoError(t, err)
		require.Equal(t, 1, resp.Meta.Status)
		require.True(t, len(v) == 1)
		require.True(t, v[0].AffiliateID == 70)
	})
}
