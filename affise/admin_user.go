package affise

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID          string       `json:"id"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Email       string       `json:"email"`
	Skype       string       `json:"skype"`
	Roles       []string     `json:"roles"`
	APIKey      string       `json:"api_key"`
	WorkHours   string       `json:"work_hours"`
	UpdatedAt   string       `json:"updated_at"`
	CreatedAt   string       `json:"created_at"`
	LastLoginAt string       `json:"last_login_at"`
	Type        string       `json:"type"`
	Avatar      string       `json:"avatar"`
	Info        string       `json:"info"`
	Password    string       `json:"password"`
	Permissions *Permissions `json:"permissions"`
}

type Permissions struct {
	json.RawMessage
}

// User roles.
const (
	RoleAdmin                       = "ROLE_ADMIN"                          // Administrator
	RoleManagerAffiliate            = "ROLE_MANAGER_AFFILIATE"              // Affiliate manager
	RoleManagerSales                = "ROLE_MANAGER_SALES"                  // Sales manager
	RoleSectionOffer                = "ROLE_SECTION_OFFER"                  // Offer section
	RoleSectionSupplier             = "ROLE_SECTION_SUPPLIER"               // Supplier section
	RoleSectionDashboard            = "ROLE_SECTION_DASHBOARD"              // Dashboard section
	RoleSectionNews                 = "ROLE_SECTION_NEWS"                   // News section
	RoleSectionCategory             = "ROLE_SECTION_CATEGORY"               // Category section
	RoleSectionPartner              = "ROLE_SECTION_PARTNER"                // Affiliate section
	RoleSectionPayment              = "ROLE_SECTION_PAYMENT"                // Billing section
	RoleSectionTicket               = "ROLE_SECTION_TICKET"                 // Ticket section
	RoleSectionStats                = "ROLE_SECTION_STATS"                  // Satistics section
	RoleSectionStatsCommon          = "ROLE_SECTION_STATS_COMMON"           // Daily stats
	RoleSectionStatsConversion      = "ROLE_SECTION_STATS_CONVERSION"       // Conversions
	RoleSectionStatsSupplier        = "ROLE_SECTION_STATS_SUPPLIER"         // Suppliers
	RoleSectionStatsOffer           = "ROLE_SECTION_STATS_OFFER"            // Offers
	RoleSectionStatsCountry         = "ROLE_SECTION_STATS_COUNTRY"          // Countries
	RoleSectionStatsCity            = "ROLE_SECTION_STATS_CITY"             // Cities
	RoleSectionStatsOs              = "ROLE_SECTION_STATS_OS"               // OS
	RoleSectionStatsGoal            = "ROLE_SECTION_STATS_GOAL"             // Goals
	RoleSectionStatsDevice          = "ROLE_SECTION_STATS_DEVICE"           // Devices
	RoleSectionStatsLimits          = "ROLE_SECTION_STATS_LIMITS"           // Limits statistics
	RoleSectionStatsPartnerPostback = "ROLE_SECTION_STATS_PARTNER_POSTBACK" // User postbacks
	RoleSectionStatsServerPostback  = "ROLE_SECTION_STATS_SERVER_POSTBACK"  // Server postbacks
	RoleSectionAutomation           = "ROLE_SECTION_AUTOMATION"             // Automation
	RoleSectionStatsComparison      = "ROLE_SECTION_STATS_COMPARISON"       // Comparison report
)

// User types.
const (
	UserCommonManager    = "common_manager"    // Administrator
	UserAffiliateManager = "affiliate_manager" // Affiliate manager
	UserAccountManager   = "account_manager"   // Account manager
)

type AdminUserService struct {
	client *Client
}

// AdminUserListOpts specifies options for List.
type AdminUserListOpts struct {
	Page      int    `schema:"page,omitempty"`       // Page of entities
	Limit     int    `schema:"limit,omitempty"`      // Limit of entities
	UpdatedAt string `schema:"updated_at,omitempty"` // Get users that have been updated from this date (format YYYY-MM-DD)
	Q         string `schema:"q,omitempty"`          // Search query
}

// adminUserListResponse specifies response for List.
type adminUserListResponse struct {
	Users []*User `json:"users"`
}

// List gets a list of users.
func (s *AdminUserService) List(ctx context.Context, opts *AdminUserListOpts) ([]*User, *Response, error) {
	path := "/3.0/admin/users"

	req, err := s.client.NewRequestOpts(ctx, http.MethodGet, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminUserListResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Users, resp, nil
}

// AdminUserCreateOpts specifies options for Create.
type AdminUserCreateOpts struct {
	Email     string   `schema:"email"`                // REQUIRED Email
	Password  string   `schema:"password"`             // REQUIRED Password (Available: at least 6 characters)
	FirstName string   `schema:"first_name"`           // REQUIRED Name
	LastName  string   `schema:"last_name"`            // REQUIRED Last name
	Roles     []string `schema:"roles"`                // REQUIRED Array off allowed roles. See roles
	Skype     string   `schema:"skype,omitempty"`      // Skype
	WorkHours string   `schema:"work_hours,omitempty"` // Working time
	Avatar    string   `schema:"avatar,omitempty"`     // Base64 encoded image. Allowed formats: jpg, jpeg
}

// adminUserGetResponse specifies response for Get.
type adminUserGetResponse struct {
	User *User `json:"user"`
}

// Get reads single user.
func (s *AdminUserService) Get(ctx context.Context, id string) (*User, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/user/%s", id)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminUserGetResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.User, resp, nil
}

// adminUserCreateResponse specifies response for Create.
type adminUserCreateResponse struct {
	User *User `json:"user"`
}

// Create adds a new user.
func (s *AdminUserService) Create(ctx context.Context, opts *AdminUserCreateOpts) (*User, *Response, error) {
	path := "/3.0/admin/user"

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminUserCreateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.User, resp, nil
}

// AdminUserUpdateOpts specifies options for Update.
type AdminUserUpdateOpts struct {
	Email     string   `schema:"email,omitempty"`      // Email
	Password  string   `schema:"password,omitempty"`   // Password (Available: at least 6 characters)
	FirstName string   `schema:"first_name,omitempty"` // Name
	LastName  string   `schema:"last_name,omitempty"`  // Last name
	Roles     []string `schema:"roles,omitempty"`      // Array off allowed roles. See roles
	Skype     string   `schema:"skype,omitempty"`      // Skype
	WorkHours string   `schema:"work_hours,omitempty"` // Working time
	Type      string   `schema:"type,omitempty"`       // User type.  See user types
	Avatar    string   `schema:"avatar,omitempty"`     // Base64 encoded image. Allowed formats: jpg, jpeg
}

// adminUserUpdateResponse specifies response for Update.
type adminUserUpdateResponse struct {
	User *User `json:"user"`
}

// Update changes the user.
func (s *AdminUserService) Update(ctx context.Context, id string, opts *AdminUserUpdateOpts) (*User, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/user/%s", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminUserUpdateResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.User, resp, nil
}

// adminUserChangeAPIKeyResponse specifies response for ChangeAPIKey.
type adminUserChangeAPIKeyResponse struct {
	User *User `json:"user"`
}

// ChangeAPIKey changes user api key.
func (s *AdminUserService) ChangeAPIKey(ctx context.Context, id string) (*User, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/user/api_key/%s", id)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminUserChangeAPIKeyResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.User, resp, nil
}

// AdminUserChangePasswordOpts specifies options for ChangePassword.
type AdminUserChangePasswordOpts struct {
	Password string `schema:"password"` // REQUIRED Password (Available: at least 6 characters)
}

// adminUserChangePasswordResponse specifies response for ChangePassword.
type adminUserChangePasswordResponse struct {
	User *User `json:"user"`
}

// ChangePassword changes user password.
func (s *AdminUserService) ChangePassword(ctx context.Context, id string, opts *AdminUserChangePasswordOpts) (*User, *Response, error) {
	path := fmt.Sprintf("/3.0/admin/user/%s/password", id)

	req, err := s.client.NewRequestOpts(ctx, http.MethodPost, path, opts, nil, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminUserChangePasswordResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.User, resp, nil
}

// AdminUserUpdatePermissionsOpts specifies options for UpdatePermissions.
type AdminUserUpdatePermissionsOpts struct {
	Permissions *Permissions `json:"permissions"`
}

// adminUserUpdatePermissionsResponse specifies response for UpdatePermissions.
type adminUserUpdatePermissionsResponse struct {
	Permissions *Permissions `json:"permissions"`
}

// UpdatePermissions updates user permissions.
func (s *AdminUserService) UpdatePermissions(ctx context.Context, id string, opts *AdminUserUpdatePermissionsOpts) (*Permissions, *Response, error) {
	path := fmt.Sprintf("/3.1/user/%s/permissions", id)

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(opts)
	if err != nil {
		return nil, nil, fmt.Errorf("encode opts err: %w", err)
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, buf, true)
	if err != nil {
		return nil, nil, err
	}

	body := new(adminUserUpdatePermissionsResponse)
	resp, err := s.client.Do(req, body)
	if err != nil {
		return nil, nil, err
	}

	return body.Permissions, resp, nil
}
