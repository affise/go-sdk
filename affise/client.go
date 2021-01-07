package affise

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// Version is the library's version.
	Version = "undefined"

	// UserAgent is the value for the library part of the User-Agent header that is sent with each request.
	UserAgent = "go-sdk/" + Version
)

type Client struct {
	httpClient *http.Client
	BaseURL    *url.URL
	AdminURL   *url.URL
	APIKey     string
	UserAgent  string

	// Services used for communicating with the API
	AdminAdvertiser        *AdminAdvertiserService
	AdminAdvertiserBilling *AdminAdvertiserBillingService
	AdminConversion        *AdminConversionService
	AdminAffiliate         *AdminAffiliateService
	AdminOffer             *AdminOfferService
	AdminOther             *AdminOtherService
	AdminPreset            *AdminPresetService
	AdminUser              *AdminUserService
	Affiliate              *AffiliateService
	Offer                  *OfferService
	Other                  *OtherService
	Statistic              *StatisticService
}

// A ClientOption is used to configure a Client.
type ClientOption func(*Client) error

// WithBaseURL is a client option for setting the base URL.
func WithBaseURL(ua string) ClientOption {
	return func(client *Client) error {
		u, err := url.Parse(ua)
		if err != nil {
			return fmt.Errorf("parse base URL err: %w", err)
		}
		client.BaseURL = u

		return nil
	}
}

func WithAdminURL(ua string) ClientOption {
	return func(client *Client) error {
		u, err := url.Parse(ua)
		if err != nil {
			return fmt.Errorf("parse admin URL err: %w", err)
		}
		client.AdminURL = u

		return nil
	}
}

// WithAPIKey configures a Client to use the specified api key for authentication.
func WithAPIKey(key string) ClientOption {
	return func(client *Client) error {
		client.APIKey = key

		return nil
	}
}

// WithBaseURL is a client option for setting the http.Client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = httpClient

		return nil
	}
}

// NewClient creates a new client.
func NewClient(options ...ClientOption) (*Client, error) {
	client := &Client{
		httpClient: http.DefaultClient,
		UserAgent:  UserAgent,
	}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	client.AdminAdvertiser = &AdminAdvertiserService{client: client}
	client.AdminAdvertiserBilling = &AdminAdvertiserBillingService{client: client}
	client.AdminConversion = &AdminConversionService{client: client}
	client.AdminAffiliate = &AdminAffiliateService{client: client}
	client.AdminOffer = &AdminOfferService{client: client}
	client.AdminOther = &AdminOtherService{client: client}
	client.AdminPreset = &AdminPresetService{client: client}
	client.AdminUser = &AdminUserService{client: client}
	client.Affiliate = &AffiliateService{client: client}
	client.Offer = &OfferService{client: client}
	client.Other = &OtherService{client: client}
	client.Statistic = &StatisticService{client: client}

	return client, nil
}

var (
	errAdminURL = errors.New("invalid admin URL")
	errBaseURL  = errors.New("invalid base URL")
)

func (c *Client) parseURL(urlStr string, isAdmin bool) (*url.URL, error) {
	var u *url.URL
	if isAdmin {
		if u = c.AdminURL; u == nil {
			return nil, errAdminURL
		}
	} else {
		if u = c.BaseURL; u == nil {
			return nil, errBaseURL
		}
	}

	return u.Parse(urlStr)
}

// NewRequest creates an HTTP request against the API. The returned request
// is assigned with ctx and has all necessary headers.
func (c *Client) NewRequest(ctx context.Context,
	method, urlStr string, body io.Reader, isAdmin bool) (*http.Request, error) {
	u, err := c.parseURL(urlStr, isAdmin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest err: %w", err)
	}
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("API-Key", c.APIKey)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) NewRequestOpts(ctx context.Context,
	method, urlStr string, opts interface{}, body io.Reader, isAdmin bool) (*http.Request, error) {
	var val url.Values
	var err error

	if valuer, ok := opts.(valuer); ok {
		val, err = valuer.values()
	} else {
		val, err = defaultEncoder.encode(opts)
	}
	if err != nil {
		return nil, err
	}

	urlStr = fmt.Sprintf("%s?%s", urlStr, val.Encode())

	return c.NewRequest(ctx, method, urlStr, body, isAdmin)
}

// Do performs an HTTP request against the API.
func (c *Client) Do(r *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do err: %w", err)
	}
	response := &Response{Response: resp}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_ = resp.Body.Close()

		return response, fmt.Errorf("ioutil.ReadAll err: %w", err)
	}
	_ = resp.Body.Close()
	resp.Body = ioutil.NopCloser(bytes.NewReader(body))

	if err := response.readMeta(body); err != nil {
		return response, fmt.Errorf("read response meta data err: %w", err)
	}

	if err := c.checkResponse(response); err != nil {
		return response, err
	}

	if v == nil {
		return response, nil
	}

	if w, ok := v.(io.Writer); ok {
		if _, err := io.Copy(w, bytes.NewReader(body)); err != nil {
			return response, fmt.Errorf("io.Copy err: %w", err)
		}
	} else if err := json.Unmarshal(body, v); err != nil {
		return response, fmt.Errorf("json.Unmarshal err: %w", err)
	}

	return response, nil
}

func (c *Client) checkResponse(r *Response) error {
	if (r.StatusCode >= 400 && r.StatusCode <= 599) || r.Meta.Status != 1 {
		return responseErr(r)
	}

	return nil
}

type ResponseErr struct {
	Method      string
	URL         *url.URL
	Status      string
	MetaStatus  int
	MetaMessage string
}

func responseErr(r *Response) *ResponseErr {
	return &ResponseErr{
		Method:      r.Request.Method,
		URL:         r.Request.URL,
		Status:      r.Status,
		MetaStatus:  r.Meta.Status,
		MetaMessage: r.Meta.Message,
	}
}

// Error implements error interface.
func (r *ResponseErr) Error() string {
	return fmt.Sprintf("%s %v %s (status %d) err: %s", r.Method, r.URL, r.Status, r.MetaStatus, r.MetaMessage)
}

// Meta represents meta information included in an API response.
type Meta struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaller.
func (m *Meta) UnmarshalJSON(data []byte) error {
	type RawMeta struct {
		Status     int         `json:"status"`
		Message    interface{} `json:"message"`
		Pagination *Pagination `json:"pagination,omitempty"`
	}

	raw := RawMeta{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("json.Unmarshal err: %w", err)
	}

	m.Status = raw.Status
	m.Pagination = raw.Pagination
	// some handles use this field as an object
	if s, ok := raw.Message.(string); ok {
		m.Message = s
	}

	return nil
}

// Pagination represents pagination meta information.
type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
	NextPage   int `json:"next_page"`
}

// Response represents a response from the API. It embeds http.Response.
type Response struct {
	*http.Response
	Meta Meta
}

func (r *Response) readMeta(body []byte) error {
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		if err := json.Unmarshal(body, &r.Meta); err != nil {
			return fmt.Errorf("json.Unmarshal err: %w", err)
		}
	}

	return nil
}
