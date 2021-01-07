package affise

import (
	"context"
	"net/http"
	"net/url"
)

type AdminConversionService struct {
	client *Client
}

type AdminConversionEditOpts struct {
	IDs      []string `json:"ids"                schema:"ids"`                // REQUIRED
	Status   string   `json:"status,omitempty"   schema:"status,omitempty"`   // (Available: confirmed, pending, declined, not_found, hold)
	Currency string   `json:"currency,omitempty" schema:"currency,omitempty"` // Example: usd
	Payouts  int      `json:"payouts,omitempty"  schema:"payouts,omitempty"`
	Revenue  int      `json:"revenue,omitempty"  schema:"revenue,omitempty"`
	Comment  string   `json:"comment,omitempty"  schema:"comment,omitempty"` // Text a comment
}

type adminConversionEditResponse struct {
	Data AdminConversionEditOpts `json:"data"`
}

type ConversionEdit struct {
	AdminConversionEditOpts
}

// Edit a conversion.
func (s *AdminConversionService) Edit(ctx context.Context,
	opts *AdminConversionEditOpts) (*ConversionEdit, *Response, error) {
	path := "/3.0/admin/conversion/edit"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminConversionEditResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	conv := &ConversionEdit{AdminConversionEditOpts: body.Data}

	return conv, resp, err
}

type AdminConversionImportOpts struct {
	Offer        int    `json:"offer"                    schema:"offer"`                    // REQUIRED Offer id
	AffiliateID  uint64 `json:"pid"                      schema:"pid"`                      // REQUIRED Partner id
	ActionID     string `json:"action_id,omitempty"      schema:"action_id,omitempty"`      // publisher conversion id
	ClickID      string `json:"click_id,omitempty"       schema:"click_id,omitempty"`       // Click ID
	Goal         int    `json:"goal,omitempty"           schema:"goal,omitempty"`           // goal number
	IP           string `json:"ip,omitempty"             schema:"ip,omitempty"`             // visitor ip
	UA           string `json:"ua,omitempty"             schema:"ua,omitempty"`             // visitor user-agent
	Comment      string `json:"comment,omitempty"        schema:"comment,omitempty"`        // comment
	Sum          int    `json:"sum,omitempty"            schema:"sum,omitempty"`            // payouts amount for conversion (for percent payment type)
	Status       string `json:"status,omitempty"         schema:"status,omitempty"`         // (Available: confirmed, pending, declined, not_found, hold)
	CustomField1 string `json:"custom_field_1,omitempty" schema:"custom_field_1,omitempty"` // custom field 1
	CustomField2 string `json:"custom_field_2,omitempty" schema:"custom_field_2,omitempty"` // custom field 2
	CustomField3 string `json:"custom_field_3,omitempty" schema:"custom_field_3,omitempty"` // custom field 3
	CustomField4 string `json:"custom_field_4,omitempty" schema:"custom_field_4,omitempty"` // custom field 4
	CustomField5 string `json:"custom_field_5,omitempty" schema:"custom_field_5,omitempty"` // custom field 5
	CustomField6 string `json:"custom_field_6,omitempty" schema:"custom_field_6,omitempty"` // custom field 6
	CustomField7 string `json:"custom_field_7,omitempty" schema:"custom_field_7,omitempty"` // custom field 7
}

type adminConversionImportResponse struct {
	Data AdminConversionImportOpts `json:"data"`
}

type ConversionImport struct {
	AdminConversionImportOpts
}

// Import a single conversion.
func (s *AdminConversionService) Import(ctx context.Context,
	opts *AdminConversionImportOpts) (*ConversionImport, *Response, error) {
	path := "/3.0/admin/conversion/import"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminConversionImportResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	conv := &ConversionImport{AdminConversionImportOpts: body.Data}

	return conv, resp, err
}

type AdminConversionImportListOpts struct {
	List []AdminConversionImportOpts `schema:"list"`
}

func (a *AdminConversionImportListOpts) values() (url.Values, error) {
	return defaultEncoder.encodeSlice("list", a.List)
}

type adminConversionImportListResponse struct {
	Data struct {
		List []AdminConversionImportOpts `json:"list"`
	} `json:"data"`
}

// ImportList imports multiple conversions.
func (s *AdminConversionService) ImportList(ctx context.Context,
	opts *AdminConversionImportListOpts) ([]*ConversionImport, *Response, error) {
	path := "/3.0/admin/conversions/import"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminConversionImportListResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	ret := make([]*ConversionImport, 0, len(body.Data.List))
	for i := range body.Data.List {
		v := body.Data.List[i]
		conv := &ConversionImport{AdminConversionImportOpts: v}
		ret = append(ret, conv)
	}

	return ret, resp, err
}
