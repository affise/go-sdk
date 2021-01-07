package affise

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Detail struct {
	OfferID    int    `json:"offer_id"`
	PayoutType string `json:"payout_type"`
	Actions    int    `json:"actions"`
	Amount     int    `json:"amount"`
	Comment    string `json:"comment"`
}

type Message struct {
	Number     int      `json:"number"`
	SupplierID string   `json:"supplier_id"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	StartDate  string   `json:"start_date"`
	EndDate    string   `json:"end_date"`
	Status     string   `json:"status"`
	Detail     []Detail `json:"detail"`
	Currency   string   `json:"currency"`
	Comment    string   `json:"comment"`
}

type AdminAdvertiserBillingService struct {
	client *Client
}

type DetailOpts struct {
	OfferID    int    `schema:"offer_id"    json:"offer_id"`    // Offer id
	PayoutType string `schema:"payout_type" json:"payout_type"` // Payout type (Available: RPA,RPS,RPA + RPS,RPC, RPM)
	Actions    int    `schema:"actions"     json:"actions"`     // Actions
	Amount     int    `schema:"amount"      json:"amount"`      // Amount
	Comment    string `schema:"comment"     json:"comment"`     // Comment for detail
}

// AdminAdvertiserBillingListOpts specifies options for List.
type AdminAdvertiserBillingListOpts struct {
	Page      int    `schema:"page,omitempty"`       // Page of entities
	Limit     int    `schema:"limit,omitempty"`      // Limit of entities
	Status    string `schema:"status,omitempty"`     // Status of invoice (Available: paid, unpaid)
	StartDate string `schema:"start_date,omitempty"` // Start date of period
	EndDate   string `schema:"end_date,omitempty"`   // End date of period
}

// adminAdvertiserBillingListResponse specifies response for List.
type adminAdvertiserBillingListResponse struct {
	Message []*Message `json:"message"`
}

// List gets list of invoices.
func (s *AdminAdvertiserBillingService) List(ctx context.Context, opts *AdminAdvertiserBillingListOpts) ([]*Message, *Response, error) {
	path := "/3.0/admin/advertiser-invoices"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAdvertiserBillingListResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Message, resp, nil
}

// adminAdvertiserBillingGetResponse specifies response for Get.
type adminAdvertiserBillingGetResponse struct {
	Message *Message `json:"message"`
}

// Get gets a invoice.
func (s *AdminAdvertiserBillingService) Get(ctx context.Context, number int) (*Message, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/advertiser-invoice/%d", number)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminAdvertiserBillingGetResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Message, resp, nil
}

// AdminAdvertiserBillingCreateOpts specifies options for Create.
type AdminAdvertiserBillingCreateOpts struct {
	SupplierID string       `schema:"supplier_id"`          // REQUIRED Advertiser Id
	StartDate  string       `schema:"start_date,omitempty"` // Start date of invoice period
	EndDate    string       `schema:"end_date,omitempty"`   // End date of invoice period
	Status     string       `schema:"status,omitempty"`     // Invoice status ([paid, unpaid])
	Currency   string       `schema:"currency"`             // REQUIRED One of the active currencies (RUB, USD, EUR etc)
	Comment    string       `schema:"comment,omitempty"`    // Comment
	Details    []DetailOpts `schema:"-"`
}

func (opts *AdminAdvertiserBillingCreateOpts) values() (url.Values, error) {
	u1, err := defaultEncoder.encode(opts)
	if err != nil {
		return nil, err
	}
	u2, err := defaultEncoder.encodeSlice("detail", opts.Details)
	if err != nil {
		return nil, err
	}

	return mergeValues(u1, u2), nil
}

// Create adds new invoice.
func (s *AdminAdvertiserBillingService) Create(ctx context.Context, opts *AdminAdvertiserBillingCreateOpts) (*Response, error) {
	path := "/3.0/admin/advertiser-invoice"

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

// AdminAdvertiserBillingUpdateOpts specifies options for Update.
type AdminAdvertiserBillingUpdateOpts struct {
	SupplierID string       `schema:"supplier_id"`          // REQUIRED Advertiser Id
	StartDate  string       `schema:"start_date,omitempty"` // Start date of invoice period
	EndDate    string       `schema:"end_date,omitempty"`   // End date of invoice period
	Status     string       `schema:"status,omitempty"`     // Invoice status ([paid, unpaid])
	Comment    string       `schema:"comment,omitempty"`    // Comment
	Details    []DetailOpts `schema:"-"`
}

func (opts *AdminAdvertiserBillingUpdateOpts) values() (url.Values, error) {
	u1, err := defaultEncoder.encode(opts)
	if err != nil {
		return nil, err
	}
	u2, err := defaultEncoder.encodeSlice("detail", opts.Details)
	if err != nil {
		return nil, err
	}

	return mergeValues(u1, u2), nil
}

// Update changes an invoice’s data.
func (s *AdminAdvertiserBillingService) Update(ctx context.Context, number int, opts *AdminAdvertiserBillingUpdateOpts) (*Response, error) {
	path := fmt.Sprintf("/3.0/admin/advertiser-invoice/%d", number)

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
