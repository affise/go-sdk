package affise

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type NewsItem struct {
	ID struct {
		ID string `json:"$id"`
	} `json:"_id"`
	Title     string `json:"title"`
	SmallDesc string `json:"small_desc"`
	Desc      string `json:"desc"`
	Status    int    `json:"status"`
	CreatedAt struct {
		Sec  int `json:"sec"`
		Usec int `json:"usec"`
	} `json:"created_at"`
}

type AffiliateService struct {
	client *Client
}

// affiliateMeResponse specifies response for Me.
type affiliateMeResponse struct {
	User *User `json:"user"`
}

// Me gets partner own data.
func (s *AffiliateService) Me(ctx context.Context) (*User, *Response, error) {
	path := "/3.1/partner/me"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateMeResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.User, resp, nil
}

// AffiliateListOffersOpts specifies options for ListOffers.
type AffiliateListOffersOpts struct {
	Q          string   `schema:"q,omitempty"`          // Search by title and id
	IDs        []string `schema:"ids,omitempty"`        // Search by string offer ID
	IntID      []int    `schema:"int_id,omitempty"`     // Search by int offer ID
	Countries  []string `schema:"countries,omitempty"`  // Array of offers countries(ISO)
	Categories []string `schema:"categories,omitempty"` // Array of offers categories
	Sort       []string `schema:"sort,omitempty"`       // Sort offers. Sample sort[id]=asc, sort[title]=desc. You can sort offers by one of (id, title, cr, epc)
	Page       int      `schema:"page,omitempty"`       // Page of offers
	Limit      int      `schema:"limit,omitempty"`      // Count offers by page
}

// affiliateListOffersResponse specifies response for ListOffers.
type affiliateListOffersResponse struct {
	Offers []*Offer `json:"offers"`
}

// ListOffers gets list of a live offers.
func (s *AffiliateService) ListOffers(ctx context.Context, opts *AffiliateListOffersOpts) ([]*Offer, *Response, error) {
	path := "/3.0/partner/offers"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateListOffersResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Offers, resp, nil
}

// AffiliateListLiveOffersOpts specifies options for ListLiveOffers.
type AffiliateListLiveOffersOpts struct {
	Q          string   `schema:"q,omitempty"`          // Search by title and id
	IDs        []string `schema:"ids,omitempty"`        // Search by string offer ID
	IntID      []int    `schema:"int_id,omitempty"`     // Search by int offer ID
	Countries  []string `schema:"countries,omitempty"`  // Array of offers countries(ISO)
	Categories []string `schema:"categories,omitempty"` // Array of offers categories
	Sort       []string `schema:"sort,omitempty"`       // Sort offers. Sample sort[id]=asc, sort[title]=desc. You can sort offers by one of (id, title, cr, epc)
	Page       int      `schema:"page,omitempty"`       // Page of offers
	Limit      int      `schema:"limit,omitempty"`      // Count offers by page
}

// affiliateListLiveOffersResponse specifies response for ListLiveOffers.
type affiliateListLiveOffersResponse struct {
	Offers []*Offer `json:"offers"`
}

// ListLiveOffers gets list of a live offers.
func (s *AffiliateService) ListLiveOffers(ctx context.Context, opts *AffiliateListLiveOffersOpts) ([]*Offer, *Response, error) {
	path := "/3.0/partner/live-offers"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateListLiveOffersResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Offers, resp, nil
}

// AffiliateActivationOfferOpts specifies options for ActivationOffer.
type AffiliateActivationOfferOpts struct {
	OfferID int    `schema:"offer_id"` // REQUIRED
	Comment string `schema:"comment"`  // REQUIRED
}

// ActivationOffer connects to an offer.
func (s *AffiliateService) ActivationOffer(ctx context.Context, opts *AffiliateActivationOfferOpts) (*Response, error) {
	path := "/3.0/partner/activation/offer"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, false)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AffiliateCreatePostbackOpts specifies options for CreatePostback.
type AffiliateCreatePostbackOpts struct {
	AffiliateID uint64 `schema:"pid"`                // REQUIRED
	OfferID     int    `schema:"offer_id,omitempty"` // Offer ID (missed parameter means creation of global postback)
	URL         string `schema:"url"`                // REQUIRED Example: http://affise.com
	Status      string `schema:"status,omitempty"`   // Postback status (Available: by_creating, confirmed, pending, declined, hold, not_found)
	Goal        string `schema:"goal,omitempty"`     // Postback goal (value)
}

// affiliateCreatePostbackResponse specifies response for CreatePostback.
type affiliateCreatePostbackResponse struct {
	Postback *Postback `json:"postback"`
}

// CreatePostback adds postback.
func (s *AffiliateService) CreatePostback(ctx context.Context, opts *AffiliateCreatePostbackOpts) (*Postback, *Response, error) {
	path := "/3.0/partner/postback"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateCreatePostbackResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Postback, resp, nil
}

// AffiliateUpdatePostbackOpts specifies options for UpdatePostback.
type AffiliateUpdatePostbackOpts struct {
	URL    string `schema:"url"`              // REQUIRED Example: http://affise.com
	Status string `schema:"status,omitempty"` // Postback status (Available: by_creating, confirmed, pending, declined, hold, not_found)
	Goal   string `schema:"goal,omitempty"`   // Postback goal (value)
}

// affiliateUpdatePostbackResponse specifies response for UpdatePostback.
type affiliateUpdatePostbackResponse struct {
	Postback *Postback `json:"postback"`
}

// UpdatePostback edits postback.
func (s *AffiliateService) UpdatePostback(ctx context.Context, id int, opts *AffiliateUpdatePostbackOpts) (*Postback, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/postback/%d", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateUpdatePostbackResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Postback, resp, nil
}

// affiliateDeletePostbackResponse specifies response for DeletePostback.
type affiliateDeletePostbackResponse struct {
	Postback *Postback `json:"postback"`
}

// DeletePostback deletes postback.
func (s *AffiliateService) DeletePostback(ctx context.Context, id int) (*Postback, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/postback/%d/remove", id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateDeletePostbackResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Postback, resp, nil
}

// AffiliateDeletePostbacksByAffiliatesOpts specifies options for DeletePostbacksByAffiliates.
type AffiliateDeletePostbacksByAffiliatesOpts struct {
	IDs []int `schema:"ids,omitempty"`
}

func (opts *AffiliateDeletePostbacksByAffiliatesOpts) values() (url.Values, error) {
	u := url.Values{}
	u.Set("ids", commaSeparatedInts(opts.IDs))

	return u, nil
}

// DELETE /3.0/partner/postbacks/by-affiliates.
func (s *AffiliateService) DeletePostbacksByAffiliates(ctx context.Context, opts *AffiliateDeletePostbacksByAffiliatesOpts) (*Response, error) {
	path := "/3.0/partner/postbacks/by-affiliates"

	req, err := s.client.NewRequestOpts(ctx, http.MethodDelete, path, opts, nil, false)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AffiliateDeletePostbacksByOffersOpts specifies options for DeletePostbacksByOffers.
type AffiliateDeletePostbacksByOffersOpts struct {
	IDs []int `schema:"ids,omitempty"`
}

func (opts *AffiliateDeletePostbacksByOffersOpts) values() (url.Values, error) {
	u := url.Values{}
	u.Set("ids", commaSeparatedInts(opts.IDs))

	return u, nil
}

// DeletePostbacksByOffers deletes postbacks by offers ids.
func (s *AffiliateService) DeletePostbacksByOffers(ctx context.Context, opts *AffiliateDeletePostbacksByOffersOpts) (*Response, error) {
	path := "/3.0/partner/postbacks/by-offers"

	req, err := s.client.NewRequestOpts(ctx, http.MethodDelete, path, opts, nil, false)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// AffiliateListNewsOpts specifies options for ListNews.
type AffiliateListNewsOpts struct {
	Limit int `schema:"limit,omitempty"` // (Available: max 100  Default: 10)
	Skip  int `schema:"skip,omitempty"`  // Offset (Default: 0)
	Fixed int `schema:"fixed,omitempty"` // 1 - pinned, 0 - not pinned (Available: 1, 0)
}

// affiliateListNewsResponse specifies response for ListNews.
type affiliateListNewsResponse struct {
	Items map[string]*NewsItem `json:"items"`
}

func (a *affiliateListNewsResponse) slice() []*NewsItem {
	ret := make([]*NewsItem, 0, len(a.Items))
	for _, v := range a.Items {
		ret = append(ret, v)
	}

	return ret
}

// ListNews gets news list.
func (s *AffiliateService) ListNews(ctx context.Context, opts *AffiliateListNewsOpts) ([]*NewsItem, *Response, error) {
	path := "/3.0/news"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateListNewsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.slice(), resp, nil
}

// affiliateGetNewsResponse specifies response for GetNews.
type affiliateGetNewsResponse struct {
	News *NewsItem `json:"news"`
}

// GetNews get news by ID.
func (s *AffiliateService) GetNewsByID(ctx context.Context, id string) (*NewsItem, *Response, error) {
	path := fmt.Sprintf("/3.0/news/%s", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateGetNewsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.News, resp, nil
}

// affiliateListPixelsResponse specifies response for ListPixels.
type affiliateListPixelsResponse struct {
	Pixel map[int]*Pixel `json:"pixel"`
}

func (a *affiliateListPixelsResponse) slice() []*Pixel {
	ret := make([]*Pixel, 0, len(a.Pixel))
	for _, v := range a.Pixel {
		ret = append(ret, v)
	}

	return ret
}

// ListPixels gets list of a partner’s pixels.
func (s *AffiliateService) ListPixels(ctx context.Context) ([]*Pixel, *Response, error) {
	path := "/3.0/partner/pixels"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateListPixelsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.slice(), resp, nil
}

// AffiliateCreatePixelOpts specifies options for CreatePixel.
type AffiliateCreatePixelOpts struct {
	OfferID  int    `schema:"offer_id"`  // REQUIRED Offer’s ID
	Name     string `schema:"name"`      // REQUIRED Name
	Code     string `schema:"code"`      // REQUIRED Code (Available: <script>…code…</scipt>, <img …>, <iframe src=“…”></iframe>)
	CodeType string `schema:"code_type"` // REQUIRED Code type (Available: javascript, iframe, image)
}

// affiliateCreatePixelResponse specifies response for CreatePixel.
type affiliateCreatePixelResponse struct {
	Pixel *Pixel `json:"pixel"`
}

// CreatePixel adds a partner’s pixel.
func (s *AffiliateService) CreatePixel(ctx context.Context, opts *AffiliateCreatePixelOpts) (*Pixel, *Response, error) {
	path := "/3.0/partner/pixel"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateCreatePixelResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Pixel, resp, nil
}

// AffiliateUpdatePixelOpts specifies options for UpdatePixel.
type AffiliateUpdatePixelOpts struct {
	Name     string `schema:"name,omitempty"`      // Name
	Code     string `schema:"code,omitempty"`      // Code (Available: <script>…code…</scipt>, <img …>, <iframe src=“…”></iframe>)
	CodeType string `schema:"code_type,omitempty"` // Code type (Available: javascript, iframe, image)
}

// affiliateUpdatePixelResponse specifies response for UpdatePixel.
type affiliateUpdatePixelResponse struct {
	Pixel *Pixel `json:"pixel"`
}

// UpdatePixel edits a partner’s pixel.
func (s *AffiliateService) UpdatePixel(ctx context.Context, id int, opts *AffiliateUpdatePixelOpts) (*Pixel, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/pixel/%d", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateUpdatePixelResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Pixel, resp, nil
}

// affiliateDeletePixelResponse specifies response for DeletePixel.
type affiliateDeletePixelResponse struct {
	Pixel *Pixel `json:"pixel"`
}

// Pixel remove deletes a partner’s pixel.
func (s *AffiliateService) DeletePixel(ctx context.Context, id int) (*Pixel, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/pixel/%d/remove", id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateDeletePixelResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Pixel, resp, nil
}

// affiliateGetAffiliateBalanceResponse specifies response for GetAffiliateBalance.
type affiliateGetAffiliateBalanceResponse struct {
	Balance Balance `json:"balance"`
}

// GetAffiliateBalance gets current affiliate balance.
func (s *AffiliateService) GetAffiliateBalance(ctx context.Context) (Balance, *Response, error) {
	path := "/3.0/balance"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateGetAffiliateBalanceResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Balance, resp, nil
}

// AffiliateGetSmartLinkCategoriesOpts specifies options for GetSmartLinkCategories.
type AffiliateGetSmartLinkCategoriesOpts struct {
	ID []string `schema:"id,omitempty"` // SmartLink categories ID collections
}

// affiliateGetSmartLinkCategoriesResponse specifies response for GetSmartLinkCategories.
type affiliateGetSmartLinkCategoriesResponse struct {
	Data []*SmartLinkCategory `json:"data"`
}

// GetSmartLinkCategories gets SmartLink categories list.
func (s *AffiliateService) GetSmartLinkCategories(ctx context.Context, opts *AffiliateGetSmartLinkCategoriesOpts) ([]*SmartLinkCategory, *Response, error) {
	path := "/3.0/partner/smartlink/categories"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateGetSmartLinkCategoriesResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Data, resp, nil
}

// affiliateGetSmartLinkOfferCountResponse specifies response for GetSmartLinkOfferCount.
type affiliateGetSmartLinkOfferCountResponse struct {
	Data struct {
		Count int `json:"count"`
	} `json:"data"`
}

// GetSmartLinkOfferCount gets SmartLink offer count.
func (s *AffiliateService) GetSmartLinkOfferCount(ctx context.Context, id string) (int, *Response, error) {
	path := fmt.Sprintf("/3.0/partner/smartlink/category/%s/offers-count", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return 0, nil, err
	}

	body := new(affiliateGetSmartLinkOfferCountResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return 0, nil, err
	}

	return body.Data.Count, resp, nil
}

// affiliateGetReferralsResponse specifies response for GetReferrals.
type affiliateGetReferralsResponse struct {
	Referrals map[string]*Affiliate
}

func (a *affiliateGetReferralsResponse) slice() []*Affiliate {
	ret := make([]*Affiliate, 0, len(a.Referrals))
	for _, v := range a.Referrals {
		ret = append(ret, v)
	}

	return ret
}

// GetReferrals gets referrals by partner ID.
func (s *AffiliateService) GetReferrals(ctx context.Context, affiliateID uint64) ([]*Affiliate, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/partner/%d/referrals", affiliateID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return nil, nil, err
	}

	body := new(affiliateGetReferralsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.slice(), resp, nil
}
