package uploadpost

import (
	"context"
	"net/url"
	"strings"
)

// GetAnalytics returns analytics data for all connected platforms of a profile.
func (c *Client) GetAnalytics(ctx context.Context, profileUsername string, opts AnalyticsOptions) (*AnalyticsResponse, error) {
	params := map[string]string{}
	if len(opts.Platforms) > 0 {
		params["platforms"] = strings.Join(opts.Platforms, ",")
	}
	if opts.PageID != "" {
		params["page_id"] = opts.PageID
	}
	if opts.PageURN != "" {
		params["page_urn"] = opts.PageURN
	}

	var resp AnalyticsResponse
	err := c.getJSON(ctx, "/analytics/"+url.PathEscape(profileUsername), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTotalImpressions returns aggregated impression metrics for a profile.
func (c *Client) GetTotalImpressions(ctx context.Context, profileUsername string, opts ImpressionsOptions) (*ImpressionsResponse, error) {
	params := map[string]string{}
	if opts.Period != "" {
		params["period"] = opts.Period
	}
	if opts.StartDate != "" {
		params["start_date"] = opts.StartDate
	}
	if opts.EndDate != "" {
		params["end_date"] = opts.EndDate
	}
	if opts.Date != "" {
		params["date"] = opts.Date
	}
	if len(opts.Platforms) > 0 {
		params["platform"] = strings.Join(opts.Platforms, ",")
	}
	if opts.Breakdown {
		params["breakdown"] = "true"
	}
	if len(opts.Metrics) > 0 {
		params["metrics"] = strings.Join(opts.Metrics, ",")
	}

	var resp ImpressionsResponse
	err := c.getJSON(ctx, "/uploadposts/total-impressions/"+url.PathEscape(profileUsername), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPostAnalytics returns per-platform analytics for a specific uploaded post.
// requestID is the request_id returned by an upload call.
func (c *Client) GetPostAnalytics(ctx context.Context, requestID string) (*PostAnalyticsResponse, error) {
	var resp PostAnalyticsResponse
	err := c.getJSON(ctx, "/uploadposts/post-analytics/"+url.PathEscape(requestID), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPostAnalyticsByPlatformID returns analytics for any post using its native
// platform ID (e.g. an Instagram media ID).
func (c *Client) GetPostAnalyticsByPlatformID(ctx context.Context, platformPostID, platform, user string) (*PostAnalyticsResponse, error) {
	var resp PostAnalyticsResponse
	err := c.getJSON(ctx, "/uploadposts/post-analytics", map[string]string{
		"platform_post_id": platformPostID,
		"platform":         platform,
		"user":             user,
	}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlatformMetrics returns the available metrics configuration for all
// supported platforms.
func (c *Client) GetPlatformMetrics(ctx context.Context) (map[string]PlatformMetric, error) {
	var resp map[string]PlatformMetric
	if err := c.getJSON(ctx, "/uploadposts/platform-metrics", nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
