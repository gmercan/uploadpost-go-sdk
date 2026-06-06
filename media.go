package uploadpost

import "context"

// GetMedia returns recent posts from a connected social account.
// Supported platforms: instagram, tiktok, youtube, linkedin, facebook, x,
// threads, pinterest, bluesky, reddit.
//
// For LinkedIn, pass pageURN to target a specific organisation page or "me" to
// force the personal profile.
func (c *Client) GetMedia(ctx context.Context, platform, user, pageURN string) (*MediaResponse, error) {
	params := map[string]string{
		"platform": platform,
		"user":     user,
	}
	if pageURN != "" {
		params["page_urn"] = pageURN
	}

	var resp MediaResponse
	if err := c.getJSON(ctx, "/uploadposts/media", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFacebookPages returns the Facebook Pages connected to a profile.
// Pass an empty string for profile to use the default profile.
func (c *Client) GetFacebookPages(ctx context.Context, profile string) (*PagesResponse, error) {
	params := map[string]string{}
	if profile != "" {
		params["profile"] = profile
	}
	var resp PagesResponse
	if err := c.getJSON(ctx, "/uploadposts/facebook/pages", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLinkedinPages returns the LinkedIn Pages connected to a profile.
func (c *Client) GetLinkedinPages(ctx context.Context, profile string) (*PagesResponse, error) {
	params := map[string]string{}
	if profile != "" {
		params["profile"] = profile
	}
	var resp PagesResponse
	if err := c.getJSON(ctx, "/uploadposts/linkedin/pages", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPinterestBoards returns the Pinterest boards connected to a profile.
func (c *Client) GetPinterestBoards(ctx context.Context, profile string) (*BoardsResponse, error) {
	params := map[string]string{}
	if profile != "" {
		params["profile"] = profile
	}
	var resp BoardsResponse
	if err := c.getJSON(ctx, "/uploadposts/pinterest/boards", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGoogleBusinessLocations returns all Google Business Profile locations for a profile.
func (c *Client) GetGoogleBusinessLocations(ctx context.Context, profile string) (*GoogleBusinessLocationsResponse, error) {
	params := map[string]string{}
	if profile != "" {
		params["profile"] = profile
	}
	var resp GoogleBusinessLocationsResponse
	if err := c.getJSON(ctx, "/uploadposts/google-business/locations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelectGoogleBusinessLocation selects a specific Google Business Profile
// location for a profile.
func (c *Client) SelectGoogleBusinessLocation(ctx context.Context, locationID, profile string) (*SimpleResponse, error) {
	payload := map[string]string{"location_id": locationID}
	if profile != "" {
		payload["profile"] = profile
	}
	var resp SimpleResponse
	if err := c.postJSON(ctx, "/uploadposts/google-business/locations/select", payload, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}