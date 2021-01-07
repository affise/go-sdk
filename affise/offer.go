package affise

import (
	"context"
	"fmt"
	"net/http"
)

type OfferService struct {
	client *Client
}

// OfferListOpts specifies options for List.
type OfferListOpts struct {
	Q          string   `schema:"q,omitempty"`          // Search by title and id
	IDs        []string `schema:"ids,omitempty"`        // Search by string offer ID
	IntID      []int    `schema:"int_id,omitempty"`     // Search by int offer ID
	Countries  []string `schema:"countries,omitempty"`  // Array of offers countries(ISO)
	OS         []string `schema:"os,omitempty"`         // OS (Available: web, wp, ios, android)
	Categories []string `schema:"categories,omitempty"` // Array of offers categories
	Sort       []string `schema:"sort,omitempty"`       // Sort offers. Sample sort[id]=asc, sort[title]=desc. (Available: id, title, cr, epc, is_top, created, revenue, daily_cap, total_cap)
	Page       int      `schema:"page,omitempty"`       // Page of offers
	Limit      int      `schema:"limit,omitempty"`      // Count offers by page
	Status     []string `schema:"status,omitempty"`     // ONLY FOR ADMIN (Default: active  Available: active, stopped, suspended)
	Advertiser []string `schema:"advertiser,omitempty"` // ONLY FOR ADMIN Advertiser ID
	Privacy    []int    `schema:"privacy,omitempty"`    // ONLY FOR ADMIN Privacy filter: Public(0), Premoderated(1), Private(2)
	UpdatedAt  string   `schema:"updated_at,omitempty"` // Get offers that have been updated from this date (format YYYY-MM-DD)
	IsTop      int      `schema:"is_top,omitempty"`     // Get TOP-offers (Available: 0, 1)
	BundleID   string   `schema:"bundle_id,omitempty"`  // Search by bundle id
}

// offerListResponse specifies response for List.
type offerListResponse struct {
	Offers []*Offer `json:"offers"`
}

// Offers list.
func (s *OfferService) List(ctx context.Context, opts *OfferListOpts) ([]*Offer, *Response, error) {
	path := "/3.0/offers"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(offerListResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Offers, resp, nil
}

// offerGetResponse specifies response for Get.
type offerGetResponse struct {
	Offer *Offer `json:"offer"`
}

// Offer by id.
func (s *OfferService) Get(ctx context.Context, id int) (*Offer, *Response, error) {
	path := fmt.Sprintf("/3.0/offer/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(offerGetResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Offer, resp, nil
}

// OfferListCategoriesOpts specifies options for ListCategories.
type OfferListCategoriesOpts struct {
	Page  int `schema:"page,omitempty"`  // Page of entities
	Limit int `schema:"limit,omitempty"` // Limit of entities
}

// offerListCategoriesResponse specifies response for ListCategories.
type offerListCategoriesResponse struct {
	Categories []*Category `json:"categories"`
}

// Categories.
func (s *OfferService) ListCategories(ctx context.Context, opts *OfferListCategoriesOpts) ([]*Category, *Response, error) {
	path := "/3.0/offer/categories"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(offerListCategoriesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Categories, resp, nil
}
