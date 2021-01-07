package affise

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type BalanceItem struct {
	Balance   json.Number `json:"balance"`
	Hold      json.Number `json:"hold"`
	Available json.Number `json:"available"`
}

type Balance map[string]BalanceItem

type Manager struct {
	ID        string   `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	WorkHours string   `json:"work_hours"`
	Email     string   `json:"email"`
	Skype     string   `json:"skype"`
	APIKey    string   `json:"api_key"`
	Roles     []string `json:"roles"`
	UpdatedAt string   `json:"updated_at"`
}

type PaymentSystem struct {
	ID        int         `json:"id"`
	LangLabel string      `json:"lang_label"`
	Currency  string      `json:"currency,omitempty"`
	Fields    interface{} `json:"fields"` // todo map[string]string or []CustomField
}

type CustomField struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	Required  bool   `json:"required,omitempty"`
	LangLabel string `json:"lang_label,omitempty"`
	// todo interface{} fields
	Label interface{} `json:"label,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Affiliate struct {
	ID             uint64                `json:"id"`
	CreatedAt      string                `json:"created_at"`
	UpdatedAt      string                `json:"updated_at"`
	Email          string                `json:"email"`
	Login          string                `json:"login"`
	RefPercent     string                `json:"ref_percent"`
	Name           string                `json:"name"`
	Notes          string                `json:"notes"`
	Manager        Manager               `json:"manager"`
	Status         string                `json:"status"`
	PaymentSystems []PaymentSystem       `json:"payment_systems"`
	CustomFields   []CustomField         `json:"customFields"`
	Balance        Balance               `json:"balance"`
	OffersCount    int                   `json:"offersCount"`
	APIKey         string                `json:"api_key"`
	Address1       string                `json:"address_1"`
	Address2       string                `json:"address_2"`
	City           string                `json:"city"`
	Country        string                `json:"country"`
	ZipCode        string                `json:"zip_code"`
	Phone          string                `json:"phone"`
	Ref            string                `json:"ref"`
	SubAccounts    map[string]SubAccount `json:"sub_accounts"`
	ContactPerson  string                `json:"contactPerson"`
}

type Postback struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Status      string `json:"status"`
	Goal        string `json:"goal"`
	Created     string `json:"created"`
	UpdatedAt   string `json:"updated_at"`
	Forced      string `json:"forced"`
	AffiliateID uint64 `json:"pid,string"`
}

type AdminAffiliateService struct {
	client *Client
}

// adminAffiliateGetResponse specifies response for Get.
type adminAffiliateGetResponse struct {
	Partner *Affiliate `json:"partner"`
}

// Get gets affiliate.
func (s *AdminAffiliateService) Get(ctx context.Context, id uint64) (*Affiliate, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/partner/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateGetResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Partner, resp, nil
}

// AdminAffiliateListPartnersOpts specifies options for ListPartners.
type AdminAffiliateListPartnersOpts struct {
	ID            []uint64 `schema:"id,omitempty"`             // Search by affiliate IDs
	WithBalance   int      `schema:"with_balance,omitempty"`   // Show partners with balance (Available: 1)
	Limit         int      `schema:"limit,omitempty"`          // Limit of entities
	Page          int      `schema:"page,omitempty"`           // Page of entities
	UpdatedAt     string   `schema:"updated_at,omitempty"`     // Get partners that have been updated from this date (format YYYY-MM-DD)
	StatusPartner string   `schema:"status_partner,omitempty"` // Filter (Available  0 - Inactive, 1 - Active, 2 - Banned, 3 - On moderation)
}

// adminAffiliateListPartnersResponse specifies response for ListPartners.
type adminAffiliateListPartnersResponse struct {
	Partners []*Affiliate `json:"partners"`
}

// ListPartners gets list of a partners.
func (s *AdminAffiliateService) ListPartners(ctx context.Context, opts *AdminAffiliateListPartnersOpts) ([]*Affiliate, *Response, error) {
	path := "/3.0/admin/partners"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateListPartnersResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Partners, resp, nil
}

type PaymentSystemOpts struct {
	SystemID int               `schema:"system_id"` // Integer ID of partners systems
	Currency string            `schema:"currency"`  // String the currency code.
	Fields   map[string]string `schema:"fields"`    // An array of Advanced fields. For example {‘Integer key’: ‘String value’} where the key is ID from /3.0/admin/payment_systems and value it’s your value of this field.
}

// AdminAffiliateCreateOpts specifies options for Create.
type AdminAffiliateCreateOpts struct {
	Email             string              `schema:"email"`                          // REQUIRED Partners e-mail
	Password          string              `schema:"password"`                       // REQUIRED Partners password
	Country           string              `schema:"country"`                        // REQUIRED Country ISO name
	Login             string              `schema:"login,omitempty"`                // Company name
	ContactPerson     string              `schema:"contact_person,omitempty"`       // Contact person
	RefPercent        string              `schema:"ref_percent,omitempty"`          // Percentage of referral program
	Notes             string              `schema:"notes,omitempty"`                // Notes
	Status            string              `schema:"status,omitempty"`               // Partners status (Available: ‘not active’, ‘active’, ‘banned’, ‘on moderation’)
	ManagerID         string              `schema:"manager_id,omitempty"`           // Manager id
	CustomFields      []string            `schema:"custom_fields,omitempty"`        // An array of custom fields (See /admin/custom_fields)
	Ref               int                 `schema:"ref,omitempty"`                  // Referral partner
	SubAccount1       string              `schema:"sub_account_1,omitempty"`        // Sub1 list, separated by commas
	SubAccount2       string              `schema:"sub_account_2,omitempty"`        // Sub2 list, separated by commas
	SubAccount1Except int                 `schema:"sub_account_1_except,omitempty"` // Except Sub1 list (Default: 0  Available: 0, 1)
	SubAccount2Except int                 `schema:"sub_account_2_except,omitempty"` // Except Sub2 list (Default: 0  Available: 0, 1)
	Notify            int                 `schema:"notify,omitempty"`               // Send welcome email to affiliate  login field should be set (Default: 0  Available: 0, 1)
	TipaltiPayeeID    int                 `schema:"tipalti_payee_id,omitempty"`     // Tipalti Payee ID
	Tags              []string            `schema:"tags,omitempty"`                 // An array of affiliates’ tags
	PaymentSystems    []PaymentSystemOpts `schema:"-"`                              // An array of payments (See Structure and /admin/payment_systems)
}

func (opts *AdminAffiliateCreateOpts) values() (url.Values, error) {
	u, err := defaultEncoder.encode(opts)
	if err != nil {
		return nil, err
	}

	for i := range opts.PaymentSystems {
		ps := opts.PaymentSystems[i]
		u.Set(fmt.Sprintf("payment_systems[%d][currency]", i), ps.Currency)
		u.Set(fmt.Sprintf("payment_systems[%d][system_id]", i), strconv.Itoa(ps.SystemID))
		for key, field := range ps.Fields {
			u.Set(fmt.Sprintf("payment_systems[%d][fields][%s]", i, key), field)
		}
	}

	return u, nil
}

// adminAffiliateCreateResponse specifies response for Create.
type adminAffiliateCreateResponse struct {
	Partner *Affiliate `json:"partner"`
}

// Create adds new partner.
func (s *AdminAffiliateService) Create(ctx context.Context, opts *AdminAffiliateCreateOpts) (*Affiliate, *Response, error) {
	path := "/3.0/admin/partner"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateCreateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Partner, resp, nil
}

// AdminAffiliateUpdateOpts specifies options for Update.
type AdminAffiliateUpdateOpts struct {
	Password          string              `schema:"password,omitempty"`             // Partners password
	Login             string              `schema:"login,omitempty"`                // Company name
	Country           string              `schema:"country,omitempty"`              // Country ISO name
	ContactPerson     string              `schema:"contact_person,omitempty"`       // Contact person
	RefPercent        string              `schema:"ref_percent,omitempty"`          // Percentage of referral program
	Notes             string              `schema:"notes,omitempty"`                // Notes
	Status            string              `schema:"status,omitempty"`               // Partners status (Available: “, ‘not active’, ‘active’, ‘banned’, ‘on moderation’)
	ManagerID         string              `schema:"manager_id,omitempty"`           // Manager id
	CustomFields      []string            `schema:"custom_fields,omitempty"`        // An array of custom fields (See /admin/custom_fields)
	Ref               int                 `schema:"ref,omitempty"`                  // Referral partner
	SubAccount1       string              `schema:"sub_account_1,omitempty"`        // Sub1 list, separated by commas
	SubAccount2       string              `schema:"sub_account_2,omitempty"`        // Sub2 list, separated by commas
	SubAccount1Except int                 `schema:"sub_account_1_except,omitempty"` // Except Sub1 list (Default: 0  Available: 0, 1)
	SubAccount2Except int                 `schema:"sub_account_2_except,omitempty"` // Except Sub2 list (Default: 0  Available: 0, 1)
	TipaltiPayeeID    int                 `schema:"tipalti_payee_id,omitempty"`     // Tipalti Payee ID
	Tags              []string            `schema:"tags,omitempty"`                 // An array of tags (All the previous tags will be overwritten by new ones)
	PaymentSystems    []PaymentSystemOpts `schema:"-"`                              // An array of payments (See the add affiliate method and /admin/payment_systems)
}

func (opts *AdminAffiliateUpdateOpts) values() (url.Values, error) {
	u, err := defaultEncoder.encode(opts)
	if err != nil {
		return nil, err
	}

	for i := range opts.PaymentSystems {
		ps := opts.PaymentSystems[i]
		u.Set(fmt.Sprintf("payment_systems[%d][currency]", i), ps.Currency)
		u.Set(fmt.Sprintf("payment_systems[%d][system_id]", i), strconv.Itoa(ps.SystemID))
		for key, field := range ps.Fields {
			u.Set(fmt.Sprintf("payment_systems[%d][fields][%s]", i, key), field)
		}
	}

	return u, nil
}

// adminAffiliateUpdateResponse specifies response for Update.
type adminAffiliateUpdateResponse struct {
	Partner *Affiliate `json:"partner"`
}

// Update edits a partner.
func (s *AdminAffiliateService) Update(ctx context.Context, id uint64, opts *AdminAffiliateUpdateOpts) (*Affiliate, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/partner/%d", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateUpdateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Partner, resp, nil
}

// AdminAffiliateMassUpdateOpts specifies options for MassUpdate.
type AdminAffiliateMassUpdateOpts struct {
	ID        []uint64 `schema:"id"`                   // REQUIRED Affiliate IDs
	ManagerID string   `schema:"manager_id,omitempty"` // Manager ID
	Status    string   `schema:"status,omitempty"`     // Status (Available: ‘not active’, ‘active’, ‘banned’, ‘on moderation’)
}

// MassUpdate updates status and manager.
func (s *AdminAffiliateService) MassUpdate(ctx context.Context, opts *AdminAffiliateMassUpdateOpts) (*Response, error) {
	path := "/3.0/admin/partners/mass-update"

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

// adminAffiliateChangePasswordResponse specifies response for ChangePassword.
type adminAffiliateChangePasswordResponse struct {
	Partner *Affiliate `json:"partner"`
}

// ChangePassword changes a partner’s password.
func (s *AdminAffiliateService) ChangePassword(ctx context.Context, id uint64) (*Affiliate, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/partner/password/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateChangePasswordResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Partner, resp, nil
}

// AdminAffiliateAddPostbackOpts specifies options for AddPostback.
type AdminAffiliateAddPostbackOpts struct {
	OfferID     int    `schema:"offer_id,omitempty"` // Offer ID (missed parameter means creation of global postback)
	URL         string `schema:"url"`                // REQUIRED Example: http://affise.com
	Status      string `schema:"status,omitempty"`   // Postback status (Available: by_creating, confirmed, pending, declined, hold, not_found)
	Goal        string `schema:"goal,omitempty"`     // Postback goal (value)
	AffiliateID uint64 `schema:"pid"`                // REQUIRED
}

// adminAffiliateAddPostbackResponse specifies response for AddPostback.
type adminAffiliateAddPostbackResponse struct {
	Postback *Postback `json:"postback"`
}

// AddPostback adds a partner’s postback.
func (s *AdminAffiliateService) AddPostback(ctx context.Context, opts *AdminAffiliateAddPostbackOpts) (*Postback, *Response, error) {
	path := "/3.0/partner/postback"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateAddPostbackResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Postback, resp, nil
}

// AdminAffiliateEditPostbackOpts specifies options for EditPostback.
type AdminAffiliateEditPostbackOpts struct {
	URL    string `schema:"url"`              // REQUIRED Example: http://affise.com
	Status string `schema:"status,omitempty"` // Postback status (Available: by_creating, confirmed, pending, declined, hold, not_found)
	Goal   string `schema:"goal,omitempty"`   // Postback goal (value)
}

// adminAffiliateEditPostbackResponse specifies response for EditPostback.
type adminAffiliateEditPostbackResponse struct {
	Postback *Postback `json:"postback"`
}

// EditPostback edits a partner’s postback.
func (s *AdminAffiliateService) EditPostback(ctx context.Context, id int, opts *AdminAffiliateEditPostbackOpts) (*Postback, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/postback/%d", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateEditPostbackResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Postback, resp, nil
}

// adminAffiliateDeletePostbackResponse specifies response for DeletePostback.
type adminAffiliateDeletePostbackResponse struct {
	Postback *Postback `json:"postback"`
}

// DeletePostback deletes a partner’s postback.
func (s *AdminAffiliateService) DeletePostback(ctx context.Context, id int) (*Postback, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/postback/%d/remove", id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateDeletePostbackResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Postback, resp, nil
}

// AdminAffiliateDeletePostbacksByAffiliatesOpts specifies options for DeletePostbacksByAffiliates.
type AdminAffiliateDeletePostbacksByAffiliatesOpts struct {
	IDs []int `schema:"ids,omitempty"`
}

func (o *AdminAffiliateDeletePostbacksByAffiliatesOpts) values() (url.Values, error) {
	s := make([]string, 0, len(o.IDs))
	for _, v := range o.IDs {
		s = append(s, strconv.Itoa(v))
	}
	ret := url.Values{}
	ret.Add("ids", strings.Join(s, ","))

	return ret, nil
}

// DeletePostbacksByAffiliates Deletes partners postbacks by affiliates ids.
func (s *AdminAffiliateService) DeletePostbacksByAffiliates(ctx context.Context, opts *AdminAffiliateDeletePostbacksByAffiliatesOpts) (*Response, error) {
	path := "/3.0/partner/postbacks/by-affiliates"

	req, err := s.client.NewRequestOpts(ctx, http.MethodDelete, path, opts, nil, true)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AdminAffiliateDeletePostbacksByOffersOpts specifies options for DeletePostbacksByOffers.
type AdminAffiliateDeletePostbacksByOffersOpts struct {
	IDs []int `schema:"ids,omitempty"`
}

func (o *AdminAffiliateDeletePostbacksByOffersOpts) values() (url.Values, error) {
	s := make([]string, 0, len(o.IDs))
	for _, v := range o.IDs {
		s = append(s, strconv.Itoa(v))
	}
	ret := url.Values{}
	ret.Add("ids", strings.Join(s, ","))

	return ret, nil
}

// DeletePostbacksByOffers delete partners postbacks by offers ids.
func (s *AdminAffiliateService) DeletePostbacksByOffers(ctx context.Context, opts *AdminAffiliateDeletePostbacksByOffersOpts) (*Response, error) {
	path := "/3.0/partner/postbacks/by-offers"

	req, err := s.client.NewRequestOpts(ctx, http.MethodDelete, path, opts, nil, true)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AdminAffiliateListPostbacksOpts specifies options for ListPostbacks.
type AdminAffiliateListPostbacksOpts struct {
	AffiliateID uint64 `schema:"partner_id"`      // REQUIRED
	Limit       int    `schema:"limit,omitempty"` // Limit of entities
	Page        int    `schema:"page,omitempty"`  // Page of entities
}

// adminAffiliateListPostbacksResponse specifies response for ListPostbacks.
type adminAffiliateListPostbacksResponse struct {
	Postbacks []*Postback `json:"postbacks"`
}

// ListPostbacks gets a list of partner postbacks.
func (s *AdminAffiliateService) ListPostbacks(ctx context.Context, opts *AdminAffiliateListPostbacksOpts) ([]*Postback, *Response, error) {
	path := "/3.0/admin/postbacks"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateListPostbacksResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Postbacks, resp, nil
}

// adminAffiliateChangeAPIKeyResponse specifies response for ChangeAPIKey.
type adminAffiliateChangeAPIKeyResponse struct {
	User *User `json:"user"`
}

// ChangeAPIKey changes partner api key.
func (s *AdminAffiliateService) ChangeAPIKey(ctx context.Context) (*User, *Response, error) {
	path := "/3.1/partner/api_key"

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateChangeAPIKeyResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.User, resp, nil
}

// AdminAffiliateUpdateLocaleOpts specifies options for UpdateLocale.
type AdminAffiliateUpdateLocaleOpts struct {
	Lang     string `schema:"lang,omitempty"`     // Language
	Timezone string `schema:"timezone,omitempty"` // Timezone
}

// UpdateLocale updates a partner’s locale.
func (s *AdminAffiliateService) UpdateLocale(ctx context.Context, affiliateID uint64, opts *AdminAffiliateUpdateLocaleOpts) (*Response, error) {
	path := fmt.Sprintf("/3.0/admin/partner/%d/locale", affiliateID)

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

// adminAffiliateGetReferralsResponse specifies response for GetReferrals.
type adminAffiliateGetReferralsResponse struct {
	Referrals map[string]*Affiliate
}

func (a *adminAffiliateGetReferralsResponse) slice() []*Affiliate {
	ret := make([]*Affiliate, 0, len(a.Referrals))
	for _, v := range a.Referrals {
		ret = append(ret, v)
	}

	return ret
}

// GetReferrals gets referrals by partner ID.
func (s *AdminAffiliateService) GetReferrals(ctx context.Context, affiliateID uint64) ([]*Affiliate, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/partner/%d/referrals", affiliateID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAffiliateGetReferralsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.slice(), resp, nil
}
