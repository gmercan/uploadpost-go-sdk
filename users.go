package uploadpost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// ListUsers returns all profiles associated with the API key.
func (c *Client) ListUsers(ctx context.Context) (*UsersResponse, error) {
	var resp UsersResponse
	if err := c.getJSON(ctx, "/uploadposts/users", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateUser creates a new profile.
func (c *Client) CreateUser(ctx context.Context, username string) (*SimpleResponse, error) {
	var resp SimpleResponse
	if err := c.postJSON(ctx, "/uploadposts/users", map[string]string{"username": username}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteUser permanently deletes a profile.
func (c *Client) DeleteUser(ctx context.Context, username string) (*SimpleResponse, error) {
	var resp SimpleResponse
	if err := c.deleteJSON(ctx, "/uploadposts/users", map[string]string{"username": username}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateJWT generates a JWT token and connection URL for white-label
// platform integration.
func (c *Client) GenerateJWT(ctx context.Context, username string, opts JWTOptions) (*JWTResponse, error) {
	payload := map[string]interface{}{
		"username": username,
	}
	if opts.RedirectURL != "" {
		payload["redirect_url"] = opts.RedirectURL
	}
	if opts.LogoImage != "" {
		payload["logo_image"] = opts.LogoImage
	}
	if opts.RedirectButtonText != "" {
		payload["redirect_button_text"] = opts.RedirectButtonText
	}
	if len(opts.Platforms) > 0 {
		payload["platforms"] = opts.Platforms
	}
	if opts.ShowCalendar != nil {
		payload["show_calendar"] = *opts.ShowCalendar
	}
	if opts.ReadonlyCalendar != nil {
		payload["readonly_calendar"] = *opts.ReadonlyCalendar
	}
	if opts.ConnectTitle != "" {
		payload["connect_title"] = opts.ConnectTitle
	}
	if opts.ConnectDescription != "" {
		payload["connect_description"] = opts.ConnectDescription
	}
	if opts.Language != "" {
		payload["language"] = opts.Language
	}

	var resp JWTResponse
	if err := c.postJSON(ctx, "/uploadposts/users/generate-jwt", payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateJWT validates an existing JWT token. Per the API spec, this endpoint
// uses Bearer token authorization instead of the standard Apikey header.
func (c *Client) ValidateJWT(ctx context.Context, jwt string) (*SimpleResponse, error) {
	var resp SimpleResponse
	if err := c.postJSONBearer(ctx, "/uploadposts/users/validate-jwt", jwt, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUserPreferences returns the current user preferences (e.g. calendar settings).
func (c *Client) GetUserPreferences(ctx context.Context) (*UserPreferences, error) {
	var raw struct {
		Success     bool            `json:"success"`
		Preferences json.RawMessage `json:"preferences,omitempty"`
	}
	if err := c.getJSON(ctx, "/uploadposts/users/preferences", nil, &raw); err != nil {
		return nil, err
	}
	var prefs UserPreferences
	if raw.Preferences != nil {
		if err := json.Unmarshal(raw.Preferences, &prefs); err != nil {
			return nil, fmt.Errorf("uploadpost: failed to parse preferences: %w", err)
		}
	}
	return &prefs, nil
}

// UpdateUserPreferences updates user preferences.
func (c *Client) UpdateUserPreferences(ctx context.Context, opts UserPreferencesOptions) (*SimpleResponse, error) {
	payload := map[string]interface{}{}
	if opts.WeekStartDay != nil {
		payload["week_start_day"] = *opts.WeekStartDay
	}
	var resp SimpleResponse
	if err := c.postJSON(ctx, "/uploadposts/users/preferences", payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNotificationConfig returns the current webhook/notification configuration.
func (c *Client) GetNotificationConfig(ctx context.Context) (*NotificationConfig, error) {
	var raw struct {
		Success bool              `json:"success"`
		Config  *NotificationConfig `json:"config,omitempty"`
	}
	if err := c.getJSON(ctx, "/uploadposts/notification-config", nil, &raw); err != nil {
		return nil, err
	}
	if raw.Config == nil {
		return &NotificationConfig{}, nil
	}
	return raw.Config, nil
}

// GetUserProfile returns a single profile by username.
func (c *Client) GetUserProfile(ctx context.Context, username string) (*UserProfile, error) {
	if username == "" {
		return nil, fmt.Errorf("uploadpost: username is required")
	}
	var raw struct {
		Success bool        `json:"success"`
		Profile UserProfile `json:"profile,omitempty"`
	}
	if err := c.getJSON(ctx, "/uploadposts/users/"+url.PathEscape(username), nil, &raw); err != nil {
		return nil, err
	}
	return &raw.Profile, nil
}

// GetCurrentUser validates the API key and returns account info (email, plan,
// preferences).
func (c *Client) GetCurrentUser(ctx context.Context) (*CurrentUser, error) {
	var resp CurrentUser
	if err := c.getJSON(ctx, "/uploadposts/me", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateNotificationConfig updates the webhook/notification configuration.
func (c *Client) UpdateNotificationConfig(ctx context.Context, opts NotificationConfigOptions) (*SimpleResponse, error) {
	payload := map[string]interface{}{}
	if len(opts.WebhookEvents) > 0 {
		payload["webhook_events"] = opts.WebhookEvents
	}
	if opts.WebhookURL != "" {
		payload["webhook_url"] = opts.WebhookURL
	}
	var resp SimpleResponse
	if err := c.postJSON(ctx, "/uploadposts/notification-config", payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}