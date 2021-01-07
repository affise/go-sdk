package affise

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Preset struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Permissions *Permissions `json:"permissions"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

type AdminPresetService struct {
	client *Client
}

// AdminPresetListOpts specifies options for List.
type AdminPresetListOpts struct {
	Page  int `schema:"page,omitempty"`  // page with results (Available: >=1; 1)
	Limit int `schema:"limit,omitempty"` // results per page (Available: >=1; 50)
}

// adminPresetListResponse specifies response for List.
type adminPresetListResponse struct {
	Presets []*Preset `json:"presets"`
}

// List gets list of presets.
func (s *AdminPresetService) List(ctx context.Context, opts *AdminPresetListOpts) ([]*Preset, *Response, error) {
	path := "/3.1/presets"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminPresetListResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Presets, resp, nil
}

// AdminPresetCreateOpts specifies options for Create.
type AdminPresetCreateOpts struct {
	Name        string       `schema:"name"`        // Preset name REQUIRED (String)
	Permissions *Permissions `schema:"permissions"` // REQUIRED Permissions for preset (full scope)
	Type        string       `schema:"type"`        // REQUIRED reset type (affiliate_manager; account_manager; eq=common_manager)
}

// adminPresetCreateResponse specifies response for Create.
type adminPresetCreateResponse struct {
	Preset *Preset `json:"preset"`
}

// Create creates preset using JSON dataset.
func (s *AdminPresetService) Create(ctx context.Context, opts *AdminPresetCreateOpts) (*Preset, *Response, error) {
	path := "/3.1/presets"

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(opts)
	if err != nil {
		return nil, nil, fmt.Errorf("encode opts err: %w", err)
	}

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, buf, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminPresetCreateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Preset, resp, nil
}

// AdminPresetUpdateOpts specifies options for Update.
type AdminPresetUpdateOpts struct {
	Name        string       `schema:"name,omitempty"` // Preset name (String)
	Permissions *Permissions `schema:"permissions"`    // REQUIRED Permissions for update
}

// adminPresetUpdateResponse specifies response for Update.
type adminPresetUpdateResponse struct {
	Preset *Preset `json:"preset"`
}

// Update updates preset using JSON dataset.
func (s *AdminPresetService) Update(ctx context.Context, id string, opts *AdminPresetUpdateOpts) (*Preset, *Response, error) {
	path := fmt.Sprintf("/3.1/presets/%s", id)

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(opts)
	if err != nil {
		return nil, nil, fmt.Errorf("encode opts err: %w", err)
	}

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, buf, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminPresetUpdateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Preset, resp, nil
}

// Delete deletes preset by ID.
func (s *AdminPresetService) Delete(ctx context.Context, id string) (*Response, error) {
	path := fmt.Sprintf("/3.1/presets/%s", id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil, true)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
