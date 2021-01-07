package affise

import (
	"context"
	"encoding/json"
	"net/http"
)

type Conversion struct {
	ID            string      `json:"id"`
	ActionID      string      `json:"action_id"`
	Status        string      `json:"status"`
	ConversionID  string      `json:"conversion_id"`
	Cbid          string      `json:"cbid"`
	Currency      string      `json:"currency"`
	Offer         *Offer      `json:"offer"`
	OfferID       uint64      `json:"offer_id"`
	Goal          string      `json:"goal"`
	IP            string      `json:"ip"`
	Country       string      `json:"country"`
	CountryName   string      `json:"country_name"`
	District      string      `json:"district"`
	City          string      `json:"city"`
	CityID        int         `json:"city_id"`
	IspCode       string      `json:"isp_code"`
	UA            string      `json:"ua"`
	Browser       string      `json:"browser"`
	OS            string      `json:"os"`
	Device        string      `json:"device"`
	DeviceType    string      `json:"device_type"`
	Sub1          string      `json:"sub1"`
	Sub2          string      `json:"sub2"`
	Sub3          string      `json:"sub3"`
	Sub4          string      `json:"sub4"`
	Sub5          string      `json:"sub5"`
	Sub6          string      `json:"sub6"`
	Sub7          string      `json:"sub7"`
	Sub8          string      `json:"sub8"`
	CustomField1  string      `json:"custom_field_1"`
	CustomField2  string      `json:"custom_field_2"`
	CustomField3  string      `json:"custom_field_3"`
	CustomField4  string      `json:"custom_field_4"`
	CustomField5  string      `json:"custom_field_5"`
	CustomField6  string      `json:"custom_field_6"`
	CustomField7  string      `json:"custom_field_7"`
	Comment       string      `json:"comment"`
	CreatedAt     string      `json:"created_at"`
	ClickTime     string      `json:"click_time"`
	Referrer      string      `json:"referrer"`
	UpdatedAt     string      `json:"updatedAt"`
	Clickid       string      `json:"clickid"`
	Partner       *Affiliate  `json:"partner"`
	AdvertiserID  string      `json:"supplier_id"`
	AffiliateID   uint64      `json:"partner_id"`
	GoalValue     string      `json:"goal_value"`
	Sum           float32     `json:"sum"`
	Revenue       float32     `json:"revenue"`
	Payouts       float32     `json:"payouts"`
	Earnings      float32     `json:"earnings"`
	Advertiser    *Advertiser `json:"advertiser"`
	PaymentType   string      `json:"payment_type"`
	PaymentStatus string      `json:"payment_status"`
	IsPaid        string      `json:"is_paid"`
	IosIdfa       string      `json:"ios_idfa"`
	AndroidID     string      `json:"android_id"`
	Price         float32     `json:"price"`
	LandingID     uint64      `json:"landing_id"`
	PrelandingID  uint64      `json:"prelanding_id"`
	// todo interface{} fields
	CurrencyID     interface{} `json:"currency_id"`
	Forensiq       interface{} `json:"forensiq"`
	HoldDateExpire interface{} `json:"hold_date_expire"`
}

type Click struct {
	ID           string     `json:"id"`
	IP           string     `json:"ip"`
	UA           string     `json:"ua"`
	Country      string     `json:"country"`
	City         string     `json:"city"`
	Device       string     `json:"device"`
	OS           string     `json:"os"`
	Browser      string     `json:"browser"`
	Referrer     string     `json:"referrer"`
	Sub1         string     `json:"sub1"`
	Sub2         string     `json:"sub2"`
	Sub3         string     `json:"sub3"`
	Sub4         string     `json:"sub4"`
	Sub5         string     `json:"sub5"`
	Sub6         string     `json:"sub6"`
	Sub7         string     `json:"sub7"`
	Sub8         string     `json:"sub8"`
	Offer        *Offer     `json:"offer"`
	ConversionID string     `json:"conversion_id"`
	IosIdfa      string     `json:"ios_idfa"`
	AndroidID    string     `json:"android_id"`
	CreatedAt    string     `json:"created_at"`
	Uniq         bool       `json:"uniq"`
	Cbid         string     `json:"cbid"`
	AffiliateID  uint64     `json:"partner_id"`
	Partner      *Affiliate `json:"partner"`
}

type AdvertiserManagerID struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AffiliateManagerID struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LandingInfo struct {
	URL        string `json:"url"`
	PreviewURL string `json:"preview_url"`
	Title      string `json:"title"`
}

type StatSlice struct {
	Year                int                  `json:"year,omitempty"`
	Quarter             int                  `json:"quarter,omitempty"`
	Month               int                  `json:"month,omitempty"`
	Day                 int                  `json:"day,omitempty"`
	Hour                int                  `json:"hour,omitempty"`
	Country             string               `json:"country,omitempty"`
	OS                  string               `json:"os,omitempty"`
	OSVersion           string               `json:"os_version,omitempty"`
	Device              string               `json:"device,omitempty"`
	DeviceModel         string               `json:"device_model,omitempty"`
	Browser             string               `json:"browser,omitempty"`
	BrowserVersion      string               `json:"browser_version,omitempty"`
	Landing             string               `json:"landing,omitempty"`
	Prelanding          json.Number          `json:"prelanding,omitempty"`
	Sub1                string               `json:"sub1,omitempty"`
	Sub2                string               `json:"sub2,omitempty"`
	Sub3                string               `json:"sub3,omitempty"`
	Sub4                string               `json:"sub4,omitempty"`
	Sub5                string               `json:"sub5,omitempty"`
	Goal                string               `json:"goal,omitempty"`
	City                string               `json:"city,omitempty"`
	ISP                 string               `json:"isp,omitempty"`
	ConnType            string               `json:"conn_type,omitempty"`
	TrafficbackReason   string               `json:"trafficback_reason,omitempty"`
	Offer               *Offer               `json:"offer,omitempty"`
	Advertiser          *Advertiser          `json:"advertiser,omitempty"`
	AdvertiserManagerID *AdvertiserManagerID `json:"advertiser_manager_id,omitempty"`
	AffiliateManagerID  *AffiliateManagerID  `json:"affiliate_manager_id,omitempty"`
	Affiliate           *Affiliate           `json:"affiliate,omitempty"`
}

type StatTraffic struct {
	Raw  string `json:"raw,omitempty"`
	Uniq string `json:"uniq,omitempty"`
}

type StatAction struct {
	Revenue float32 `json:"revenue,omitempty"`
	Charge  float32 `json:"charge,omitempty"`
	Earning float32 `json:"earning,omitempty"`
	Null    float32 `json:"null,omitempty"`
	Count   float32 `json:"count,omitempty"`
}

type Stat struct {
	Slice        StatSlice              `json:"slice"`
	Traffic      StatTraffic            `json:"traffic"`
	Actions      map[string]StatAction  `json:"actions"`
	Ratio        string                 `json:"ratio"`
	Epc          int                    `json:"epc"`
	LandingsInfo map[string]LandingInfo `json:"landings_info"`
}

type RefPayment struct {
	AffiliateID             uint64 `json:"pid,string"`
	Ref                     string `json:"ref"`
	Status                  string `json:"status"`
	IsPaid                  string `json:"is_paid"`
	Currency                string `json:"currency"`
	Count                   string `json:"count"`
	MaxCreatedAt            string `json:"max_created_at"`
	DateRegistrationPartner string `json:"date_registration_partner"`
	SumRevenue              string `json:"sum_revenue"`
}

type Sub map[string]string

type Track struct {
	ID        string `json:"id"`
	IP        string `json:"ip"`
	Ua        string `json:"ua"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Device    string `json:"device"`
	Os        string `json:"os"`
	Browser   string `json:"browser"`
	Offer     *Offer `json:"offer"`
	Referrer  string `json:"referrer"`
	ClickID   string `json:"click_id"`
	Sub1      string `json:"sub1"`
	Sub2      string `json:"sub2"`
	Sub3      string `json:"sub3"`
	Sub4      string `json:"sub4"`
	Sub5      string `json:"sub5"`
	OfferID   string `json:"offer_id"`
	CreatedAt string `json:"created_at"`
	Uniq      int    `json:"uniq"`
	Partner   struct {
		ID    string `json:"id"`
		Login string `json:"login"`
		Email string `json:"email"`
	} `json:"partner"`
	// todo interface{} fields
	ConversionID interface{} `json:"conversion_id"`
	IosIdfa      interface{} `json:"ios_idfa"`
	AndroidID    interface{} `json:"android_id"`
	Cbid         interface{} `json:"cbid"`
}

type StatPostback struct {
	IDStruct struct {
		ID string `json:"$id"`
	} `json:"_id"`
	GetStruct struct {
		Clickid string `json:"clickid"`
	} `json:"_get"`
	PostStruct []interface{} `json:"_post"`
	Date       struct {
		Sec  int `json:"sec"`
		Usec int `json:"usec"`
	} `json:"date"`
	Get         string `json:"get"`
	Post        string `json:"post"`
	Server      string `json:"server"`
	Response    string `json:"response"`
	Track       *Track `json:"track"`
	AffiliateID uint64 `json:"pid"`
	LeadID      string `json:"lead_id"`
	HTTPCode    int    `json:"http_code"`
	PostbackURL string `json:"postback_url"`
	OfferID     int    `json:"offer_id"`
	JobID       string `json:"job_id"`
	Goal        string `json:"goal"`
	Status      int    `json:"status"`
}

type RetentionRate struct {
	AffiliateID  uint64      `json:"affiliate_id"`
	Date         string      `json:"date"`
	RrInstall    json.Number `json:"rr_install"`
	RrOther1     json.Number `json:"rr_other1"`
	RrOther2     json.Number `json:"rr_other2"`
	InstallCount json.Number `json:"install_count"`
}

type TimeToAction struct {
	AffiliateID      uint64 `json:"affiliate_id"`
	Clicks           int    `json:"clicks"`
	TotalConversions int    `json:"total_conversions"`
	Tta30            int    `json:"tta_30"`
	Tta600           int    `json:"tta_600"`
	TtaInf           int    `json:"tta_inf"`
}

type StatCap struct {
	OfferID int `json:"offer_id"`
	Stats   []struct {
		ID            string              `json:"id"`
		Timeframe     string              `json:"timeframe"`
		Type          string              `json:"type"`
		Value         int                 `json:"value"`
		CurrentValue  int                 `json:"current_value"`
		IsRemaining   bool                `json:"is_remaining"`
		ResetToValue  int                 `json:"reset_to_value"`
		AffiliateType string              `json:"affiliate_type"`
		Affiliates    []int               `json:"affiliates"`
		Goals         []map[string]string `json:"goals"`
		GoalType      string              `json:"goal_type"`
		Countries     []string            `json:"countries"`
		CountryType   string              `json:"country_type"`
	} `json:"stats"`
}

type StatFilter struct {
	DateFrom            string   `schema:"filter[date_from]"`                       // Date from (Available: YYYY-MM-DD)
	DateTo              string   `schema:"filter[date_to]"`                         // Date to (Available: YYYY-MM-DD)
	Currency            []string `schema:"filter[currency],omitempty"`              // The list of a currencies code you can get from API /3.0/admin/currency (Default: All currencies code)
	Advertiser          []string `schema:"filter[advertiser],omitempty"`            // Advertiser ID’s
	Offer               []int    `schema:"filter[offer],omitempty"`                 // Offers ID’s
	Manager             []string `schema:"filter[manager],omitempty"`               // Managers ID’s
	AdvertiserManagerID []string `schema:"filter[advertiser_manager_id],omitempty"` // Advertiser managers ID’s
	Partner             []string `schema:"filter[partner],omitempty"`               // ONLY FOR ADMIN  Partners ID’s. (Default: If the request from not an admin then default a ID partner)
	Country             []string `schema:"filter[country],omitempty"`               // Countries codes. Example: “US”
	OS                  []string `schema:"filter[os],omitempty"`                    // Os
	Goal                []string `schema:"filter[goal],omitempty"`                  // Goal
	Sub1                []string `schema:"filter[sub1],omitempty"`                  // Sub number 1
	Sub2                []string `schema:"filter[sub2],omitempty"`                  // Sub number 2
	Sub3                []string `schema:"filter[sub3],omitempty"`                  // Sub number 3
	Sub4                []string `schema:"filter[sub4],omitempty"`                  // Sub number 4
	Sub5                []string `schema:"filter[sub5],omitempty"`                  // Sub number 5
	Sub6                []string `schema:"filter[sub6],omitempty"`                  // Sub 6
	Sub7                []string `schema:"filter[sub7],omitempty"`                  // Sub 7
	Sub8                []string `schema:"filter[sub8],omitempty"`                  // Sub 8
	Device              []string `schema:"filter[device],omitempty"`                // Device
	SmartID             []string `schema:"filter[smart_id],omitempty"`              // Allowed only when the smart slice selected SmartLink categories ID’s
	Nonzero             int      `schema:"filter[nonzero],omitempty"`               // Non-zero conversions (Available: 1, 0)
	AdvertiserTag       string   `schema:"filter[advertiser_tag],omitempty"`        // Comma separated array of strings. Example: tag1,tag2,tag3
	AffiliateTag        string   `schema:"filter[affiliate_tag],omitempty"`         // Comma separated array of strings. Example: tag1,tag2,tag3
	OfferTag            string   `schema:"filter[offer_tag],omitempty"`             // Comma separated array of strings. Example: tag1,tag2,tag3
}

type StatisticService struct {
	client *Client
}

// StatisticCustomOpts specifies options for Custom.
type StatisticCustomOpts struct {
	StatFilter
	Slice           []string `schema:"slice"`                     // REQUIRED Custom stats slice (Available: hour, month, quarter, year, day, offer, country, city, os, os_version, device, device_model, browser, goal, sub1, sub2, sub3, sub4, sub5.  Only for admin: advertiser, affiliate, manager, smart_id.  Only for users with special permission: trafficback_reason)
	Locale          string   `schema:"locale,omitempty"`          // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	ConversionTypes []string `schema:"conversionTypes,omitempty"` // Only this conversion types will be output (Available: total, confirmed, pending, declined, hold, not_found)
	Page            int      `schema:"page,omitempty"`            // Page of stat entities (Default: 1)
	Limit           int      `schema:"limit,omitempty"`           // Limit of stat entities (Default: 100)
	OrderType       string   `schema:"orderType,omitempty"`       // Sorting order (Default: asc  Available: asc, desc)
	Order           []string `schema:"order,omitempty"`           // Sort by field (Available: hour, month, quarter, year, day, currency, offer, country, city, os, os_version, device, device_model, browser, goal, sub1, sub2, sub3, sub4, sub5, confirmed_earning, raw, uniq, total_count, total_revenue, total_null, pending_count, pending_revenue, declined_count, declined_revenue, hold_count, hold_revenue, confirmed_count, confirmed_revenue.  Only for admin: advertiser, affiliate, manager)
	Timezone        string   `schema:"timezone,omitempty"`        // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
}

// statisticCustomResponse specifies response for Custom.
type statisticCustomResponse struct {
	Stats []*Stat `json:"stats"`
}

// Custom gets custom statistics.
func (s *StatisticService) Custom(ctx context.Context, opts *StatisticCustomOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/custom"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticCustomResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// ConversionsByIDOpts specifies options for ConversionsByID.
type ConversionsByIDOpts struct {
	ID string `json:"id"`
}

// statisticConversionsByIDResponse specifies response for ConversionsByID.
type statisticConversionsByIDResponse struct {
	Conversion *Conversion `json:"conversion"`
}

// ConversionsByID gets conversion
// NOTE: Available only for admin API-Key.
func (s *StatisticService) ConversionsByID(ctx context.Context, id string) (*Conversion, *Response, error) {
	path := "/3.0/stats/conversionsbyid"
	opts := ConversionsByIDOpts{ID: id}

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticConversionsByIDResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Conversion, resp, nil
}

// StatisticConversionsOpts specifies options for Conversions.
type StatisticConversionsOpts struct {
	DateFrom       string   `schema:"date_from,omitempty"`        // Date from (Available: YYYY-MM-DD Default: day one week ago)
	DateTo         string   `schema:"date_to,omitempty"`          // Date to (Available: YYYY-MM-DD Default: date now)
	UpdateFromDate string   `schema:"update_from_date,omitempty"` // Last update date point (Available: YYYY-MM-DD)
	UpdateFromHour int      `schema:"update_from_hour,omitempty"` // Last update hour point
	Status         []int    `schema:"status,omitempty"`           // Status conversions. 1 = confirmed, 2 = pending, 3 = declined, 4 = not_found, 5 = hold (Available: 1, 2, 3, 4, 5)
	Offer          []int    `schema:"offer,omitempty"`            // Offer ID collection
	Advertiser     []string `schema:"advertiser,omitempty"`       // Advertiser ID collection
	Country        []string `schema:"country,omitempty"`          // Countries codes. Example: “US”
	Browser        string   `schema:"browser,omitempty"`          // Browser
	ActionID       string   `schema:"action_id,omitempty"`        // Cbid
	Clickid        string   `schema:"clickid,omitempty"`          // Click ID
	OS             string   `schema:"os,omitempty"`               // Os
	Goal           string   `schema:"goal,omitempty"`             // Goal
	Device         string   `schema:"device,omitempty"`           // Device (Available: tablet, desktop, mobile)
	Payouts        float64  `schema:"payouts,omitempty"`          // Payout for affiliate
	Currency       int      `schema:"currency,omitempty"`         // ID currency
	Hour           int      `schema:"hour,omitempty"`             // Hour point  Allows only for one day period (Between 0 and 23)
	Timezone       string   `schema:"timezone,omitempty"`         // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	CustomField1   string   `schema:"custom_field_1,omitempty"`   // Custom field 1
	CustomField2   string   `schema:"custom_field_2,omitempty"`   // Custom field 2
	CustomField3   string   `schema:"custom_field_3,omitempty"`   // Custom field 3
	CustomField4   string   `schema:"custom_field_4,omitempty"`   // Custom field 4
	CustomField5   string   `schema:"custom_field_5,omitempty"`   // Custom field 5
	CustomField6   string   `schema:"custom_field_6,omitempty"`   // Custom field 6
	CustomField7   string   `schema:"custom_field_7,omitempty"`   // Custom field 7
	Subid1         string   `schema:"subid1,omitempty"`           // Sub 1
	Subid2         string   `schema:"subid2,omitempty"`           // Sub 2
	Subid3         string   `schema:"subid3,omitempty"`           // Sub 3
	Subid4         string   `schema:"subid4,omitempty"`           // Sub 4
	Subid5         string   `schema:"subid5,omitempty"`           // Sub 5
	Partner        []int    `schema:"partner,omitempty"`          // ONLY FOR ADMIN  Affiliates
	Revenue        float64  `schema:"revenue,omitempty"`          // ONLY FOR ADMIN Revenue
	Page           int      `schema:"page,omitempty"`             // Page of stat entities (Default: 1)
	Limit          int      `schema:"limit,omitempty"`            // Limit of stat entities (Default: 100)
	RawExport      int      `schema:"raw_export,omitempty"`       // Without mapping related entities (For huge exports) (Default: 0)
}

// statisticConversionsResponse specifies response for Conversions.
type statisticConversionsResponse struct {
	Conversions []*Conversion `json:"conversions"`
}

// Conversions gets conversions.
func (s *StatisticService) Conversions(ctx context.Context, opts *StatisticConversionsOpts) ([]*Conversion, *Response, error) {
	path := "/3.0/stats/conversions"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticConversionsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Conversions, resp, nil
}

// StatisticClicksOpts specifies options for Clicks.
type StatisticClicksOpts struct {
	DateFrom    string   `schema:"date_from"`             // REQUIRED (Available: YYYY-MM-DD)
	DateTo      string   `schema:"date_to"`               // REQUIRED (Available: YYYY-MM-DD)
	Hour        int      `schema:"hour,omitempty"`        // Hour point  Allows only for one day period (Between 0 and 23)
	Offer       []int    `schema:"offer,omitempty"`       // Offer ID’s
	Partner     []int    `schema:"partner,omitempty"`     // Affiliates ID’s
	Country     []string `schema:"country,omitempty"`     // Countries codes. Example: “US”
	Advertisers []string `schema:"advertisers,omitempty"` // ONLY FOR ADMIN Advertiser ID collection
	Timezone    string   `schema:"timezone,omitempty"`    // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page        int      `schema:"page,omitempty"`        // Page of stat entities (Default: 1)
	Limit       int      `schema:"limit,omitempty"`       // Limit of stat entities (Default: 100)
}

// statisticClicksResponse specifies response for Clicks.
type statisticClicksResponse struct {
	Clicks []*Click `json:"clicks"`
}

// Clicks gets clicks
// NOTE: Available only for admin API-Key.
func (s *StatisticService) Clicks(ctx context.Context, opts *StatisticClicksOpts) ([]*Click, *Response, error) {
	path := "/3.0/stats/clicks"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticClicksResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Clicks, resp, nil
}

// StatisticGetByDateOpts specifies options for GetByDate.
type StatisticGetByDateOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByDateResponse specifies response for GetByDate.
type statisticGetByDateResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByDate gets statistics by date.
func (s *StatisticService) GetByDate(ctx context.Context, opts *StatisticGetByDateOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbydate"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByDateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByHourOpts specifies options for GetByHour.
type StatisticGetByHourOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByHourResponse specifies response for GetByHour.
type statisticGetByHourResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByHour gets statistics by hour.
func (s *StatisticService) GetByHour(ctx context.Context, opts *StatisticGetByHourOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbyhour"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByHourResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetBySubOpts specifies options for GetBySub.
type StatisticGetBySubOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetBySubResponse specifies response for GetBySub.
type statisticGetBySubResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetBySub gets statistics by sub
// NOTE: Available only for partner API-Key.
func (s *StatisticService) GetBySub(ctx context.Context, opts *StatisticGetBySubOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbysub"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetBySubResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByOfferOpts specifies options for GetByOffer.
type StatisticGetByOfferOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByOfferResponse specifies response for GetByOffer.
type statisticGetByOfferResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByOffer gets statistics by offer.
func (s *StatisticService) GetByOffer(ctx context.Context, opts *StatisticGetByOfferOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbyprogram"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByOfferResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByAdvertiserOpts specifies options for GetByAdvertiser.
type StatisticGetByAdvertiserOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByAdvertiserResponse specifies response for GetByAdvertiser.
type statisticGetByAdvertiserResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByAdvertiser gets statistics by advertiser
// NOTE: Available only for admin API-Key.
func (s *StatisticService) GetByAdvertiser(ctx context.Context, opts *StatisticGetByAdvertiserOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbyadvertiser"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByAdvertiserResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByAccountManagerOpts specifies options for GetByAccountManager.
type StatisticGetByAccountManagerOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByAccountManagerResponse specifies response for GetByAccountManager.
type statisticGetByAccountManagerResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByAccountManager gets statistics by accounts managers
// NOTE: Available only for admin API-Key.
func (s *StatisticService) GetByAccountManager(ctx context.Context, opts *StatisticGetByAccountManagerOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbyaccountmanager"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByAccountManagerResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByAffiliateManagerOpts specifies options for GetByAffiliateManager.
type StatisticGetByAffiliateManagerOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByAffiliateManagerResponse specifies response for GetByAffiliateManager.
type statisticGetByAffiliateManagerResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByAffiliateManager gets statistics by affiliates managers
// NOTE: Available only for admin API-Key.
func (s *StatisticService) GetByAffiliateManager(ctx context.Context, opts *StatisticGetByAffiliateManagerOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbyaffiliatemanager"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByAffiliateManagerResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByPartnerOpts specifies options for GetByAffiliate.
type StatisticGetByAffiliateOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByAffiliateResponse specifies response for GetByAffiliate.
type statisticGetByAffiliateResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByAffiliate gets statistics by affiliate
// NOTE:  Available only for admin API-Key.
func (s *StatisticService) GetByAffiliate(ctx context.Context, opts *StatisticGetByAffiliateOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbypartner"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByAffiliateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByAffiliateByDateOpts specifies options for GetByAffiliateByDate.
type StatisticGetByAffiliateByDateOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByAffiliateByDateResponse specifies response for GetByAffiliateByDate.
type statisticGetByAffiliateByDateResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByAffiliateByDate gets statistics by affiliate and date
// NOTE: Available only for admin API-Key.
func (s *StatisticService) GetByAffiliateByDate(ctx context.Context, opts *StatisticGetByAffiliateByDateOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbypartnerbydate"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByAffiliateByDateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByCountriesOpts specifies options for GetByCountries.
type StatisticGetByCountriesOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByCountriesResponse specifies response for GetByCountries.
type statisticGetByCountriesResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByCountries gets statistics by countries.
func (s *StatisticService) GetByCountries(ctx context.Context, opts *StatisticGetByCountriesOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbycountries"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByCountriesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByBrowsersOpts specifies options for GetByBrowsers.
type StatisticGetByBrowsersOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByBrowsersResponse specifies response for GetByBrowsers.
type statisticGetByBrowsersResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByBrowsers gets statistics by browser.
func (s *StatisticService) GetByBrowsers(ctx context.Context, opts *StatisticGetByBrowsersOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbybrowsers"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByBrowsersResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByBrowserVersionOpts specifies options for GetByBrowserVersion.
type StatisticGetByBrowserVersionOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByBrowserVersionResponse specifies response for GetByBrowserVersion.
type statisticGetByBrowserVersionResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByBrowserVersion gets statistics by browser version.
func (s *StatisticService) GetByBrowserVersion(ctx context.Context, opts *StatisticGetByBrowserVersionOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbybrowsersversion"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByBrowserVersionResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByLandingOpts specifies options for GetByLanding.
type StatisticGetByLandingOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByLandingResponse specifies response for GetByLanding.
type statisticGetByLandingResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByLanding gets statistics by landing.
func (s *StatisticService) GetByLanding(ctx context.Context, opts *StatisticGetByLandingOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbylanding"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByLandingResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByPrelandingOpts specifies options for GetByPrelanding.
type StatisticGetByPrelandingOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByPrelandingResponse specifies response for GetByPrelanding.
type statisticGetByPrelandingResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByPrelanding gets statistics by prelanding.
func (s *StatisticService) GetByPrelanding(ctx context.Context, opts *StatisticGetByPrelandingOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbyprelanding"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByPrelandingResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByMobileCarrierOpts specifies options for GetByMobileCarrier.
type StatisticGetByMobileCarrierOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByMobileCarrierResponse specifies response for GetByMobileCarrier.
type statisticGetByMobileCarrierResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByMobileCarrier gets statistics by mobile carrier.
func (s *StatisticService) GetByMobileCarrier(ctx context.Context, opts *StatisticGetByMobileCarrierOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbymobilecarrier"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByMobileCarrierResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByConnectionTypeOpts specifies options for GetByConnectionType.
type StatisticGetByConnectionTypeOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByConnectionTypeResponse specifies response for GetByConnectionType.
type statisticGetByConnectionTypeResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByConnectionType gets statistics by connection type.
func (s *StatisticService) GetByConnectionType(ctx context.Context, opts *StatisticGetByConnectionTypeOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbyconnectiontype"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByConnectionTypeResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByOSOpts specifies options for GetByOS.
type StatisticGetByOSOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByOSResponse specifies response for GetByOS.
type statisticGetByOSResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByOS gets statistics by OS.
func (s *StatisticService) GetByOS(ctx context.Context, opts *StatisticGetByOSOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbyos"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByOSResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByVersionsOpts specifies options for GetByVersions.
type StatisticGetByVersionsOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByVersionsResponse specifies response for GetByVersions.
type statisticGetByVersionsResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByVersions gets statistics by OS version.
func (s *StatisticService) GetByVersions(ctx context.Context, opts *StatisticGetByVersionsOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbyversions"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByVersionsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByGoalOpts specifies options for GetByGoal.
type StatisticGetByGoalOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByGoalResponse specifies response for GetByGoal.
type statisticGetByGoalResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByGoal gets statistics by goal.
func (s *StatisticService) GetByGoal(ctx context.Context, opts *StatisticGetByGoalOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbygoal"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByGoalResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByCitiesOpts specifies options for GetByCities.
type StatisticGetByCitiesOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByCitiesResponse specifies response for GetByCities.
type statisticGetByCitiesResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByCities gets statistics by cities.
func (s *StatisticService) GetByCities(ctx context.Context, opts *StatisticGetByCitiesOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbycities"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByCitiesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByDevicesOpts specifies options for GetByDevices.
type StatisticGetByDevicesOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByDevicesResponse specifies response for GetByDevices.
type statisticGetByDevicesResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByDevices gets statistics by device.
func (s *StatisticService) GetByDevices(ctx context.Context, opts *StatisticGetByDevicesOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbydevices"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByDevicesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByDeviceModelsOpts specifies options for GetByDeviceModels.
type StatisticGetByDeviceModelsOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByDeviceModelsResponse specifies response for GetByDeviceModels.
type statisticGetByDeviceModelsResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByDeviceModels gets statistics by device model.
func (s *StatisticService) GetByDeviceModels(ctx context.Context, opts *StatisticGetByDeviceModelsOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbydevicemodels"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByDeviceModelsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByReferralPaymentsOpts specifies options for GetByReferralPayments.
type StatisticGetByReferralPaymentsOpts struct {
	DateFrom    string `schema:"date_from"`          // REQUIRED  Date from (Available: DD-MM-YYYY)
	DateTo      string `schema:"date_to"`            // REQUIRED  Date to (Available: DD-MM-YYYY)
	AffiliateID uint64 `schema:"pid,omitempty"`      // Partner ID
	Ref         int    `schema:"ref,omitempty"`      // Referral partner ID
	IsPaid      int    `schema:"is_paid,omitempty"`  // Status (Available: 0 => payouts, 1 => paid, 2 => pending)
	Status      int    `schema:"status,omitempty"`   // Active (Available: 0 => no, 1 => yes)
	Page        int    `schema:"page,omitempty"`     // Page of stat entities (Default: 1)
	Limit       int    `schema:"limit,omitempty"`    // Limit of stat entities (Default: 100)
	Currency    int    `schema:"currency,omitempty"` // ID currency
}

// statisticGetByReferralPaymentsResponse specifies response for GetByReferralPayments.
type statisticGetByReferralPaymentsResponse struct {
	RefPayments []*RefPayment `json:"ref_payments"`
}

// GetByReferralPayments gets statistics by referral payments.
func (s *StatisticService) GetByReferralPayments(ctx context.Context, opts *StatisticGetByReferralPaymentsOpts) ([]*RefPayment, *Response, error) {
	path := "/3.0/stats/getreferralpayments"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByReferralPaymentsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.RefPayments, resp, nil
}

// StatisticFindSubsOpts specifies options for FindSubs.
type StatisticFindSubsOpts struct {
	Sub1  string `schema:"sub1,omitempty"`  // Sub 1
	Sub2  string `schema:"sub2,omitempty"`  // Sub 2
	Sub3  string `schema:"sub3,omitempty"`  // Sub 3
	Sub4  string `schema:"sub4,omitempty"`  // Sub 4
	Sub5  string `schema:"sub5,omitempty"`  // Sub 5
	Page  int    `schema:"page,omitempty"`  // Page of sub entities (Default: 1)
	Limit int    `schema:"limit,omitempty"` // Limit of sub entities (Default: 100)
}

// statisticFindSubsResponse specifies response for FindSubs.
type statisticFindSubsResponse struct {
	Subs []*Sub `json:"subs"`
}

// FindSubs gets sub accounts
// NOTE: Available only for partner API-Key.
func (s *StatisticService) FindSubs(ctx context.Context, opts *StatisticFindSubsOpts) ([]*Sub, *Response, error) {
	path := "/3.0/stats/find-subs"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticFindSubsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Subs, resp, nil
}

// StatisticServerPostbacksOpts specifies options for ServerPostbacks.
type StatisticServerPostbacksOpts struct {
	DateFrom string   `schema:"date_from"`           // REQUIRED (Available: YYYY-MM-DD)
	DateTo   string   `schema:"date_to"`             // REQUIRED (Available: YYYY-MM-DD)
	Offer    []int    `schema:"offer,omitempty"`     // Offers ID’s
	Partner  []int    `schema:"partner,omitempty"`   // Partners ID’s.
	Supplier []string `schema:"supplier,omitempty"`  // Advertiser ID’s.
	ActionID string   `schema:"action_id,omitempty"` // Action id
	ClickID  string   `schema:"click_id,omitempty"`  // Click id
	Goal     string   `schema:"goal,omitempty"`      // Goal
	Status   string   `schema:"status,omitempty"`    // Status
	Timezone string   `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page     int      `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit    int      `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
}

// statisticServerPostbacksResponse specifies response for ServerPostbacks.
type statisticServerPostbacksResponse struct {
	Postbacks []*StatPostback `json:"postbacks"`
}

// ServerPostbacks gets server postbacks
// NOTE  Available only for admin API-Key.
func (s *StatisticService) ServerPostbacks(ctx context.Context, opts *StatisticServerPostbacksOpts) ([]*StatPostback, *Response, error) {
	path := "/3.0/stats/serverpostbacks"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticServerPostbacksResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Postbacks, resp, nil
}

// StatisticAffiliatePostbacksOpts specifies options for AffiliatePostbacks.
type StatisticAffiliatePostbacksOpts struct {
	DateFrom string `schema:"date_from"`           // REQUIRED (Available: YYYY-MM-DD)
	DateTo   string `schema:"date_to"`             // REQUIRED (Available: YYYY-MM-DD)
	Offer    []int  `schema:"offer,omitempty"`     // Offers ID’s
	Partner  []int  `schema:"partner,omitempty"`   // Partners ID’s.
	Goal     string `schema:"goal,omitempty"`      // Goal
	Status   int    `schema:"status,omitempty"`    // Status
	HTTPCode int    `schema:"http_code,omitempty"` // Http code
	Timezone string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page     int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit    int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
}

// statisticAffiliatePostbacksResponse specifies response for AffiliatePostbacks.
type statisticAffiliatePostbacksResponse struct {
	Postbacks []*StatPostback `json:"postbacks"`
}

// AffiliatePostbacks gets partner postbacks
// NOTE  Available only for admin API-Key.
func (s *StatisticService) AffiliatePostbacks(ctx context.Context, opts *StatisticAffiliatePostbacksOpts) ([]*StatPostback, *Response, error) {
	path := "/3.0/stats/affiliatepostbacks"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticAffiliatePostbacksResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Postbacks, resp, nil
}

// StatisticCapsOpts specifies options for Caps.
type StatisticCapsOpts struct {
	OfferID []int `schema:"offer_id"` // REQUIRED  Offers ID’s
}

// statisticCapsResponse specifies response for Caps.
type statisticCapsResponse struct {
	Stats []*StatCap `json:"stats"`
}

// Caps gets stats by cap.
func (s *StatisticService) Caps(ctx context.Context, opts *StatisticCapsOpts) ([]*StatCap, *Response, error) {
	path := "/3.1/stats/caps"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticCapsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticGetByTrafficbackOpts specifies options for GetByTrafficback.
type StatisticGetByTrafficbackOpts struct {
	StatFilter
	Locale    string `schema:"locale,omitempty"`    // Locale for output a cities data when you use the city slice (Default: en  Available: ru, en, es)
	Timezone  string `schema:"timezone,omitempty"`  // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
	Page      int    `schema:"page,omitempty"`      // Page of stat entities (Default: 1)
	Limit     int    `schema:"limit,omitempty"`     // Limit of stat entities (Default: 100)
	OrderType string `schema:"orderType,omitempty"` // Sorting order (Default: asc  Available: asc, desc)
}

// statisticGetByTrafficbackResponse specifies response for GetByTrafficback.
type statisticGetByTrafficbackResponse struct {
	Stats []*Stat `json:"stats"`
}

// GetByTrafficback gets statistics by trafficback.
func (s *StatisticService) GetByTrafficback(ctx context.Context, opts *StatisticGetByTrafficbackOpts) ([]*Stat, *Response, error) {
	path := "/3.0/stats/getbytrafficback"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticGetByTrafficbackResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticRetentionRateOpts specifies options for RetentionRate.
type StatisticRetentionRateOpts struct {
	DateFrom    string   `schema:"date_from"`              // REQUIRED  Date from (Available: YYYY-MM-DD)
	DateTo      string   `schema:"date_to"`                // REQUIRED  Date to (Available: YYYY-MM-DD)
	Offer       int      `schema:"offer"`                  // REQUIRED
	BaseEvent   string   `schema:"base_event"`             // REQUIRED Name based goal (Available: ^[a-zA-Z])
	Events      []string `schema:"events"`                 // REQUIRED events (Available: ^[a-zA-Z])
	AffiliateID uint64   `schema:"affiliate_id,omitempty"` // Affiliates filter
	Timezone    string   `schema:"timezone,omitempty"`     // Timezone name. Example: “Europe/Berlin” (Default: Timezone of your platform)
}

// statisticRetentionRateResponse specifies response for RetentionRate.
type statisticRetentionRateResponse struct {
	Stats []*RetentionRate `json:"stats"`
}

// RetentionRate gets stats retention rate.
func (s *StatisticService) RetentionRate(ctx context.Context, opts *StatisticRetentionRateOpts) ([]*RetentionRate, *Response, error) {
	path := "/3.0/stats/retentionrate"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticRetentionRateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Stats, resp, nil
}

// StatisticTimeToActionOpts specifies options for TimeToAction.
type StatisticTimeToActionOpts struct {
	DateFrom     string   `schema:"date_from"`               // REQUIRED  Date from (Available: YYYY-MM-DD)
	DateTo       string   `schema:"date_to"`                 // REQUIRED  Date to (Available: YYYY-MM-DD)
	OfferID      int      `schema:"offer_id"`                // REQUIRED An offer id
	Timezone     string   `schema:"timezone"`                // Timezone name. Example: “Europe/Berlin” (REQUIRED Timezone)
	Goal         string   `schema:"goal,omitempty"`          // Name based goal
	AffiliateIDs []uint64 `schema:"affiliate_ids,omitempty"` // Affiliates filter. Comma separated int values
	Page         int      `schema:"page,omitempty"`          // Page of stat entities (Default: 1)
	Limit        int      `schema:"limit,omitempty"`         // Limit of stat entities (Default: 100)
} // todo values

// statisticTimeToActionResponse specifies response for TimeToAction.
type statisticTimeToActionResponse struct {
	Data []*TimeToAction `json:"data"`
}

// TimeToAction gets Time to action report.
func (s *StatisticService) TimeToAction(ctx context.Context, opts *StatisticTimeToActionOpts) ([]*TimeToAction, *Response, error) {
	path := "/3.0/stats/time-to-action"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(statisticTimeToActionResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Data, resp, nil
}
