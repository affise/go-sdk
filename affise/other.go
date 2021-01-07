package affise

import (
	"context"
	"fmt"
	"net/http"
)

type OtherService struct {
	client *Client
}

type Country struct {
	Code string `json:"code"` // Country in ISO format
	Name string `json:"name"` // Name
}

type Region struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
}

// OtherListISPOpts specifies options for ListISP.
type OtherListISPOpts struct {
	Country string `schema:"country"`     // REQUIRED Country code. Example: “US”
	Q       string `schema:"q,omitempty"` // Search query
}

// otherListISPResponse specifies response for ListISP.
type otherListISPResponse struct {
	ISPs []*ISP `json:"isps"`
}

// ListISP gets ISP list.
func (s *OtherService) ListISP(ctx context.Context, opts *OtherListISPOpts) ([]*ISP, *Response, error) {
	path := "/3.1/isp"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(otherListISPResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.ISPs, resp, nil
}

// otherListCountriesResponse specifies response for ListCountries.
type otherListCountriesResponse struct {
	Countries []*Country `json:"countries"`
}

// ListCountries gets countries list.
func (s *OtherService) ListCountries(ctx context.Context) ([]*Country, *Response, error) {
	path := "/3.1/countries"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(otherListCountriesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Countries, resp, nil
}

// OtherListRegionsOpts specifies options for ListRegions.
type OtherListRegionsOpts struct {
	Country string `schema:"country"` // REQUIRED Country code. Example: “US”
}

// otherListRegionsResponse specifies response for ListRegions.
type otherListRegionsResponse struct {
	Regions []*Region `json:"regions"`
}

// ListRegions gets region list.
func (s *OtherService) ListRegions(ctx context.Context, opts *OtherListRegionsOpts) ([]*Region, *Response, error) {
	path := "/3.1/regions"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(otherListRegionsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Regions, resp, nil
}

// otherListConnectionTypesResponse specifies response for ListConnectionTypes.
type otherListConnectionTypesResponse struct {
	Types []string `json:"types"`
}

// ListConnectionTypes gets connection types list.
func (s *OtherService) ListConnectionTypes(ctx context.Context) ([]string, *Response, error) {
	path := "/3.1/connection-types"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(otherListConnectionTypesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Types, resp, nil
}

// OtherListVendorsOpts specifies options for ListVendors.
type OtherListVendorsOpts struct {
	Q string `schema:"q,omitempty"` // Search query
}

// otherListVendorsResponse specifies response for ListVendors.
type otherListVendorsResponse struct {
	Vendors []string `json:"vendors"`
}

// ListVendors gets vendors list.
func (s *OtherService) ListVendors(ctx context.Context, opts *OtherListVendorsOpts) ([]string, *Response, error) {
	path := "/3.1/vendors"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(otherListVendorsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Vendors, resp, nil
}

// otherListOSResponse specifies response for ListOS.
type otherListOSResponse struct {
	Oses map[string]string
}

// ListOS gets oses list.
func (s *OtherService) ListOS(ctx context.Context) (map[string]string, *Response, error) {
	path := "/3.1/oses"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(otherListOSResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Oses, resp, nil
}

// otherListOSVersionsResponse specifies response for ListOSVersions.
type otherListOSVersionsResponse struct {
	Versions []string `json:"versions"`
}

// ListOSVersions gets os versions list.
func (s *OtherService) ListOSVersions(ctx context.Context, os string) ([]string, *Response, error) {
	path := fmt.Sprintf("/3.1/oses/%s", os)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(otherListOSVersionsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Versions, resp, nil
}
