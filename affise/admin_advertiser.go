package affise

import (
	"context"
	"fmt"
	"net/http"
)

type ManagerObj struct {
	ID        string   `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Skype     string   `json:"skype"`
	Roles     []string `json:"roles"`
	APIKey    string   `json:"api_key"`
	CreatedAt string   `json:"created_at"`
}

type Advertiser struct {
	ID                            string                `json:"id"`
	Title                         string                `json:"title"`
	Contact                       string                `json:"contact"`
	Email                         string                `json:"email"`
	URL                           string                `json:"url"`
	Manager                       string                `json:"manager"`
	ManagerObj                    *ManagerObj           `json:"manager_obj"`
	AllowedIP                     []string              `json:"allowed_ip"`
	DisallowedIP                  []string              `json:"disallowed_ip"`
	Skype                         string                `json:"skype"`
	Note                          string                `json:"note"`
	Address1                      string                `json:"address_1"`
	Address2                      string                `json:"address_2"`
	City                          string                `json:"city"`
	Country                       string                `json:"country"`
	ZipCode                       string                `json:"zip_code"`
	VatCode                       string                `json:"vat_code"`
	SubAccounts                   map[string]SubAccount `json:"sub_accounts"`
	HashPassword                  string                `json:"hash_password"`
	ConsiderPersonalTargetingOnly bool                  `json:"consider_personal_targeting_only"`
}

type AdminAdvertiserService struct {
	client *Client
}

// adminAdvertiserGetResponse specifies response for Get.
type adminAdvertiserGetResponse struct {
	Advertiser *Advertiser `json:"advertiser"`
}

// Get gets advertiser.
func (s *AdminAdvertiserService) Get(ctx context.Context, id string) (*Advertiser, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/advertiser/%s", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAdvertiserGetResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Advertiser, resp, nil
}

// AdminAdvertiserListOpts specifies options for List.
type AdminAdvertiserListOpts struct {
	Page      int    `schema:"page,omitempty"`       // Page of entities
	Limit     int    `schema:"limit,omitempty"`      // Limit of entities
	Order     string `schema:"order,omitempty"`      // Sort by field (Default: _id  Available: _id, title, email)
	OrderType string `schema:"orderType,omitempty"`  // Sorting order (Default: asc  Available: desc, asc)
	UpdatedAt string `schema:"updated_at,omitempty"` // Get advertisers that have been updated from this date (format YYYY-MM-DD)
}

// adminAdvertiserListResponse specifies response for List.
type adminAdvertiserListResponse struct {
	Advertisers []*Advertiser `json:"advertisers"`
}

// List gets a list of advertisers.
func (s *AdminAdvertiserService) List(ctx context.Context, opts *AdminAdvertiserListOpts) ([]*Advertiser, *Response, error) {
	path := "/3.0/admin/advertisers"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAdvertiserListResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Advertisers, resp, nil
}

// AdminAdvertiserCreateOpts specifies options for Create.
type AdminAdvertiserCreateOpts struct {
	Title                         string   `schema:"title"`                                      // REQUIRED Company name
	Contact                       string   `schema:"contact,omitempty"`                          // Contact person name
	Skype                         string   `schema:"skype,omitempty"`                            // IM/Skype
	Manager                       string   `schema:"manager,omitempty"`                          // Manager ID
	URL                           string   `schema:"url,omitempty"`                              // Site Url
	Email                         string   `schema:"email,omitempty"`                            // Email
	AllowedIP                     string   `schema:"allowed_ip,omitempty"`                       // Allowed IP. Example: 127.0.0.1\n127.0.1.1-127.0.2.1
	Address1                      string   `schema:"address_1,omitempty"`                        // Main address string
	Address2                      string   `schema:"address_2,omitempty"`                        // Additional address
	City                          string   `schema:"city,omitempty"`                             // City name
	Country                       string   `schema:"country,omitempty"`                          // Country ISO name
	ZipCode                       string   `schema:"zip_code,omitempty"`                         // Zip code
	VatCode                       string   `schema:"vat_code,omitempty"`                         // Vat code
	SubAccount1                   string   `schema:"sub_account_1,omitempty"`                    // Allowed sub1 values (Available only letters(a-z), numbers(0-9) and these symbols: ,._-{}+=/:~)
	SubAccount2                   string   `schema:"sub_account_2,omitempty"`                    // Allowed sub2 values (Available only letters(a-z), numbers(0-9) and these symbols: ,._-{}+=/:~)
	SubAccount1Except             int      `schema:"sub_account_1_except,omitempty"`             // Block sub1 values (Default: 0  Available: 0, 1)
	SubAccount2Except             int      `schema:"sub_account_2_except,omitempty"`             // Block sub2 values (Default: 0  Available: 0, 1)
	ConsiderPersonalTargetingOnly string   `schema:"consider_personal_targeting_only,omitempty"` // (Available: true/false)
	Tags                          []string `schema:"tags,omitempty"`                             // An array of advertiser’s tags
}

// adminAdvertiserCreateResponse specifies response for Create.
type adminAdvertiserCreateResponse struct {
	Advertiser *Advertiser `json:"advertiser"`
}

// Create adds new advertiser.
func (s *AdminAdvertiserService) Create(ctx context.Context, opts *AdminAdvertiserCreateOpts) (*Advertiser, *Response, error) {
	path := "/3.0/admin/advertiser"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAdvertiserCreateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Advertiser, resp, nil
}

// AdminAdvertiserUpdateOpts specifies options for Update.
type AdminAdvertiserUpdateOpts struct {
	Title                         string   `schema:"title,omitempty"`                            // Company name
	Contact                       string   `schema:"contact,omitempty"`                          // Contact person name
	Skype                         string   `schema:"skype,omitempty"`                            // IM/Skype
	Manager                       string   `schema:"manager,omitempty"`                          // Manager ID
	URL                           string   `schema:"url,omitempty"`                              // Site Url
	Email                         string   `schema:"email,omitempty"`                            // Email
	AllowedIP                     string   `schema:"allowed_ip,omitempty"`                       // Allowed IP. Example: 127.0.0.1\n127.0.1.1-127.0.2.1
	Note                          string   `schema:"note,omitempty"`                             // Note
	Address1                      string   `schema:"address_1,omitempty"`                        // Main address string
	Address2                      string   `schema:"address_2,omitempty"`                        // Additional address
	City                          string   `schema:"city,omitempty"`                             // City name
	Country                       string   `schema:"country,omitempty"`                          // Country ISO name
	ZipCode                       string   `schema:"zip_code,omitempty"`                         // Zip code
	VatCode                       string   `schema:"vat_code,omitempty"`                         // Vat code
	SubAccount1                   string   `schema:"sub_account_1,omitempty"`                    // Sub1 list, separated by commas
	SubAccount2                   string   `schema:"sub_account_2,omitempty"`                    // Sub2 list, separated by commas
	SubAccount1Except             int      `schema:"sub_account_1_except,omitempty"`             // Except Sub1 list (Default: 0  Available: 0, 1)
	SubAccount2Except             int      `schema:"sub_account_2_except,omitempty"`             // Except Sub2 list (Default: 0  Available: 0, 1)
	ConsiderPersonalTargetingOnly string   `schema:"consider_personal_targeting_only,omitempty"` // (Available: true/false)
	Tags                          []string `schema:"tags,omitempty"`                             // An array of tags (All the previous tags will be overwritten by new ones)
}

// adminAdvertiserUpdateResponse specifies response for Update.
type adminAdvertiserUpdateResponse struct {
	Advertiser *Advertiser `json:"advertiser"`
}

// Update changes an advertiser’s data.
func (s *AdminAdvertiserService) Update(ctx context.Context, id string, opts *AdminAdvertiserUpdateOpts) (*Advertiser, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/advertiser/%s", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAdvertiserUpdateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Advertiser, resp, nil
}

// SendPassword changes an advertiser password and send it by email.
func (s *AdminAdvertiserService) SendPassword(ctx context.Context, id string) (*Response, error) {
	path := fmt.Sprintf("/3.0/admin/advertiser/%s/sendpass", id)

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

// AdminAdvertiserEnableAffiliateOpts specifies options for EnableAffiliate.
type AdminAdvertiserEnableAffiliateOpts struct {
	AdvertisersID []string `schema:"advertisers_id"` // REQUIRED Array of advertiser IDs to connect
	AffiliateID   uint64   `schema:"pid"`            // REQUIRED affiliate ID
}

// EnableAffiliate un-puts affiliate from blacklist for specified advertisers.
func (s *AdminAdvertiserService) EnableAffiliate(ctx context.Context, opts *AdminAdvertiserEnableAffiliateOpts) (*Response, error) {
	path := "/3.0/admin/advertiser/enable-affiliate"

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

// AdminAdvertiserDisableAffiliateOpts specifies options for DisableAffiliate.
type AdminAdvertiserDisableAffiliateOpts struct {
	AdvertisersID []string `schema:"advertisers_id"` // REQUIRED Array of advertiser IDs to connect
	AffiliateID   uint64   `schema:"pid"`            // REQUIRED affiliate ID
}

// DisableAffiliate puts affiliate to blacklist for specified advertisers.
func (s *AdminAdvertiserService) DisableAffiliate(ctx context.Context, opts *AdminAdvertiserDisableAffiliateOpts) (*Response, error) {
	path := "/3.0/admin/advertiser/disable-affiliate"

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
