package affise

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Domain struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	UseHTTPS bool   `json:"use_https"`
}

type ExtendedCurrency struct {
	ID         int    `json:"_id"`
	Code       string `json:"code"`
	Active     bool   `json:"active"`
	Default    bool   `json:"default"`
	Rate       int    `json:"rate"`
	MinPayment int    `json:"min_payment"`
	IsCrypto   bool   `json:"is_crypto"`
}

type Comments struct {
	Answers int `json:"answers"`
	Unread  int `json:"unread"`
}

type City struct {
	CountryCode string `json:"country_code"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	RegionCode  string `json:"region_code"`
}

type Ticket struct {
	ID          string    `json:"id"`
	Status      string    `json:"status"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     string    `json:"created"`
	Updated     string    `json:"updated"`
	Attachments []string  `json:"attachments"`
	Comments    Comments  `json:"comments"`
	Partner     Affiliate `json:"partner"`
	Offer       Offer     `json:"offer"`
}

type Pixel struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Code             string `json:"code"`
	CodeType         string `json:"code_type"`
	OfferID          string `json:"offer_id"`
	AffiliateID      uint64 `json:"pid,string"`
	IsActive         string `json:"is_active"`
	ModerationStatus string `json:"moderation_status"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type SmartLinkCategory struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Domain      string `json:"domain"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type AdminOtherService struct {
	client *Client
}

// AdminOtherListCitiesOpts specifies options for ListCities.
type AdminOtherListCitiesOpts struct {
	Q       string   `schema:"q,omitempty"`    // Search query
	Code    []int    `schema:"code,omitempty"` // City codes for filter
	Country []string `schema:"country"`        // REQUIRED Country code. Example : US
}

func (o *AdminOtherListCitiesOpts) values() (url.Values, error) {
	values := url.Values{}
	if o.Q != "" {
		values.Set("q", o.Q)
	}
	if len(o.Country) != 0 {
		values.Set("country", commaSeparatedStrings(o.Country))
	}
	if len(o.Code) != 0 {
		values.Set("code", commaSeparatedInts(o.Code))
	}

	return values, nil
}

// adminOtherListCitiesResponse specifies response for ListCities.
type adminOtherListCitiesResponse struct {
	Cities []*City `json:"cities"`
}

// ListCities gets city list.
func (s *AdminOtherService) ListCities(ctx context.Context, opts *AdminOtherListCitiesOpts) ([]*City, *Response, error) {
	path := "/3.1/cities"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListCitiesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Cities, resp, nil
}

// adminOtherListDevicesResponse specifies response for ListDevices.
type adminOtherListDevicesResponse struct {
	Types []string `json:"types"`
}

// ListDevices gets list of devices.
func (s *AdminOtherService) ListDevices(ctx context.Context) ([]string, *Response, error) {
	path := "/3.1/devices"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListDevicesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Types, resp, nil
}

// adminOtherListBrowsersResponse specifies response for ListBrowsers.
type adminOtherListBrowsersResponse struct {
	Browsers []string `json:"browsers"`
}

// ListBrowsers gets browsers list.
func (s *AdminOtherService) ListBrowsers(ctx context.Context) ([]string, *Response, error) {
	path := "/3.1/browsers"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListBrowsersResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Browsers, resp, nil
}

// AdminOtherListCurrenciesOpts specifies options for ListCurrencies.
type AdminOtherListCurrenciesOpts struct {
	GetOnlyActive int `schema:"get_only_active,omitempty"` // Ignore inactive currencies (Default: 0  Available: 1)
	Extended      int `schema:"extended,omitempty"`        // Extended currencies information (Default: 0  Available: 1)
}

// adminOtherListCurrenciesResponse specifies response for ListCurrencies.
type adminOtherListCurrenciesResponse struct {
	Quotes map[string]json.Number `json:"quotes"`
}

// ListCurrencies gets list of currency.
func (s *AdminOtherService) ListCurrencies(ctx context.Context, opts *AdminOtherListCurrenciesOpts) (map[string]json.Number, *Response, error) {
	path := "/3.0/admin/currency"

	if opts != nil {
		opts.Extended = 0
	}

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListCurrenciesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Quotes, resp, nil
}

// adminOtherListCurrenciesExtendedResponse specifies response for ListCurrencies.
type adminOtherListCurrenciesExtendedResponse struct {
	Quotes []*ExtendedCurrency `json:"quotes"`
}

// ListCurrencies gets extended list of currency.
func (s *AdminOtherService) ListCurrenciesExtended(ctx context.Context, opts *AdminOtherListCurrenciesOpts) ([]*ExtendedCurrency, *Response, error) {
	path := "/3.0/admin/currency"

	if opts != nil {
		opts.Extended = 1
	}

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListCurrenciesExtendedResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Quotes, resp, nil
}

// adminOtherListPaymentSystemsResponse specifies response for ListPaymentSystems.
type adminOtherListPaymentSystemsResponse struct {
	PaymentSystems []*PaymentSystem `json:"payment_systems"`
}

// ListPaymentSystems gets list of payment systems.
func (s *AdminOtherService) ListPaymentSystems(ctx context.Context) ([]*PaymentSystem, *Response, error) {
	path := "/3.0/admin/payment_systems"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListPaymentSystemsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.PaymentSystems, resp, nil
}

// adminOtherListCustomFieldsResponse specifies response for ListCustomFields.
type adminOtherListCustomFieldsResponse struct {
	Fields []*CustomField `json:"fields"`
}

// ListCustomFields gets list of signup settings.
func (s *AdminOtherService) ListCustomFields(ctx context.Context) ([]*CustomField, *Response, error) {
	path := "/3.0/admin/custom_fields"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListCustomFieldsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Fields, resp, nil
}

// adminOtherListDomainsResponse specifies response for ListDomains.
type adminOtherListDomainsResponse struct {
	Domains []*Domain `json:"domains"`
}

// ListDomains gets domains.
func (s *AdminOtherService) ListDomains(ctx context.Context) ([]*Domain, *Response, error) {
	path := "/3.0/admin/domains"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListDomainsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Domains, resp, nil
}

// adminOtherGetTicketResponse specifies response for GetTicket.
type adminOtherGetTicketResponse struct {
	Ticket *Ticket `json:"ticket"`
}

// GetTicket gets ticket.
func (s *AdminOtherService) GetTicket(ctx context.Context, id string) (*Ticket, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/ticket/%s", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherGetTicketResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Ticket, resp, nil
}

// AdminOtherListTicketsOpts specifies options for ListTickets.
type AdminOtherListTicketsOpts struct {
	Page   int    `schema:"page,omitempty"`   // Page of stat entities (Default: 1)
	Limit  int    `schema:"limit,omitempty"`  // Limit of entities (Default: 100)
	Status string `schema:"status,omitempty"` // (Available: open, closed)
}

// adminOtherListTicketsResponse specifies response for ListTickets.
type adminOtherListTicketsResponse struct {
	Tickets []*Ticket `json:"tickets"`
}

// ListTickets gets list of tickets for connection to offers.
func (s *AdminOtherService) ListTickets(ctx context.Context, opts *AdminOtherListTicketsOpts) ([]*Ticket, *Response, error) {
	path := "/3.0/admin/tickets"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListTicketsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Tickets, resp, nil
}

// AdminOtherApproveTicketOpts specifies options for ApproveTicket.
type AdminOtherApproveTicketOpts struct {
	Do string `schema:"do,omitempty"` // What need to do with a ticket (Available: approve, reject)
}

// ApproveTicket approves or rejects ticket for connect affiliate to offer.
func (s *AdminOtherService) ApproveTicket(ctx context.Context, id int, opts *AdminOtherApproveTicketOpts) (*Response, error) {
	path := fmt.Sprintf("/3.0/admin/ticket/%d/offer", id)

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

// adminOtherListPixelsResponse specifies response for ListPixels.
type adminOtherListPixelsResponse struct {
	Pixel map[string]*Pixel `json:"pixel"`
}

// ListPixels gets list of a partner’s pixels.
func (s *AdminOtherService) ListPixels(ctx context.Context, affiliateID uint64) ([]*Pixel, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/pixels/%d", affiliateID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListPixelsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	ret := make([]*Pixel, 0, len(body.Pixel))
	for _, v := range body.Pixel {
		ret = append(ret, v)
	}

	return ret, resp, nil
}

// AdminOtherCreatePixelOpts specifies options for CreatePixel.
type AdminOtherCreatePixelOpts struct {
	AffiliateID      uint64 `schema:"pid"`                         // REQUIRED affiliate’s ID
	OfferID          uint64 `schema:"offer_id"`                    // REQUIRED Offer’s ID
	Name             string `schema:"name"`                        // REQUIRED Name
	Code             string `schema:"code"`                        // REQUIRED Code (Available: <script>…code…</scipt>, <img …>, <iframe src=“…”></iframe>)
	CodeType         string `schema:"code_type"`                   // REQUIRED Code type (Available: javascript, iframe, image)
	IsActive         int    `schema:"is_active,omitempty"`         // Active or not (Available: 0, 1)
	ModerationStatus int    `schema:"moderation_status,omitempty"` // Moderation status (Available: Pending: 0, Rejected: -1, Approved: 1)
}

// adminOtherCreatePixelResponse specifies response for CreatePixel.
type adminOtherCreatePixelResponse struct {
	Pixel *Pixel `json:"pixel"`
}

// CreatePixel adds a partner’s pixel.
func (s *AdminOtherService) CreatePixel(ctx context.Context, opts *AdminOtherCreatePixelOpts) (*Pixel, *Response, error) {
	path := "/3.0/partner/pixel"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherCreatePixelResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Pixel, resp, nil
}

// AdminOtherUpdatePixelOpts specifies options for UpdatePixel.
type AdminOtherUpdatePixelOpts struct {
	Name             string `schema:"name,omitempty"`              // Name
	Code             string `schema:"code,omitempty"`              // Code (Available: <script>…code…</scipt>, <img …>, <iframe src=“…”></iframe>)
	CodeType         string `schema:"code_type,omitempty"`         // Code type (Available: javascript, iframe, image)
	IsActive         int    `schema:"is_active,omitempty"`         // Active or not (Available: 0, 1)
	ModerationStatus int    `schema:"moderation_status,omitempty"` // Moderation status (Available: Pending: 0, Rejected: -1, Approved: 1)
}

// adminOtherUpdatePixelResponse specifies response for UpdatePixel.
type adminOtherUpdatePixelResponse struct {
	Pixel *Pixel `json:"pixel"`
}

// UpdatePixel edits a partner’s pixel.
func (s *AdminOtherService) UpdatePixel(ctx context.Context, id int, opts *AdminOtherUpdatePixelOpts) (*Pixel, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/pixel/%d", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherUpdatePixelResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Pixel, resp, nil
}

// adminOtherDeletePixelResponse specifies response for DeletePixel.
type adminOtherDeletePixelResponse struct {
	Pixel *Pixel `json:"pixel"`
}

// DeletePixel deletes a partner’s pixel.
func (s *AdminOtherService) DeletePixel(ctx context.Context, id int) (*Pixel, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/pixel/%d/remove", id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherDeletePixelResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Pixel, resp, nil
}

// AdminOtherListSmartLinkCategoriesOpts specifies options for ListSmartLinkCategories.
type AdminOtherListSmartLinkCategoriesOpts struct {
	ID   []string `schema:"id,omitempty"`   // SmartLink categories ID (Available only letters (a-f), numbers (0-9). Line length must be 24 characters)
	Name string   `schema:"name,omitempty"` // SmartLink title
}

// adminOtherListSmartLinkCategoriesResponse specifies response for ListSmartLinkCategories.
type adminOtherListSmartLinkCategoriesResponse struct {
	Data []*SmartLinkCategory `json:"data"`
}

// ListSmartLinkCategories gets SmartLink categories list.
func (s *AdminOtherService) ListSmartLinkCategories(ctx context.Context, opts *AdminOtherListSmartLinkCategoriesOpts) ([]*SmartLinkCategory, *Response, error) {
	path := "/3.0/admin/smartlink/categories"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherListSmartLinkCategoriesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Data, resp, nil
}

// AdminOtherCreateSmartLinkCategoryOpts specifies options for CreateSmartLinkCategory.
type AdminOtherCreateSmartLinkCategoryOpts struct {
	Name        string `schema:"name"`                  // REQUIRED Category name
	DomainID    int    `schema:"domain_id,omitempty"`   // (Keep it empty to set the default TDS domain or use domain ID from GET /3.0/admin/domains)
	Description string `schema:"description,omitempty"` // Category description
}

// adminOtherCreateSmartLinkCategoryResponse specifies response for CreateSmartLinkCategory.
type adminOtherCreateSmartLinkCategoryResponse struct {
	Data *SmartLinkCategory `json:"data"`
}

// CreateSmartLinkCategory adds new SmartLink category.
func (s *AdminOtherService) CreateSmartLinkCategory(ctx context.Context, opts *AdminOtherCreateSmartLinkCategoryOpts) (*SmartLinkCategory, *Response, error) {
	path := "/3.0/admin/smartlink/category"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherCreateSmartLinkCategoryResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Data, resp, nil
}

// AdminOtherUpdateSmartLinkCategoryOpts specifies options for UpdateSmartLinkCategory.
type AdminOtherUpdateSmartLinkCategoryOpts struct {
	Name        string `schema:"name,omitempty"`        // Name of category
	DomainID    int    `schema:"domain_id,omitempty"`   // If you stay it as empty will set the default TDS domain. The domain ID from your domain list.
	Description string `schema:"description,omitempty"` // Description of the category
}

// adminOtherUpdateSmartLinkCategoryResponse specifies response for UpdateSmartLinkCategory.
type adminOtherUpdateSmartLinkCategoryResponse struct {
	Data *SmartLinkCategory `json:"data"`
}

// UpdateSmartLinkCategory edits a SmartLink category.
func (s *AdminOtherService) UpdateSmartLinkCategory(ctx context.Context, id string, opts *AdminOtherUpdateSmartLinkCategoryOpts) (*SmartLinkCategory, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/smartlink/category/%s", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherUpdateSmartLinkCategoryResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Data, resp, nil
}

// adminOtherDeleteSmartLinkCategoryResponse specifies response for DeleteSmartLinkCategory.
type adminOtherDeleteSmartLinkCategoryResponse struct {
	Data *SmartLinkCategory `json:"data"`
}

// DeleteSmartLinkCategory removes a SmartLink category.
func (s *AdminOtherService) DeleteSmartLinkCategory(ctx context.Context, id string) (*SmartLinkCategory, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/smartlink/category/%s/remove", id)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminOtherDeleteSmartLinkCategoryResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Data, resp, nil
}

// adminOtherGetSmartLinkOffersCountResponse specifies response for GetSmartLinkOffersCount.
type adminOtherGetSmartLinkOffersCountResponse struct {
	Data struct {
		Count int `json:"count"`
	} `json:"data"`
}

// GetSmartLinkOffersCount adds new SmartLink category.
func (s *AdminOtherService) GetSmartLinkOffersCount(ctx context.Context, id string) (int, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/smartlink/category/%s/offers-count", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return 0, nil, err
	}

	body := new(adminOtherGetSmartLinkOffersCountResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return 0, nil, err
	}

	return body.Data.Count, resp, nil
}
