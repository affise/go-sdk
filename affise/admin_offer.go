package affise

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Source struct {
	ID        string            `json:"id"`
	Title     string            `json:"title"`
	TitleLang map[string]string `json:"title_lang"`
	Allowed   int               `json:"allowed"`
}

type TargetingGroup struct {
	CountryAllow  []string                     `schema:"country[allow],omitempty"`   // list of allowed countries. Example: country[allow][]=US (countries ISO codes)
	CountryDeny   []string                     `schema:"country[deny],omitempty"`    // list of denied
	IPAllow       []string                     `schema:"ip[allow],omitempty"`        // list of allowed ip ranges (IPv4 , IPv6). Available formats : “100.0.0.1” (single IP), “100.0.0.1-100.0.0.255” (IP range), “222.1.1.20/26” (IP/mask)
	IPDeny        []string                     `schema:"ip[deny],omitempty"`         // list of denied ip ranges
	BrowserAllow  []string                     `schema:"browser[allow],omitempty"`   // list of allowed browsers. Example: browser[deny][] = “Edge”
	BrowserDeny   []string                     `schema:"browser[deny],omitempty"`    // list of denied browsers
	BrandAllow    []string                     `schema:"brand[allow],omitempty"`     // list of allowed device brands. Example: brand[deny][] = “SAMTEL”
	BrandDeny     []string                     `schema:"brand[deny],omitempty"`      // list of denied device brands
	DeviceType    []string                     `schema:"device_type,omitempty"`      // list of allowed device types. (“mobile”, “tablet”, “desktop”, “mediahub”, “ereader”, “console”, “tv”, “smartwatch”)
	Connection    []string                     `schema:"connection,omitempty"`       // list of allowed connection types. (“wi-fi”, “cellular”)
	AffiliateID   []uint64                     `schema:"affiliate_id,omitempty"`     // list of affiliates for personal targeting groups.
	RegionAllow   map[string][]int             `schema:"region[allow],omitempty"`    // list of allowed regions for chosen country(ISO). Example: region[allow][US]=33 (region codes)
	RegionDeny    map[string][]int             `schema:"region[deny],omitempty"`     // list of denied regions
	CityAllow     map[string][]int             `schema:"city[allow],omitempty"`      // list of allowed cities for chosen country(ISO). Example: city[allow][US]=57 (city codes)
	CityDeny      map[string][]int             `schema:"city[deny],omitempty"`       // list of denied cities
	ISPAllow      map[string][]string          `schema:"isp[allow],omitempty"`       // list of allowed ISP for chosen country(ISO). Example: isp[allow][US]=Att (ISP list)
	SubAllow      map[string][]string          `schema:"sub[allow],omitempty"`       // list of allowed subs for chosen sub parameter. Example: sub[allow][2][]=“subValue”
	SubDeny       map[string][]string          `schema:"sub[deny],omitempty"`        // list of denied subs for chosen sub parameter.
	SubDenyGroups map[string]map[string]string `schema:"sub[deny_groups],omitempty"` // list of denied sub restricted groups (when is needed to block sub pairs(and more): sub1=“A” + sub2=“B”). Example: sub[deny_groups][0][1]=“A” + sub[deny_groups][0][2]=“B”. To implement “Block traffic if empty sub” option, put empty string in the group : sub[deny_groups][1][8]=“”
	BlockProxy    int                          `schema:"block_proxy,omitempty"`      // enable/disable “Click-level Anti-fraud” feature (0, 1)
	URL           string                       `schema:"url,omitempty"`              // Additional Tracking URL
	OSAllow       []OS                         `schema:"os[allow],omitempty"`        // list of allowed OSes and them versions. To deny specific version should use according comparison operation. (See OS Structure)
	Urls          []URLWeight                  `schema:"urls,omitempty"`             // URLs with weights for traffic redistribution between several track-links
}

type URLWeight struct {
	URL    string `schema:"url,omitempty"`    // Tracking URL
	Weight int    `schema:"weight,omitempty"` // track-link weight (0-100)
}

type OS struct {
	Name       string `json:"name,omitempty"`       // OS name that could be found at OSes list method.
	Comparison string `json:"comparison,omitempty"` // Comparison operation for OS version. (LT, LTE, EQ, GT, GTE)
	Version    string `json:"version,omitempty"`    // OS version for comparison.
}

// Payment item structure.
type Payment struct {
	Partners       []int    `json:"partners,omitempty"` // Array of partner ID, which include payments (It’s available only for personal payments)
	Countries      []string `json:"countries"`          // An array of countries in ISO format (or put empty string to clear existing items)
	CountryExclude bool     `json:"country_exclude"`    // Exclude these countries
	Cities         []City   `json:"cities"`             // An array of id cities (or put empty string to clear existing items)
	Devices        []string `json:"devices"`            // The array of devices. Possible values: mediahub, mobile, ereader, console, tv, tablet, desktop, smartwatch (or put empty string to clear existing items)
	OS             []string `json:"os"`                 // The array of OSes
	Goal           string   `json:"goal"`               // Value targets
	Total          int      `json:"total"`              // The amount of payment
	Revenue        int      `json:"revenue"`            // Payment webmaster
	Currency       string   `json:"currency"`           // Currency (Code in ECB format)
	Type           string   `json:"type"`               // Type of payment. Possible values: fixed, percent, mixed
	Title          string   `json:"title"`
	URL            string   `json:"url"`
	WithRegions    bool     `json:"with_regions"`
}

// Landing structure.
type Landing struct {
	Title      string `json:"title"`       // Title
	URL        string `json:"url"`         // Tracking URL
	URLPreview string `json:"url_preview"` // View URL
	Type       string `json:"type"`        // Type (Possible values: landing, transit; By default: landing)
}

// Strictly identify the operating system.
type OSStrictly struct {
	OS string `json:"os"`
	// OS versions with possible special chars >= or <.
	// Possible values When adding OS targeting to the API offer,
	// the version is a required parameter but can be an empty array value
	Versions string `json:"versions"`
}

// Strictly identify ISP.
type ISP struct {
	Country string `json:"country"` // Country in ISO format
	Name    string `json:"name"`    // Name
}

// todo interface{} field
// Cap item structure.
type Cap struct {
	Period        string      `json:"period,omitempty"`         // Possible values: day, month, all
	Type          string      `json:"type,omitempty"`           // Possible values: budget, conversions, clicks
	Value         json.Number `json:"value,omitempty"`          // The integer value for the type of conversion and the float value for the budget type.
	GoalType      string      `json:"goal_type,omitempty"`      // Values: “all” , “each”, “exact”. “goals” field is mandatory to be specified for “exact” value.
	Goals         interface{} `json:"goals,omitempty"`          // Either specifies goal value or is empty. Empty field requires “goal_type” values of “all”/“each”.
	AffiliateType string      `json:"affiliate_type,omitempty"` // Values: “all” , “each”, “exact”. “affiliates” field is mandatory to be specified for “exact” value.
	Affiliates    []int       `json:"affiliates,omitempty"`     // Either specifies affiliate ID or is empty filed. Empty field requires “affiliate_type” values of “all”/“each”.
	CountryType   string      `json:"country_type,omitempty"`   // Values: “all” , “each”, “exact”. “country” field is mandatory to be specified for “exact” value.
	Country       []string    `json:"country,omitempty"`        // Country codes.
}

// Commission tier item structure.
type CommissionTier struct {
	Timeframe           string      `json:"timeframe"`             //  Possible values: day, week, month, all
	Type                string      `json:"type"`                  //  Possible values: budget, conversions
	Value               json.Number `json:"value"`                 // The integer value for the type of conversion and the float value for the budget type.
	ModifierValue       float64     `json:"modifier_value"`        // The float value.
	ModifierType        string      `json:"modifier_type"`         // Possible values: by_fix, by_percent, to_fix, to_percent.
	Goals               []string    `json:"goals"`                 // Either specifies goal value or is empty.
	TargetGoals         []string    `json:"target_goals"`          // Either specifies target goal value or is empty.
	AffiliateType       string      `json:"affiliate_type"`        // Possible values: all, each, exact. Default: each.
	Affiliates          []int       `json:"affiliates"`            // Either specifies affiliate ID or is empty filed.
	ModifierPaymentType string      `json:"modifier_payment_type"` // Possible values: payout, total, payout_and_total. Default: payout.
	ConversionStatus    []string    `json:"conversion_status"`     // Possible values: confirmed, pending, declined, not_found, hold.
}

type SubAccount struct {
	Value  string     `json:"value"`
	Except CustomBool `json:"except"`
}

type Offer struct {
	ID                           int                   `json:"id"`
	OfferID                      string                `json:"offer_id"`
	Advertiser                   string                `json:"advertiser"`
	ExternalOfferID              string                `json:"external_offer_id"`
	BundleID                     string                `json:"bundle_id"`
	HidePayments                 bool                  `json:"hide_payments"`
	Title                        string                `json:"title"`
	MacroURL                     string                `json:"macro_url"`
	URL                          string                `json:"url"`
	CrossPostbackURL             string                `json:"cross_postback_url"`
	URLPreview                   string                `json:"url_preview"`
	PreviewURL                   string                `json:"preview_url"`
	DomainURL                    string                `json:"domain_url"`
	UseHTTPS                     bool                  `json:"use_https"`
	UseHTTP                      bool                  `json:"use_http"`
	DescriptionLang              map[string]string     `json:"description_lang"`
	Sources                      []Source              `json:"sources"`
	Logo                         string                `json:"logo"`
	LogoSource                   string                `json:"logo_source"`
	Status                       string                `json:"status"`
	Tags                         []string              `json:"tags"`
	Privacy                      string                `json:"privacy"`
	IsTop                        int                   `json:"is_top"`
	Payments                     []Payment             `json:"payments"`
	PartnerPayments              []Payment             `json:"partner_payments"`
	Landings                     []Landing             `json:"landings"`
	StrictlyCountry              int                   `json:"strictly_country"`
	StrictlyOS                   []string              `json:"strictly_os"`
	StrictlyBrands               []string              `json:"strictly_brands"`
	StrictlyConnectionType       string                `json:"strictly_connection_type"`
	IsRedirectOvercap            bool                  `json:"is_redirect_overcap"`
	NoticePercentOvercap         int                   `json:"notice_percent_overcap"`
	HoldPeriod                   int                   `json:"hold_period"`
	Categories                   []string              `json:"categories"`
	FullCategories               []Category            `json:"full_categories"`
	CR                           float64               `json:"cr"`
	EPC                          float64               `json:"epc"`
	Notes                        string                `json:"notes"`
	AllowedIP                    string                `json:"allowed_ip"`
	DisallowedIP                 string                `json:"disallowed_ip"`
	HashPassword                 string                `json:"hash_password"`
	AllowDeeplink                int                   `json:"allow_deeplink"`
	HideReferer                  int                   `json:"hide_referer"`
	StartAt                      string                `json:"start_at"`
	StopAt                       string                `json:"stop_at"`
	AutoOfferConnect             int                   `json:"auto_offer_connect"`
	RequiredApproval             bool                  `json:"required_approval"`
	IsCPI                        bool                  `json:"is_cpi"`
	KPI                          map[string]string     `json:"kpi"`
	SubRestrictions              []map[string]string   `json:"sub_restrictions"`
	Creatives                    []int                 `json:"creatives"`
	CreativesZip                 interface{}           `json:"creatives_zip"`
	SubAccounts                  map[string]SubAccount `json:"sub_accounts"`
	RedirectType                 string                `json:"redirect_type"`
	Caps                         []Cap                 `json:"caps"`
	CommissionTiers              []CommissionTier      `json:"commission_tiers"`
	CapsTimezone                 string                `json:"caps_timezone"`
	StrictlyISP                  []string              `json:"strictly_isp"`
	RestrictionISP               []ISP                 `json:"restriction_isp"`
	StrictlyDevices              []string              `json:"strictly_devices"`
	DisabledChoicePostbackStatus bool                  `json:"disabled_choice_postback_status"`
	UpdatedAt                    string                `json:"updated_at"`
	CreatedAt                    string                `json:"created_at"`
	CapsStatus                   []string              `json:"caps_status"`
	SearchEmptySub               int                   `json:"search_empty_sub"`
	AllowImpressions             bool                  `json:"allow_impressions"`
	SmartlinkCategories          []string              `json:"smartlink_categories"`
	ClickSession                 string                `json:"click_session"`
	MinimalClickSession          string                `json:"minimal_click_session"`
	IoDocument                   bool                  `json:"io_document"`
	UniqIPOnly                   bool                  `json:"uniq_ip_only"`
	RejectNotUniqIP              bool                  `json:"reject_not_uniq_ip"`
}

type Category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type AdminOfferService struct {
	client *Client
}

// adminOfferGetCountOffersResponse specifies response for GetCountOffers.
type adminOfferGetCountOffersResponse struct {
	Count int `json:"count"`
}

// GetCountOffers gets count of offers in status "active".
func (s *AdminOfferService) GetCountOffers(ctx context.Context) (int, *Response, error) {
	path := "/3.0/offers/count"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return 0, nil, err
	}

	body := new(adminOfferGetCountOffersResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return 0, nil, err
	}

	return body.Count, resp, nil
}

// AdminOfferCreateOfferOpts specifies options for CreateOffer.
type AdminOfferCreateOfferOpts struct {
	Title                         string            `schema:"title"`                                      // REQUIRED Title
	Advertiser                    string            `schema:"advertiser"`                                 // REQUIRED Advertiser ID
	URL                           string            `schema:"url"`                                        // REQUIRED Tracking URL
	CrossPostbackURL              string            `schema:"cross_postback_url,omitempty"`               // Cross-postback URL
	MacroURL                      string            `schema:"macro_url,omitempty"`                        // Additional macro
	URLPreview                    string            `schema:"url_preview,omitempty"`                      // View URL
	TrafficbackURL                string            `schema:"trafficback_url,omitempty"`                  // Trafficback URL
	DomainURL                     int               `schema:"domain_url,omitempty"`                       // The domain Id for the tracking URL
	DescriptionLang               []string          `schema:"description_lang,omitempty"`                 // Offer description on specified language. Example: description_lang[en] = ‘English description’
	StopDate                      string            `schema:"stopDate,omitempty"`                         // Stop date (Available: YYYY-MM-DD)
	CreativeFiles                 []string          `schema:"creativeFiles,omitempty"`                    // An array of creative FILES to upload (Available: image/jpeg, image/png, image/gif, application/zip)
	CreativeUrls                  []string          `schema:"creativeUrls,omitempty"`                     // An array of URLs to external creative resources
	CreativeDownloads             []string          `schema:"creativeDownloads,omitempty"`                // An array of URLs to external creative resources for download
	Sources                       []string          `schema:"sources,omitempty"`                          // An array of traffic sources The list of available sources of traffic in the section
	Logo                          string            `schema:"logo,omitempty"`                             // logo File (Available: image/jpeg, image/pjpeg, image/png, image/gif)
	Status                        string            `schema:"status,omitempty"`                           // Offer status (Default: stopped  Available: stopped, active, suspended)
	Tags                          []string          `schema:"tags,omitempty"`                             // Offer tags
	Privacy                       string            `schema:"privacy,omitempty"`                          // Privacy level (Available: public, protected, private)
	IsTop                         int               `schema:"is_top,omitempty"`                           // The top offer (Available: 0, 1)
	IsCpi                         int               `schema:"is_cpi,omitempty"`                           // CPI (Available: 0, 1)
	Payments                      []string          `schema:"payments,omitempty"`                         // Payments array (See Structure)
	PartnerPayments               []string          `schema:"partner_payments,omitempty"`                 // An array of personal paymentsy (See Structure)
	NoticePercentOvercap          int               `schema:"notice_percent_overcap,omitempty"`           // The percentage conversions to achieve the daily limit at which the messages will be sent
	Landings                      []string          `schema:"landings,omitempty"`                         // An array of landings(See Structure)
	StrictlyCountry               int               `schema:"strictly_country,omitempty"`                 // Strictly identify the country (Available: 0, 1)
	StrictlyConnectionType        string            `schema:"strictly_connection_type,omitempty"`         // Strictly identify the connection type. Set a value to empty for choosing the all strictly connection type. (Available: “”, wi-fi, cellular)
	StrictlyOs                    []string          `schema:"strictly_os,omitempty"`                      // Deprecated : use restriction_os
	RestrictionOs                 []string          `schema:"restriction_os,omitempty"`                   // Strictly identify the operating system (See Structure)
	StrictlyDevices               []string          `schema:"strictly_devices,omitempty"`                 // Strictly identify the device (See Possible values)
	StrictlyBrands                []string          `schema:"strictly_brands,omitempty"`                  // Vendors (See Vendors)
	CapsStatus                    []string          `schema:"caps_status,omitempty"`                      // Array of conversion statuses for caps calculation. Available values: “confirmed”, “pending”, “hold”, “not_found”, “declined”
	CapsTimezone                  string            `schema:"caps_timezone,omitempty"`                    // Select timezone of conversions calculating for caps with periods day/month
	EnabledCommissionTiers        int               `schema:"enabled_commission_tiers,omitempty"`         // Enable commission tiers (Available: 0, 1 Default: 0)
	HoldPeriod                    int               `schema:"hold_period,omitempty"`                      // Hold time (Available: between 0 and 60)
	Categories                    []string          `schema:"categories,omitempty"`                       // An array of categories
	Notes                         string            `schema:"notes,omitempty"`                            // Offer notes
	AllowedIP                     string            `schema:"allowed_ip,omitempty"`                       // Allowed IP. Example: 127.0.0.1\n127.0.1.1-127.0.2.1
	AllowDeeplink                 int               `schema:"allow_deeplink,omitempty"`                   // Allow diplinks (Available: 0, 1)
	HideReferer                   int               `schema:"hide_referer,omitempty"`                     // Hide referrer (Available: 0, 1)
	RedirectType                  string            `schema:"redirect_type,omitempty"`                    // Redirect types: http302 - usual http redirect with code 302. Without referrer passing: http302hidden, meta (meta-tag redirect), js (javascript redirect) (http302, http302hidden, js, meta)
	StartAt                       string            `schema:"start_at,omitempty"`                         // Date time of launch (Available: YYYY-MM-DD HH:MM:SS)
	SendEmails                    int               `schema:"send_emails,omitempty"`                      // Send emails to affiliates by offer changing. (Default: 0  Available: 0, 1)
	IsRedirectOvercap             int               `schema:"is_redirect_overcap,omitempty"`              // Send traffic to trafficback by daily overcaps. (Default: 0  Available: 0, 1)
	HidePayments                  int               `schema:"hide_payments,omitempty"`                    // Hide the percentage of contributions to offer for partners if it is the type of Percent payment. (Default: 0  Available: 0, 1)
	ClickSession                  string            `schema:"click_session,omitempty"`                    // Click Session Lifespan  Example: 1y2m3w4d5h6i7s  Scales must be one from: y(year), m(month), w(week), d(day), h(hour), i(minute), s(second) (Default: 1y)
	MinimalClickSession           string            `schema:"minimal_click_session,omitempty"`            // Minimal click session lifespan  Example: 1y2m3w4d5h6i7s  Scales must be one from: y(year), m(month), w(week), d(day), h(hour), i(minute), s(second) (Default: 0s)
	SubAccount1                   string            `schema:"sub_account_1,omitempty"`                    // Allowed sub1 values (Available only letters(a-z), numbers(0-9) and these symbols: ,._-{}+=/:~)
	SubAccount2                   string            `schema:"sub_account_2,omitempty"`                    // Allowed sub2 values (Available only letters(a-z), numbers(0-9) and these symbols: ,._-{}+=/:~)
	SubAccount1Except             int               `schema:"sub_account_1_except,omitempty"`             // Block sub1 values, set only with sub_account_1 (Default: 0  Available: 0, 1)
	SubAccount2Except             int               `schema:"sub_account_2_except,omitempty"`             // Block sub2 values, set only with sub_account_2 (Default: 0  Available: 0, 1)
	SmartlinkCategories           []string          `schema:"smartlink_categories,omitempty"`             // Smartlink category ID. Use /3.0/admin/smartlink/categories to get an ID. Use empty value to remove a Smartlink category from an offer.
	Kpi                           []string          `schema:"kpi,omitempty"`                              // KPI description on specified language. Example: kpi[en] = ‘English text’
	UniqIPOnly                    int               `schema:"uniqIpOnly,omitempty"`                       // Unique IP only flag (Default:  0  Available: 0, 1)
	RejectNotUniqIP               int               `schema:"rejectNotUniqIp,omitempty"`                  // Reject not unique Ip flag (Default:  0  Available: 0, 1)
	StrictlyIsp                   []string          `schema:"strictly_isp,omitempty"`                     // Deprecated : use restriction_isp
	RestrictionIsp                []string          `schema:"restriction_isp,omitempty"`                  // Stricly ISP (See Structure)
	ExternalOfferID               string            `schema:"external_offer_id,omitempty"`                // External offer id
	BundleID                      string            `schema:"bundle_id,omitempty"`                        // Bundle id
	NoteAff                       string            `schema:"note_aff,omitempty"`                         // Note for affiliate
	NoteSales                     string            `schema:"note_sales,omitempty"`                       // Note for sales
	DisallowedIP                  string            `schema:"disallowed_ip,omitempty"`                    // disallowed ip
	HideCaps                      int               `schema:"hide_caps,omitempty"`                        // Hide caps in partner interface (Available: 0, 1)
	SearchEmptySub                int               `schema:"search_empty_sub,omitempty"`                 // Search for an empty sub with this number (Available: 1..8)
	CapsGoalOvercap               string            `schema:"caps_goal_overcap,omitempty"`                // Enabled - When cap for chosen default goal is reached, clicks would be redirected to Trafficback url
	AllowImpressions              int               `schema:"allow_impressions,omitempty"`                // Allow impressions for offer (Available: 0, 1)
	ImpressionsURL                string            `schema:"impressions_url,omitempty"`                  // Impressions destination URL
	ConsiderPersonalTargetingOnly string            `schema:"consider_personal_targeting_only,omitempty"` // (Available: true/false)
	Caps                          []Cap             `schema:"-"`                                          // Caps (See CapStructure)
	CommissionTiers               []CommissionTier  `schema:"-"`                                          // Commission tiers (See CommissionTierStructure). Commission tier list replaces existing list. To prevent a counter reset do not change fields in new list except value and modifier_value. To delete commission tiers set empty field.
	Targeting                     []TargetingGroup  `schema:"-"`                                          // Array of targeting groups (See Structure)
	SubRestrictions               map[string]string `schema:"-"`                                          // Sub restriction pair. Example or structure: sub_restrictions[0][sub1] = ‘sub1_val’, sub_restrictions[0][sub2] = ‘sub2_val’, sub_restrictions[1][sub1] = ‘sub2_val’, etc..
}

func (opts *AdminOfferCreateOfferOpts) values() (_ url.Values, err error) {
	// todo encode Caps, CommissionTiers, Targeting, SubRestrictions
	return defaultEncoder.encode(opts)
}

// adminOfferCreateOfferResponse specifies response for CreateOffer.
type adminOfferCreateOfferResponse struct {
	Offer *Offer `json:"offer"`
}

// CreateOffer creates a new offer.
func (s *AdminOfferService) CreateOffer(ctx context.Context, opts *AdminOfferCreateOfferOpts) (*Offer, *Response, error) {
	path := "/3.0/admin/offer"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOfferCreateOfferResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Offer, resp, nil
}

// AdminOfferUpdateOfferOpts specifies options for UpdateOffer.
type AdminOfferUpdateOfferOpts struct {
	Title                         string            `schema:"title,omitempty"`                            // Title
	Advertiser                    string            `schema:"advertiser,omitempty"`                       // Advertiser ID
	URL                           string            `schema:"url,omitempty"`                              // Tracking URL
	CrossPostbackURL              string            `schema:"cross_postback_url,omitempty"`               // Cross-postback URL
	MacroURL                      string            `schema:"macro_url,omitempty"`                        // Additional macro
	URLPreview                    string            `schema:"url_preview,omitempty"`                      // View URL
	TrafficbackURL                string            `schema:"trafficback_url,omitempty"`                  // Trafficback URL
	DomainURL                     int               `schema:"domain_url,omitempty"`                       // The domain Id for the tracking URL
	DescriptionLang               []string          `schema:"description_lang,omitempty"`                 // Offer description on specified language. Example: description_lang[en] = ‘English description’
	Kpi                           []string          `schema:"kpi,omitempty"`                              // KPI description on specified language. Example: kpi[en] = ‘English text’
	StopDate                      string            `schema:"stopDate,omitempty"`                         // Stop date (Available: YYYY-MM-DD)
	CreativeFiles                 []string          `schema:"creativeFiles,omitempty"`                    // An array of creative FILES to upload (Available: image/jpeg, image/png, image/gif, application/zip)
	CreativeUrls                  []string          `schema:"creativeUrls,omitempty"`                     // An array of URLs to external creative resources
	CreativeDownloads             []string          `schema:"creativeDownloads,omitempty"`                // An array of URLs to external creative resources for download
	Sources                       []string          `schema:"sources,omitempty"`                          // An array of traffic sources The list of available sources of traffic in the section
	Logo                          string            `schema:"logo,omitempty"`                             // logo File (Available: image/jpeg, image/pjpeg, image/png, image/gif)
	Status                        string            `schema:"status,omitempty"`                           // Offer status (Default: stopped  Available: stopped, active, suspended)
	Tags                          []string          `schema:"tags,omitempty"`                             // Offer tags
	Privacy                       string            `schema:"privacy,omitempty"`                          // Privacy level (Available: public, protected, private)
	IsTop                         int               `schema:"is_top,omitempty"`                           // The top offer (Available: 0, 1)
	IsCpi                         int               `schema:"is_cpi,omitempty"`                           // CPI (Available: 0, 1)
	Payments                      []string          `schema:"payments,omitempty"`                         // Payments array (See Structure)
	PartnerPayments               []string          `schema:"partner_payments,omitempty"`                 // An array of personal paymentsy (See add offer)
	NoticePercentOvercap          int               `schema:"notice_percent_overcap,omitempty"`           // The percentage conversions to achieve the daily limit at which the messages will be sent
	Landings                      []string          `schema:"landings,omitempty"`                         // An array of landings(See Structure)
	StrictlyCountry               int               `schema:"strictly_country,omitempty"`                 // Strictly identify the country (Available: 0, 1)
	StrictlyConnectionType        string            `schema:"strictly_connection_type,omitempty"`         // Strictly identify the connection type. Set a value to empty for choosing the all strictly connection type. (Available: “”, wi-fi, cellular)
	StrictlyOs                    []string          `schema:"strictly_os,omitempty"`                      // Deprecated : use restriction_os
	RestrictionOs                 []string          `schema:"restriction_os,omitempty"`                   // Strictly identify the operating system (See add offer)
	StrictlyDevices               []string          `schema:"strictly_devices,omitempty"`                 // Strictly identify the device (See Possible values)
	CapsStatus                    []string          `schema:"caps_status,omitempty"`                      // Array of conversion statuses for caps calculation. Available values: “confirmed”, “pending”, “hold”, “not_found”, “declined”
	CapsTimezone                  string            `schema:"caps_timezone,omitempty"`                    // Select timezone of conversions calculating for caps with periods day/month
	EnabledCommissionTiers        int               `schema:"enabled_commission_tiers,omitempty"`         // Enable commission tiers (Available: 0, 1 Default: 0)
	HoldPeriod                    int               `schema:"hold_period,omitempty"`                      // Hold time (Available: between 0 and 60)
	Categories                    []string          `schema:"categories,omitempty"`                       // An array of categories
	Notes                         string            `schema:"notes,omitempty"`                            // Offer notes
	AllowedIP                     string            `schema:"allowed_ip,omitempty"`                       // Allowed IP. Example: 127.0.0.1\n127.0.1.1-127.0.2.1
	AllowDeeplink                 int               `schema:"allow_deeplink,omitempty"`                   // Allow diplinks (Available: 0, 1)
	HideReferer                   int               `schema:"hide_referer,omitempty"`                     // Hide referrer. Deprecated: use redirect_type (Available: 0, 1)
	RedirectType                  string            `schema:"redirect_type,omitempty"`                    // Redirect types: http302 - usual http redirect with code 302. Without referrer passing: http302hidden, meta (meta-tag redirect), js (javascript redirect) (http302, http302hidden, js, meta)
	StartAt                       string            `schema:"start_at,omitempty"`                         // Date time of launch (Available: YYYY-MM-DD HH:MM:SS)
	SendEmails                    int               `schema:"send_emails,omitempty"`                      // Send emails to affiliates by offer changing. (Default: 0  Available: 0, 1)
	IsRedirectOvercap             int               `schema:"is_redirect_overcap,omitempty"`              // Send traffic to trafficback by daily overcaps. (Default: 0  Available: 0, 1)
	HidePayments                  int               `schema:"hide_payments,omitempty"`                    // Hide the percentage of contributions to offer for partners if it is the type of Percent payment. (Default: 0  Available: 0, 1)
	ClickSession                  string            `schema:"click_session,omitempty"`                    // Click Session Lifespan  Example: 1y2m3w4d5h6i7s  Scales must be one from: y(year), m(month), w(week), d(day), h(hour), i(minute), s(second) (Default: 1y)
	MinimalClickSession           string            `schema:"minimal_click_session,omitempty"`            // Minimal click session lifespan  Example: 1y2m3w4d5h6i7s  Scales must be one from: y(year), m(month), w(week), d(day), h(hour), i(minute), s(second) (Default: 0s)
	SubAccount1                   string            `schema:"sub_account_1,omitempty"`                    // Sub1 list, separated by commas
	SubAccount2                   string            `schema:"sub_account_2,omitempty"`                    // Sub2 list, separated by commas
	SubAccount1Except             int               `schema:"sub_account_1_except,omitempty"`             // Except Sub1 list set only with sub_account_1 (Default: 0  Available: 0, 1)
	SubAccount2Except             int               `schema:"sub_account_2_except,omitempty"`             // Except Sub2 list set only with sub_account_2 (Default: 0  Available: 0, 1)
	SmartlinkCategories           []string          `schema:"smartlink_categories,omitempty"`             // Smartlink category ID. Use /3.0/admin/smartlink/categories to get an ID. Use empty value to remove a Smartlink category from an offer.
	UniqIPOnly                    int               `schema:"uniqIpOnly,omitempty"`                       // Unique IP only flag (Default:  0  Available: 0, 1)
	RejectNotUniqIP               int               `schema:"rejectNotUniqIp,omitempty"`                  // Reject not unique Ip flag (Default:  0  Available: 0, 1)
	StrictlyIsp                   []string          `schema:"strictly_isp,omitempty"`                     // Deprecated : use restriction_isp
	RestrictionIsp                []string          `schema:"restriction_isp,omitempty"`                  // Stricly ISP (See Structure)
	ExternalOfferID               string            `schema:"external_offer_id,omitempty"`                // External offer id
	BundleID                      string            `schema:"bundle_id,omitempty"`                        // Bundle id
	HideCaps                      int               `schema:"hide_caps,omitempty"`                        // Hide caps in partner interface (Available: 0, 1)
	SearchEmptySub                int               `schema:"search_empty_sub,omitempty"`                 // Search for an empty sub with this number (Available: 1..8)
	CapsGoalOvercap               string            `schema:"caps_goal_overcap,omitempty"`                // Enabled - When cap for chosen default goal is reached, clicks would be redirected to Trafficback url
	AllowImpressions              int               `schema:"allow_impressions,omitempty"`                // Allow impressions for offer (Available: 0, 1)
	ImpressionsURL                string            `schema:"impressions_url,omitempty"`                  // Impressions destination URL
	ConsiderPersonalTargetingOnly string            `schema:"consider_personal_targeting_only,omitempty"` // (Available: true/false)
	Caps                          []Cap             `schema:"-"`                                          // Caps (See CapStructure)
	CommissionTiers               []CommissionTier  `schema:"-"`                                          // Commission tiers (See CommissionTierStructure). Commission tier list replaces existing list. To prevent a counter reset do not change fields in new list except value and modifier_value. To delete commission tiers set empty field.
	Targeting                     []TargetingGroup  `schema:"-"`                                          // Array of targeting groups (See Structure)
	SubRestrictions               map[string]string `schema:"-"`                                          // Sub restriction pair. Example or structure: sub_restrictions[0][sub1] = ‘sub1_val’, sub_restrictions[0][sub2] = ‘sub2_val’, sub_restrictions[1][sub1] = ‘sub2_val’, etc..
}

func (opts *AdminOfferUpdateOfferOpts) values() (url.Values, error) {
	// todo encode Caps, CommissionTiers, Targeting, SubRestrictions
	return defaultEncoder.encode(opts)
}

// adminOfferUpdateOfferResponse specifies response for UpdateOffer.
type adminOfferUpdateOfferResponse struct {
	Offer *Offer `json:"offer"`
}

// UpdateOffer changes an offer settings.
func (s *AdminOfferService) UpdateOffer(ctx context.Context, id int, opts *AdminOfferUpdateOfferOpts) (*Offer, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/offer/%d", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOfferUpdateOfferResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Offer, resp, nil
}

// AdminOfferDeleteOfferOpts specifies options for DeleteOffer.
type AdminOfferDeleteOfferOpts struct {
	OfferID []int `schema:"offer_id"` // REQUIRED
}

func (o *AdminOfferDeleteOfferOpts) values() (url.Values, error) {
	res := url.Values{}
	for i, n := range o.OfferID {
		res.Set(fmt.Sprintf("offer_id[%d]", i), strconv.Itoa(n))
	}

	return res, nil
}

// DeleteOffer deletes the offer.
func (s *AdminOfferService) DeleteOffer(ctx context.Context, opts *AdminOfferDeleteOfferOpts) (*Response, error) {
	path := "/3.0/admin/offer/delete"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// adminOfferListSourcesResponse specifies response for ListSources.
type adminOfferListSourcesResponse struct {
	Sources []*Source `json:"sources"`
}

// ListSources gets list of sources.
func (s *AdminOfferService) ListSources(ctx context.Context) ([]*Source, *Response, error) {
	path := "/3.0/admin/offer/sources"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOfferListSourcesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Sources, resp, nil
}

// AdminOfferCreateSourceOpts specifies options for CreateSource.
type AdminOfferCreateSourceOpts struct {
	TitleLang map[string]string `schema:"title_lang"` // REQUIRED  Key-value pair of title on different languages (Available keys: ru, en, es, ka, vi)
}

func (o *AdminOfferCreateSourceOpts) values() (url.Values, error) {
	return defaultEncoder.encodeMap("title_lang", o.TitleLang)
}

// adminOfferCreateSourceResponse specifies response for CreateSource.
type adminOfferCreateSourceResponse struct {
	Source *Source `json:"source"`
}

// CreateSource creates a source.
func (s *AdminOfferService) CreateSource(ctx context.Context, opts *AdminOfferCreateSourceOpts) (*Source, *Response, error) {
	path := "/3.0/admin/offer/source"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOfferCreateSourceResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Source, resp, nil
}

// AdminOfferUpdateSourceOpts specifies options for UpdateSource.
type AdminOfferUpdateSourceOpts struct {
	TitleLang map[string]string `schema:"title_lang"` // REQUIRED  Key-value pair of title on different languages (Available keys: ru, en, es, ka, vi)
}

func (o *AdminOfferUpdateSourceOpts) values() (url.Values, error) {
	return defaultEncoder.encodeMap("title_lang", o.TitleLang)
}

// adminOfferUpdateSourceResponse specifies response for UpdateSource.
type adminOfferUpdateSourceResponse struct {
	Source *Source `json:"source"`
}

// UpdateSource updates a source.
func (s *AdminOfferService) UpdateSource(ctx context.Context, id string, opts *AdminOfferUpdateSourceOpts) (*Source, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/offer/source/%s", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOfferUpdateSourceResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Source, resp, nil
}

// adminOfferDeleteSourceResponse specifies response for DeleteSource.
type adminOfferDeleteSourceResponse struct {
	Source *Source `json:"source"`
}

// DeleteSource deletes a source by ID.
func (s *AdminOfferService) DeleteSource(ctx context.Context, id string) (*Source, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/offer/source/%s", id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOfferDeleteSourceResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Source, resp, nil
}

// AdminOfferCreateCategoryOpts specifies options for CreateCategory.
type AdminOfferCreateCategoryOpts struct {
	Title string `schema:"title"` // REQUIRED Category title
}

// adminOfferCreateCategoryResponse specifies response for CreateCategory.
type adminOfferCreateCategoryResponse struct {
	Category *Category `json:"category"`
}

// CreateCategory adds new category.
func (s *AdminOfferService) CreateCategory(ctx context.Context, opts *AdminOfferCreateCategoryOpts) (*Category, *Response, error) {
	path := "/3.0/admin/category"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOfferCreateCategoryResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Category, resp, nil
}

// AdminOfferUpdateCategoryOpts specifies options for UpdateCategory.
type AdminOfferUpdateCategoryOpts struct {
	Title string `schema:"title"` // REQUIRED Category title
}

// adminOfferUpdateCategoryResponse specifies response for UpdateCategory.
type adminOfferUpdateCategoryResponse struct {
	Category *Category `json:"category"`
}

// POST /3.0/admin/category/{ID}.
func (s *AdminOfferService) UpdateCategory(ctx context.Context, id string, opts *AdminOfferUpdateCategoryOpts) (*Category, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/category/%s", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOfferUpdateCategoryResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Category, resp, nil
}

// AdminOfferEnableAffiliateOpts specifies options for EnableAffiliate.
type AdminOfferEnableAffiliateOpts struct {
	OfferID     []int  `schema:"offer_id"`         // REQUIRED
	AffiliateID uint64 `schema:"pid"`              // REQUIRED Affiliate ID
	Notice      int    `schema:"notice,omitempty"` // Send notice to affiliate (Default: 1  Available: 0 or 1)
}

func (o *AdminOfferEnableAffiliateOpts) values() (url.Values, error) {
	values := url.Values{}
	values.Set("pid", strconv.FormatUint(o.AffiliateID, 10))
	values.Set("notice", strconv.Itoa(o.Notice))
	for i, v := range o.OfferID {
		key := fmt.Sprintf("offer_id[%d]", i)
		values.Set(key, strconv.Itoa(v))
	}

	return values, nil
}

// EnableAffiliate connections an affiliate to offer.
func (s *AdminOfferService) EnableAffiliate(ctx context.Context, opts *AdminOfferEnableAffiliateOpts) (*Response, error) {
	path := "/3.0/offer/enable-affiliate"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AdminOfferDisableAffiliateOpts specifies options for DisableAffiliate.
type AdminOfferDisableAffiliateOpts struct {
	OfferID     []int  `schema:"offer_id"`         // REQUIRED
	AffiliateID uint64 `schema:"pid"`              // REQUIRED Affiliate ID
	Notice      int    `schema:"notice,omitempty"` // Send notice to affiliate (Default: 1  Available: 0 or 1)
}

func (o *AdminOfferDisableAffiliateOpts) values() (url.Values, error) {
	values := url.Values{}
	values.Set("pid", strconv.FormatUint(o.AffiliateID, 10))
	values.Set("notice", strconv.Itoa(o.Notice))
	for i, v := range o.OfferID {
		key := fmt.Sprintf("offer_id[%d]", i)
		values.Set(key, strconv.Itoa(v))
	}

	return values, nil
}

// DisableAffiliate disconnects an affiliate from offer.
func (s *AdminOfferService) DisableAffiliate(ctx context.Context, opts *AdminOfferDisableAffiliateOpts) (*Response, error) {
	path := "/3.0/offer/disable-affiliate"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AdminOfferMassUpdateOffersOpts specifies options for MassUpdateOffers.
type AdminOfferMassUpdateOffersOpts struct {
	OfferID []int  `schema:"offer_id"`          // REQUIRED
	Status  string `schema:"status,omitempty"`  // Status (Available: active stopped suspended)
	Privacy string `schema:"privacy,omitempty"` // Privacy level (Available: public protected private)
}

func (o *AdminOfferMassUpdateOffersOpts) values() (url.Values, error) {
	values := url.Values{}
	if o.Status != "" {
		values.Set("status", o.Status)
	}
	if o.Privacy != "" {
		values.Set("privacy", o.Privacy)
	}
	for i, v := range o.OfferID {
		key := fmt.Sprintf("offer_id[%d]", i)
		values.Set(key, strconv.Itoa(v))
	}

	return values, nil
}

// MassUpdateOffers updates offer`s status.
func (s *AdminOfferService) MassUpdateOffers(ctx context.Context, opts *AdminOfferMassUpdateOffersOpts) (*Response, error) {
	path := "/3.0/admin/offer/mass-update"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DisableAffiliates disconnects all affiliates from private or protected offer.
func (s *AdminOfferService) DisableAffiliates(ctx context.Context, offerID string) (*Response, error) {
	path := fmt.Sprintf("/3.0/admin/offer/%s/disable-affiliates", offerID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil, true)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DisableOffers disconnects all private or protected offers from affiliate.
func (s *AdminOfferService) DisableOffers(ctx context.Context, affiliateID uint64) (*Response, error) {
	path := fmt.Sprintf("/3.0/admin/affiliate/%d/disable-offers", affiliateID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil, true)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AdminOfferRemoveCreativeOpts specifies options for RemoveCreative.
type AdminRemoveOfferCreativesOpts struct {
	Creatives []int `schema:"creatives"` // REQUIRED Creative IDs
}

func (o *AdminRemoveOfferCreativesOpts) values() (url.Values, error) {
	res := url.Values{}
	for i, n := range o.Creatives {
		res.Set(fmt.Sprintf("creatives[%d]", i), strconv.Itoa(n))
	}

	return res, nil
}

// adminOfferRemoveCreativeResponse specifies response for RemoveCreative.
type adminRemoveOfferCreativesResponse struct {
	Removed []int `json:"removed"`
}

// RemoveOfferCreatives removes creative from offer by creative id.
func (s *AdminOfferService) RemoveOfferCreatives(ctx context.Context, offerID string, opts *AdminRemoveOfferCreativesOpts) ([]int, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/offer/%s/remove-creative", offerID)

	req, err := s.client.NewRequestOpts(ctx, http.MethodDelete, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminRemoveOfferCreativesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Removed, resp, nil
}
